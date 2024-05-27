package main

import (
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

	origins []string

	port int

	dbUser, dbPassword, dbHost, dbName string
	dbPort                             int
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	dbPassword = os.Getenv("DB_PASSWORD")
	dbUser = os.Getenv("DB_USER")
	dbHost = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_NAME")
	dbPort, err = strconv.Atoi(strings.TrimSpace(os.Getenv("DB_PORT")))
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

	tokenLifetime, err = strconv.Atoi(strings.TrimSpace(os.Getenv("TOKEN_LIFETIME")))
	if err != nil {
		log.Fatal("Can't parse token lifetime: ", err)
	}

	secret = os.Getenv("SECRET")
	if secret == "" {
		log.Fatal("Missing secret in environvent variables")
	}

	origins = strings.Split(os.Getenv("ALLOW_ORIGINS"), ",")
	if len(origins) == 0 {
		log.Fatal("Missing allow origins")
	}

	port, err = strconv.Atoi(strings.TrimSpace(os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Can't parse port: ", err)
	}
}

func main() {
	dbPostgres, err := setupst.GetConnect(dbUser, dbPassword, dbHost, dbName, dbPort, 5)
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
	config.AllowOrigins = origins
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

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router.Handler(),
	}

	// ctx, cancel := context.WithCancel(context.Background())

	// go func(ctx context.Context, srv *http.Server) {
	// 	<-ctx.Done()
	// 	err := server.Shutdown(ctx)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }(ctx, server)

	// go func(cancel context.CancelFunc) {
	// 	fmt.Scanln()
	// 	cancel()
	// }(cancel)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
