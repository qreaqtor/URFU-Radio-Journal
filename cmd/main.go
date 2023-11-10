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

	editionPath := router.Group("/edition")
	edition := controllers.NewEditionController()
	edition.RegisterRoutes(editionPath)

	log.Fatal(router.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
