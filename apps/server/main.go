package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.Static("/", "../web/build")

	// Start and run the server
	router.Run(":8080")
}
