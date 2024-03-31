package main

import (
	"fmt"
	"log"
	"os"

	"urfu-radio-journal/internal/controllers/auth"
	"urfu-radio-journal/internal/controllers/article"
	"urfu-radio-journal/internal/controllers/comments"
	"urfu-radio-journal/internal/controllers/edition"
	"urfu-radio-journal/internal/controllers/filePaths"
	"urfu-radio-journal/internal/controllers/council"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	frontend := os.Getenv("FRONTEND_ADRESS")

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{frontend}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization", "Cookie")

	router.Use(cors.New(config))

	authPath := router.Group("/admin/auth")
	auth := auth.NewAuthController()
	auth.RegisterRoutes(authPath)

	publicFilesPath := router.Group("/files")

	adminFilesPath := router.Group("/admin/files")
	adminFilesPath.Use(auth.AuthMiddleware())

	files := filePaths.NewFilesController()
	files.RegisterRoutes(publicFilesPath, adminFilesPath)

	publicCommentsPath := router.Group("/comments")

	adminCommentsPath := router.Group("/admin/comments")
	adminCommentsPath.Use(auth.AuthMiddleware())

	comments := comments.NewCommentsController()
	comments.RegisterRoutes(publicCommentsPath, adminCommentsPath)

	publicArticlePath := router.Group("/articles")

	adminArticlePath := router.Group("/admin/articles")
	adminArticlePath.Use(auth.AuthMiddleware())

	article := article.NewArticleController(files.GetDeleteHandler(), comments.GetDeleteHandler())
	article.RegisterRoutes(publicArticlePath, adminArticlePath)

	publicEditionPath := router.Group("/editions")

	adminEditionPath := router.Group("/admin/editions")
	adminEditionPath.Use(auth.AuthMiddleware())

	edition := edition.NewEditionController(files.GetDeleteHandler(), article.GetDeleteHandler())
	edition.RegisterRoutes(publicEditionPath, adminEditionPath)

	councilPublicPath := router.Group("/council/members")

	councilAdminPath := router.Group("/admin/council/members")
	councilAdminPath.Use(auth.AuthMiddleware())

	council := council.NewCouncilController(files.GetDeleteHandler())
	council.RegisterRoutes(councilPublicPath, councilAdminPath)

	port := os.Getenv("PORT")
	log.Fatal(router.Run(fmt.Sprintf(":%s", port)))
}
