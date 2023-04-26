package main

import (
	"commercial-propfloor-users/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	routes.Routes(router)

	log.Fatal(router.Run(":" + port))
}
