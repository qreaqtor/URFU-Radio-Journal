package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	articlehand "urfu-radio-journal/internal/handlers/article"
	authand "urfu-radio-journal/internal/handlers/auth"
	commentshand "urfu-radio-journal/internal/handlers/comments"
	councilhand "urfu-radio-journal/internal/handlers/council"
	editionhand "urfu-radio-journal/internal/handlers/edition"
	"urfu-radio-journal/internal/handlers/middleware"
	redactionhand "urfu-radio-journal/internal/handlers/redaction"
	articlesrv "urfu-radio-journal/internal/services/article"
	authsrv "urfu-radio-journal/internal/services/auth"
	commentsrv "urfu-radio-journal/internal/services/comments"
	councilsrv "urfu-radio-journal/internal/services/council"
	editionsrv "urfu-radio-journal/internal/services/edition"
	redactionsrv "urfu-radio-journal/internal/services/redaction"
	articlest "urfu-radio-journal/internal/storage/mysql/article"
	commentst "urfu-radio-journal/internal/storage/mysql/comments"
	councilst "urfu-radio-journal/internal/storage/mysql/council"
	editionst "urfu-radio-journal/internal/storage/mysql/edition"
	redactionst "urfu-radio-journal/internal/storage/mysql/redaction"
	setupst "urfu-radio-journal/internal/storage/mysql/setup"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// для этого надо завести отдельный файл с конфигами
var (
	// userMongo, passwordMongo, dbNameMongo string

	adminPassword, loginAdmin, secret string
	tokenLifetime                     int

	frontend string

	port string

	dbUser, dbPassword, addr, dbName, dbPort string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// userMongo = os.Getenv("MONGO_USER")
	// passwordMongo = os.Getenv("MONGO_PASSWORD")
	// dbNameMongo = os.Getenv("DB_NAME")

	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	addr = os.Getenv("DB_ADDRES")
	dbName = os.Getenv("DB_NAME")
	dbPort = os.Getenv("DB_PORT")

	adminPassword = os.Getenv("ADMIN_PASSWORD")
	// if passwordMongo == "" {
	// 	log.Fatal("Missing admin password in environvent variables.")
	// }
	loginAdmin = os.Getenv("ADMIN_LOGIN")
	if loginAdmin == "" {
		log.Fatal("Missing admin username in environvent variables.")
	}
	tokenLifetime, err = strconv.Atoi(os.Getenv("TOKEN_LIFETIME"))
	if err != nil {
		log.Fatal("Can't parse token lifetime.")
	}
	secret = os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("Missing secret in environvent variables.")
	}

	frontend = os.Getenv("FRONTEND_ADRESS")

	port = os.Getenv("PORT")
}

