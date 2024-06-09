package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"urfu-radio-journal/internal/config"
	articlehand "urfu-radio-journal/internal/handlers/article"
	authand "urfu-radio-journal/internal/handlers/auth"
	commentshand "urfu-radio-journal/internal/handlers/comments"
	councilhand "urfu-radio-journal/internal/handlers/council"
	editionhand "urfu-radio-journal/internal/handlers/edition"
	filehand "urfu-radio-journal/internal/handlers/files"
	"urfu-radio-journal/internal/handlers/middleware"
	redactionhand "urfu-radio-journal/internal/handlers/redaction"
	"urfu-radio-journal/internal/monitoring"
	articlesrv "urfu-radio-journal/internal/services/article"
	authsrv "urfu-radio-journal/internal/services/auth"
	commentsrv "urfu-radio-journal/internal/services/comments"
	councilsrv "urfu-radio-journal/internal/services/council"
	editionsrv "urfu-radio-journal/internal/services/edition"
	filesrv "urfu-radio-journal/internal/services/files"
	bucketsrv "urfu-radio-journal/internal/services/files/buckets"
	redactionsrv "urfu-radio-journal/internal/services/redaction"
	filest "urfu-radio-journal/internal/storage/minio/files"
	miniost "urfu-radio-journal/internal/storage/minio/setup"
	articlest "urfu-radio-journal/internal/storage/postgres/article"
	authorst "urfu-radio-journal/internal/storage/postgres/author"
	commentst "urfu-radio-journal/internal/storage/postgres/comments"
	councilst "urfu-radio-journal/internal/storage/postgres/council"
	editionst "urfu-radio-journal/internal/storage/postgres/edition"
	fileinfost "urfu-radio-journal/internal/storage/postgres/fileinfo"
	redactionst "urfu-radio-journal/internal/storage/postgres/redaction"
	postgrest "urfu-radio-journal/internal/storage/postgres/setup"

	"github.com/gin-contrib/cors"
	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	fileInfoTable  = "fileinfo"
	articlesTable  = "articles"
	commentsTable  = "comments"
	counsilTable   = "council"
	editionsTable  = "editions"
	redactionTable = "redaction"
	authorsTable   = "authors"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	server, err := getServerWithConf(ctx, cancel, conf)
	if err != nil {
		log.Fatal(err)
	}

	dbPostgres, err := postgrest.GetConnect(conf.PostgresConfig, conf.Ssl)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbPostgres)

	minioClient, err := miniost.GetConnect(conf.MinioConfig, conf.Ssl)
	if err != nil {
		log.Fatal(err)
	}
	err = miniost.InitBuckets(ctx, minioClient, conf.Buckets...)
	if err != nil {
		log.Fatal(err)
	}

	corsConf := cors.DefaultConfig()
	corsConf.AllowOrigins = conf.Origins
	corsConf.AllowMethods = conf.Methods
	corsConf.AllowCredentials = true
	corsConf.AddAllowHeaders(conf.Methods...)

	// тут инициализация всех стореджей
	articleStorage := articlest.NewArticleStorage(dbPostgres, articlesTable)
	commentStorage := commentst.NewCommentStorage(dbPostgres, commentsTable)
	councilStorage := councilst.NewCouncilStorage(dbPostgres, counsilTable)
	editionStorage := editionst.NewEditionStorage(dbPostgres, editionsTable)
	redactionStorage := redactionst.NewRedactionStorage(dbPostgres, redactionTable)
	fileInfoStorage := fileinfost.NewFileInfoStorage(dbPostgres, fileInfoTable)
	authorStorage := authorst.NewAuthorStorage(dbPostgres, authorsTable)

	buckets := make([]bucketsrv.FileRepo, len(conf.Buckets))
	for _, bucketConf := range conf.Buckets {
		buckets = append(buckets, filest.NewFileStorage(minioClient, bucketConf.Name, bucketConf.ContentTypes...))
	}

	// тут всех сервисов
	articleService := articlesrv.NewArticleService(articleStorage, authorStorage)
	authService := authsrv.NewAuthService(conf.AuthConfig)
	commentService := commentsrv.NewCommentsService(commentStorage)
	councilService := councilsrv.NewCouncilService(councilStorage)
	editionService := editionsrv.NewEditionService(editionStorage)
	redactionService := redactionsrv.NewRedactionService(redactionStorage)
	fileService := filesrv.NewFileService(fileInfoStorage, buckets...)

	monitoring := monitoring.NewMonitoring(conf.Monitoring.ContentTypes...)

	// тут хендлеров
	articleHandler := articlehand.NewArticleHandler(articleService)
	authHandler := authand.NewAuthHandler(authService)
	commentHandler := commentshand.NewCommentsHandler(commentService)
	councilHandler := councilhand.NewCouncilHandler(councilService)
	editionHandler := editionhand.NewEditionHandler(editionService)
	redactionHandler := redactionhand.NewRedactionHandler(redactionService)
	fileHandler := filehand.NewFilesHandler(fileService, monitoring)

	engine := gin.Default()
	engine.Use(cors.New(corsConf))

	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router := engine.Group(fmt.Sprintf("/api/v%d", conf.ApiVersion))

	authMiddleware := middleware.Auth(authService.ValidateToken)

	// тут роутов
	articleRouter := router.Group("/article")
	articleRouter.GET("/get/all", articleHandler.GetAllArticles)
	articleRouter.GET("/get/:articleId", articleHandler.GetArticleById)

	articleRouter.POST("/create", authMiddleware, articleHandler.Create)
	articleRouter.PUT("/update", authMiddleware, articleHandler.Update)
	articleRouter.DELETE("/delete/:id", authMiddleware, articleHandler.Delete)

	authRouter := router.Group("/admin/auth")
	authRouter.POST("/login", authHandler.Login)

	commentRouter := router.Group("/comments")
	commentRouter.GET("/get/all", commentHandler.GetAll)

	commentRouter.POST("/create", authMiddleware, commentHandler.Create)
	commentRouter.PATCH("/update", authMiddleware, commentHandler.Update)
	commentRouter.PATCH("/approve", authMiddleware, commentHandler.Approve)
	commentRouter.DELETE("/delete/:id", authMiddleware, commentHandler.Delete)

	councilRouter := router.Group("/council/members")
	councilRouter.GET("/get/all", councilHandler.GetAll)
	councilRouter.GET("/get/:memberId", councilHandler.GetMemberById)

	councilRouter.POST("/create", authMiddleware, councilHandler.Create)
	councilRouter.PUT("/update/:id", authMiddleware, councilHandler.Update)
	councilRouter.DELETE("/delete/:id", authMiddleware, councilHandler.Delete)

	editionRouter := router.Group("/editions")
	editionRouter.GET("/get/all", editionHandler.GetAllEditions)
	editionRouter.GET("/get/:editionId", editionHandler.GetEditionById)

	editionRouter.POST("/create", authMiddleware, editionHandler.CreateEdition)
	editionRouter.PUT("/update", authMiddleware, editionHandler.UpdateEdition)
	editionRouter.DELETE("/delete/:id", authMiddleware, editionHandler.DeleteEdition)

	redactionRouter := router.Group("/redaction/members")
	redactionRouter.GET("/get/all", redactionHandler.GetAll)
	redactionRouter.GET("/get/:memberId", redactionHandler.GetMemberById)

	redactionRouter.POST("/create", authMiddleware, redactionHandler.Create)
	redactionRouter.PUT("/update/:id", authMiddleware, redactionHandler.Update)
	redactionRouter.DELETE("/delete/:id", authMiddleware, redactionHandler.Delete)

	fileRouter := router.Group("/files")
	fileRouter.GET("/download/:fileID", fileHandler.DownloadFile)

	fileRouter.DELETE("/delete/:fileID", authMiddleware, fileHandler.DeleteFile)
	fileRouter.POST("/upload/", authMiddleware, limits.RequestSizeLimiter(conf.MaxFileSize), fileHandler.UploadFile)

	server.Handler = engine.Handler()

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

const (
	prod = "prod"
	dev  = "dev"
)

// This func is not set handler for server
func getServerWithConf(ctx context.Context, cancel context.CancelFunc, conf *config.ServerConfig) (*http.Server, error) {
	server := &http.Server{
		Addr: fmt.Sprintf(":%d", conf.Port),
	}

	switch conf.Env {
	case dev:
		go func(cancel context.CancelFunc) {
			fmt.Scanln()
			cancel()
		}(cancel)

		go func(server *http.Server, ctx context.Context) {
			<-ctx.Done()
			server.Shutdown(ctx)
		}(server, ctx)
	case prod:
		gin.SetMode(gin.ReleaseMode)
	default:
		return nil, fmt.Errorf("bad env value in configuration file")
	}

	return server, nil
}
