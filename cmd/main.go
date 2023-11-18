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

	publicEditionPath := router.Group("/editions")
	adminEditionPath := router.Group("/admin/editions")
	adminEditionPath.Use(auth.SessionsHandler())
	adminEditionPath.Use(auth.AuthMiddleware())
	edition := controllers.NewEditionController()
	edition.RegisterRoutes(publicEditionPath, adminEditionPath)

	log.Fatal(router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