func main() {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{frontend, "http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "Cookie")

	dbPostgres, _ := setupst.GetConnect(dbUser, dbPassword, addr, dbPort, dbName)

	// тут инициализация всех стореджей
	articleStorage := articlest.NewArticleStorage(dbPostgres)
	commentStorage := commentst.NewCommentStorage(dbPostgres)
	councilStorage := councilst.NewCouncilStorage(dbPostgres)
	editionStorage := editionst.NewEditionStorage(dbPostgres)
	redactionStorage := redactionst.NewRedactionStorage(dbPostgres)

	// тут всех сервисов
	articleService := articlesrv.NewArticleService(articleStorage)
	authService := authsrv.NewAuthService(adminPassword, loginAdmin, secret, tokenLifetime)
	commentService := commentsrv.NewCommentsService(commentStorage)
	councilService := councilsrv.NewCouncilService(councilStorage)
	editionService := editionsrv.NewEditionService(editionStorage)
	redactionService := redactionsrv.NewRedactionService(redactionStorage)

	// тут хендлеров
	articleHandler := articlehand.NewArticleHandler(articleService)
	authHandler := authand.NewAuthHandler(authService)
	commentHandler := commentshand.NewCommentsHandler(commentService)
	councilHandler := councilhand.NewCouncilHandler(councilService)
	editionHandler := editionhand.NewEditionHandler(editionService)
	redactionHandler := redactionhand.NewRedactionHandler(redactionService)

	router := gin.Default()
	router.Use(cors.New(config))
	router.Use(middleware.PanicMiddleware()) // не уверен что тут надо это, но пусть пока что будет

	authMiddleware := middleware.AuthMiddleware(authService.ValidateToken)

	// тут роутов
	articleRouter := router.Group("/article")
	articleRouter.GET("/get/all", articleHandler.GetAllArticles)
	articleRouter.GET("/get/:articleId", articleHandler.GetArticleById)
	articleRouter.POST("/create", articleHandler.Create).Use(authMiddleware)
	articleRouter.PUT("/update", articleHandler.Update).Use(authMiddleware)
	articleRouter.DELETE("/delete/:id", articleHandler.Delete).Use(authMiddleware)

	authRouter := router.Group("/admin/auth")
	authRouter.POST("/login", authHandler.Login)

	commentRouter := router.Group("/comments")
	commentRouter.GET("/get/all", commentHandler.GetAll)
	commentRouter.POST("/create", commentHandler.Create)
	commentRouter.PATCH("/update", commentHandler.Update).Use(authMiddleware)
	commentRouter.PATCH("/approve", commentHandler.Approve).Use(authMiddleware)
	commentRouter.DELETE("/delete/:id", commentHandler.Delete).Use(authMiddleware)

	councilRouter := router.Group("/council/members")
	councilRouter.GET("/get/all", councilHandler.GetAll)
	councilRouter.GET("/get/:memberId", councilHandler.GetMemberById)
	councilRouter.POST("/create", councilHandler.Create)
	councilRouter.PUT("/update/:id", councilHandler.Update)
	councilRouter.DELETE("/delete/:id", councilHandler.Delete)

	editionRouter := router.Group("/editions")
	editionRouter.GET("/get/all", editionHandler.GetAllEditions)
	editionRouter.GET("/get/:editionId", editionHandler.GetEditionById)
	editionRouter.POST("/create", editionHandler.CreateEdition)
	editionRouter.PUT("/update", editionHandler.UpdateEdition)
	editionRouter.DELETE("/delete/:id", editionHandler.DeleteEdition)

	redactionRouter := router.Group("/redaction/members")
	redactionRouter.GET("/get/all", redactionHandler.GetAll)
	redactionRouter.GET("/get/:memberId", redactionHandler.GetMemberById)
	redactionRouter.POST("/create", redactionHandler.Create)
	redactionRouter.PUT("/update/:id", redactionHandler.Update)
	redactionRouter.DELETE("/delete/:id", redactionHandler.Delete)

	// articleStorage := article.NewArticleStorage(dbMongo, "articles")

	// authPath := router.Group("/admin/auth")
	// authHandler := auth.NewAuthHandler()
	// authHandler.RegisterRoutes(authPath)

	// publicFilesPath := router.Group("/files")

	// adminFilesPath := router.Group("/admin/files")
	// adminFilesPath.Use(authHandler.AuthMiddleware())

	// files := filePaths.NewFilesController()
	// files.RegisterRoutes(publicFilesPath, adminFilesPath)

	// publicCommentsPath := router.Group("/comments")

	// adminCommentsPath := router.Group("/admin/comments")
	// adminCommentsPath.Use(authHandler.AuthMiddleware())

	// comments := comments.NewCommentsController()
	// comments.RegisterRoutes(publicCommentsPath, adminCommentsPath)

	// publicArticlePath := router.Group("/articles")

	// adminArticlePath := router.Group("/admin/articles")
	// adminArticlePath.Use(authHandler.AuthMiddleware())

	// articleHandler := article.NewArticleController()
	// articleHandler.RegisterRoutes(publicArticlePath, adminArticlePath)

	// publicEditionPath := router.Group("/editions")

	// adminEditionPath := router.Group("/admin/editions")
	// adminEditionPath.Use(authHandler.AuthMiddleware())

	// edition := edition.NewEditionController(files.GetDeleteHandler(), articleHandler.GetDeleteHandler())
	// edition.RegisterRoutes(publicEditionPath, adminEditionPath)

	// councilPublicPath := router.Group("/council/members")

	// councilAdminPath := router.Group("/admin/council/members")
	// councilAdminPath.Use(authHandler.AuthMiddleware())

	// council := council.NewCouncilController(files.GetDeleteHandler())
	// council.RegisterRoutes(councilPublicPath, councilAdminPath)

	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
