package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
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
	articlest "urfu-radio-journal/internal/storage/postgres/article"
	commentst "urfu-radio-journal/internal/storage/postgres/comments"
	councilst "urfu-radio-journal/internal/storage/postgres/council"
	editionst "urfu-radio-journal/internal/storage/postgres/edition"
	redactionst "urfu-radio-journal/internal/storage/postgres/redaction"
	setupst "urfu-radio-journal/internal/storage/postgres/setup"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// для этого надо завести отдельный файл с конфигами
var (
	adminPassword, adminLogin, secret string
	tokenLifetime                     int

	frontend string

	port int

	dbUser, dbPassword, dbHost, dbName string
	dbPort                             int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbPassword = os.Getenv("DB_PASSWORD")
	dbUser = os.Getenv("DB_USER")
	dbHost = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_NAME")
	dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal("Can't parse dbPort.")
	}

	adminPassword = os.Getenv("ADMIN_PASSWORD")
	if adminPassword == "" {
		log.Fatal("Missing admin password in environment variables.")
	}

	adminLogin = os.Getenv("ADMIN_LOGIN")
	if adminLogin == "" {
		log.Fatal("Missing admin username in environment variables.")
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

	port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Can't parse port.")
	}
}

func main() {
	dbPostgres, err := setupst.GetConnect(dbUser, dbPassword, dbHost, dbName, dbPort)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(dbPostgres)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{frontend}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "Cookie")

	// тут инициализация всех стореджей
	articleStorage := articlest.NewArticleStorage(dbPostgres, "articles")
	commentStorage := commentst.NewCommentStorage(dbPostgres)
	councilStorage := councilst.NewCouncilStorage(dbPostgres)
	editionStorage := editionst.NewEditionStorage(dbPostgres, "editions")
	redactionStorage := redactionst.NewRedactionStorage(dbPostgres)

	// тут всех сервисов
	articleService := articlesrv.NewArticleService(articleStorage)
	authService := authsrv.NewAuthService(adminPassword, adminLogin, secret, tokenLifetime)
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

	councilRouter.POST("/create", councilHandler.Create).Use(authMiddleware)
	councilRouter.PUT("/update/:id", councilHandler.Update).Use(authMiddleware)
	councilRouter.DELETE("/delete/:id", councilHandler.Delete).Use(authMiddleware)

	editionRouter := router.Group("/editions")
	editionRouter.GET("/get/all", editionHandler.GetAllEditions)
	editionRouter.GET("/get/:editionId", editionHandler.GetEditionById)
	editionRouter.POST("/create", editionHandler.CreateEdition).Use(authMiddleware)
	editionRouter.PUT("/update", editionHandler.UpdateEdition).Use(authMiddleware)
	editionRouter.DELETE("/delete/:id", editionHandler.DeleteEdition).Use(authMiddleware)

	redactionRouter := router.Group("/redaction/members")
	redactionRouter.GET("/get/all", redactionHandler.GetAll)
	redactionRouter.GET("/get/:memberId", redactionHandler.GetMemberById)
	redactionRouter.POST("/create", redactionHandler.Create).Use(authMiddleware)
	redactionRouter.PUT("/update/:id", redactionHandler.Update).Use(authMiddleware)
	redactionRouter.DELETE("/delete/:id", redactionHandler.Delete).Use(authMiddleware)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router.Handler(),
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context, srv *http.Server) {
		<-ctx.Done()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(ctx, server)

	go func(cancel context.CancelFunc) {
		fmt.Scanln()
		cancel()
	}(cancel)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
