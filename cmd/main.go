package main

import (
	"fmt"
	"log"
	"os"
	"urfu-radio-journal/internal/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	authPath := router.Group("/admin/auth")
	auth := controllers.NewAuthController()
	auth.RegisterRoutes(authPath)

	publicFilesPath := router.Group("/public/files")

	adminFilesPath := router.Group("/admin/files")
	adminFilesPath.Use(auth.SessionsHandler())
	adminFilesPath.Use(auth.AuthMiddleware())

	files := controllers.NewFilesController()
	files.RegisterRoutes(publicFilesPath, adminFilesPath)

	publicCommentsPath := router.Group("/comments")

	adminCommentsPath := router.Group("/admin/comments")
	adminCommentsPath.Use(auth.SessionsHandler())
	adminCommentsPath.Use(auth.AuthMiddleware())

	comments := controllers.NewCommentsController()
	comments.RegisterRoutes(publicCommentsPath, adminCommentsPath)

	publicArticlePath := router.Group("/articles")

	adminArticlePath := router.Group("/admin/articles")
	adminArticlePath.Use(auth.SessionsHandler())
	adminArticlePath.Use(auth.AuthMiddleware())

	article := controllers.NewArticleController(files.GetDeleteHandler(), comments.GetDeleteHandler())
	article.RegisterRoutes(publicArticlePath, adminArticlePath)

	publicEditionPath := router.Group("/editions")

	adminEditionPath := router.Group("/admin/editions")
	adminEditionPath.Use(auth.SessionsHandler())
	adminEditionPath.Use(auth.AuthMiddleware())

	edition := controllers.NewEditionController(files.GetDeleteHandler(), article.GetDeleteHandler())
	edition.RegisterRoutes(publicEditionPath, adminEditionPath)

	log.Fatal(router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
