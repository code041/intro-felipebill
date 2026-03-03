package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "mundo")
		c.JSON(http.StatusOK, gin.H{"message": "Olá, " + name + "!"})
	})

	return r
}

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}
