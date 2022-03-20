package main

import (
	"github.com/Galen-liu/project-based-learning/go/url_shortener/controller/url_shortener"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/api/v1/short-urls", url_shortener.CreateShortenedUrl)
	router.GET("/shorten-url/:id", url_shortener.Redirect2RealUrl)

	router.Run("localhost:8080")
}
