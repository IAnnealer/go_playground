package main

import (
	"go_playground/go-crud/controllers"
	"go_playground/go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// router.GET("/", controllers.PostsCreate)
	router.POST("/posts", controllers.PostsCreate)

	router.GET("/posts", controllers.PostsIndex)

	router.GET("/posts/:id", controllers.PostsShow)

	router.PUT("/posts/:id", controllers.PostsUpdate)

	router.DELETE("/posts/:id", controllers.PostsDelete)

	router.Run()
}
