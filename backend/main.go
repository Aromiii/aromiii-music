package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(Authenticate)
	r.GET("/", HandleRoot)
	r.Run()
}
