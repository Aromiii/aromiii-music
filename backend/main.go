package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Public routes
	r.GET("/stream/:id", HandleStream)

	r.Use(Authenticate)

	// Routes behind auth
	r.GET("/", HandleRoot)

	r.Run()
}
