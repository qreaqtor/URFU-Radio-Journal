package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	articlehand "urfu-radio-journal/internal/handlers/article"
	authand "urfu-radio-journal/internal/handlers/auth"
	"urfu-radio-journal/internal/handlers/middleware"
	articlesrv "urfu-radio-journal/internal/services/article"
	authsrv "urfu-radio-journal/internal/services/auth"
	articlest "urfu-radio-journal/internal/storage/mongo/article"
	setupst "urfu-radio-journal/internal/storage/mongo/setup"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	userMongo, passwordMongo, dbNameMongo string

	adminPassword, loginAdmin, secret string
	tokenLifetime                     int

	frontend string

	port string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	userMongo = os.Getenv("MONGO_USER")
	passwordMongo = os.Getenv("MONGO_PASSWORD")
	dbNameMongo = os.Getenv("DB_NAME")

	adminPassword = os.Getenv("ADMIN_PASSWORD")
	if passwordMongo == "" {
		log.Fatal("Missing admin password in environvent variables.")
	}
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

	dbMongo := setupst.GetConnect(userMongo, passwordMongo, dbNameMongo)

	// тут инициализация все стореджей
	articleStorage := articlest.NewArticleStorage(dbMongo, "articles")

	// тут всех сервисов
	articleService := articlesrv.NewArticleService(articleStorage)
	authService := authsrv.NewAuthService(adminPassword, loginAdmin, secret, tokenLifetime)

	// тут хендлеров
	articleHandler := articlehand.NewArticleHandler(articleService)
	authHandler := authand.NewAuthHandler(authService)

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
