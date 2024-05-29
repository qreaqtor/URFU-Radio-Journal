package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	articlehand "urfu-radio-journal/internal/handlers/article"
	authand "urfu-radio-journal/internal/handlers/auth"
	commentshand "urfu-radio-journal/internal/handlers/comments"
	councilhand "urfu-radio-journal/internal/handlers/council"
	editionhand "urfu-radio-journal/internal/handlers/edition"
	filehand "urfu-radio-journal/internal/handlers/files"
	"urfu-radio-journal/internal/handlers/middleware"
	redactionhand "urfu-radio-journal/internal/handlers/redaction"
	articlesrv "urfu-radio-journal/internal/services/article"
	authsrv "urfu-radio-journal/internal/services/auth"
	commentsrv "urfu-radio-journal/internal/services/comments"
	councilsrv "urfu-radio-journal/internal/services/council"
	editionsrv "urfu-radio-journal/internal/services/edition"
	filesrv "urfu-radio-journal/internal/services/files"
	"urfu-radio-journal/internal/services/files/buckets"
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
	"github.com/gin-gonic/gin"
)

const (
	videosBucket    = "videos"
	imagesBucket    = "images"
	documentsBucket = "documents"

	fileInfoTable = "fileinfo"
)

// для этого надо завести отдельный файл с конфигами
var (
	adminPassword, adminLogin, secret string
	tokenLifetime                     int

	origins []string

	port int

	apiVersion int

	dbUser, dbPassword, dbHost, dbName string
	dbPort                             int
	connCount                          int

	minioUser, minioPassword, minioEndpoint string
	ssl                                     bool
)

const (
	articlesTable  = "articles"
	commentsTable  = "comments"
	counsilTable   = "counsil"
	editionsTable  = "editions"
	redactionTable = "redaction"
	authorsTable   = "authors"
)

func init() {
	var err error

	dbPassword = os.Getenv("DB_PASSWORD")
	dbUser = os.Getenv("DB_USER")
	dbHost = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_NAME")
	dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Can't parse dbPort: ", err)
	}

	adminPassword = os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		log.Fatal("Missing admin password in environment variables")
	}

	adminLogin = os.Getenv("ADMIN_LOGIN")
	if adminLogin == "" {
		log.Fatal("Missing admin username in environment variables")
	}

	tokenLifetime, err = strconv.Atoi(os.Getenv("TOKEN_LIFETIME"))
	if err != nil {
		log.Fatal("Can't parse token lifetime: ", err)
	}

	secret = os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("Missing secret")
	}

	origins = strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	if len(origins) == 0 {
		log.Fatal("Missing allow origins")
	}

	port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Can't parse port: ", err)
	}

	apiVersion, err = strconv.Atoi(os.Getenv("API_VERSION"))
	if err != nil {
		log.Fatal("Can't parse apiVersion: ", err)
	}

	connCount, err = strconv.Atoi(os.Getenv("CONNECT_COUNT"))
	if err != nil {
		log.Fatal("Can't parse connection count: ", err)
	}

	ssl, err = strconv.ParseBool(os.Getenv("SSL"))
	if err != nil {
		log.Fatal("Can't parse ssl: ", err)
	}

	minioUser = os.Getenv("MINIO_USER")
	if minioUser == "" {
		log.Fatal("Missing minio username in environment variables")
	}

	minioPassword = os.Getenv("MINIO_PASSWORD")
	if minioPassword == "" {
		log.Fatal("Missing minio password in environment variables")
	}

	minioEndpoint = os.Getenv("MINIO_ENDPOINT")
	if minioEndpoint == "" {
		log.Fatal("Missing minio endpoint in environment variables")
	}
}

func main() {
	ctx := context.Background()

	dbPostgres, err := postgrest.GetConnect(dbUser, dbPassword, dbHost, dbName, dbPort, connCount)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbPostgres)

	minioClient, err := miniost.GetConnect(minioUser, minioPassword, minioEndpoint, ssl)
	if err != nil {
		log.Fatal(err)
	}
	err = miniost.InitBuckets(ctx, minioClient,
		videosBucket,
		imagesBucket,
		documentsBucket,
	)
	if err != nil {
		log.Fatal(err)
	}

	config := cors.DefaultConfig()
	config.AllowOrigins = origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "Content-Type")

	// тут инициализация всех стореджей
	articleStorage := articlest.NewArticleStorage(dbPostgres, articlesTable)
	commentStorage := commentst.NewCommentStorage(dbPostgres, commentsTable)
	councilStorage := councilst.NewCouncilStorage(dbPostgres, counsilTable)
	editionStorage := editionst.NewEditionStorage(dbPostgres, editionsTable)
	redactionStorage := redactionst.NewRedactionStorage(dbPostgres, redactionTable)
	fileInfoStorage := fileinfost.NewFileInfoStorage(dbPostgres, fileInfoTable)
	videoStorage := filest.NewFileStorage(minioClient, videosBucket)
	imageStorage := filest.NewFileStorage(minioClient, imagesBucket)
	documentStorage := filest.NewFileStorage(minioClient, documentsBucket)
	authorStorage := authorst.NewAuthorStorage(dbPostgres, authorsTable)

	types := buckets.AllowedContentType{
		videoStorage:    {"video/mp4"},
		imageStorage:    {"image/jpeg"},
		documentStorage: {"application/pdf"},
	}

	// тут всех сервисов
	articleService := articlesrv.NewArticleService(articleStorage, authorStorage)
	authService := authsrv.NewAuthService(adminPassword, adminLogin, secret, tokenLifetime)
	commentService := commentsrv.NewCommentsService(commentStorage)
	councilService := councilsrv.NewCouncilService(councilStorage)
	editionService := editionsrv.NewEditionService(editionStorage)
	redactionService := redactionsrv.NewRedactionService(redactionStorage)
	fileService := filesrv.NewFileService(types, fileInfoStorage)

	// тут хендлеров
	articleHandler := articlehand.NewArticleHandler(articleService)
	authHandler := authand.NewAuthHandler(authService)
	commentHandler := commentshand.NewCommentsHandler(commentService)
	councilHandler := councilhand.NewCouncilHandler(councilService)
	editionHandler := editionhand.NewEditionHandler(editionService)
	redactionHandler := redactionhand.NewRedactionHandler(redactionService)
	fileHandler := filehand.NewFilesHandler(fileService)

	engine := gin.Default()
	engine.Use(cors.New(config))

	router := engine.Group(fmt.Sprintf("/api/v%d",apiVersion))

	authMiddleware := middleware.AuthMiddleware(authService.ValidateToken)

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
	fileRouter.POST("/upload/", authMiddleware, fileHandler.UploadFile)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: engine.Handler(),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
