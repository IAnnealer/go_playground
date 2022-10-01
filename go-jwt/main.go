package main

import (
	"github.com/gin-gonic/gin"
	"go_playground/go-jwt/initializers"
	"net/http"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
