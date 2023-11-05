package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	log.Fatal(router.Run())
}
