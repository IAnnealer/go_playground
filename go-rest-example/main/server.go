package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iannealer/go_playground/go-rest-sample/controller"
	"github.com/iannealer/go_playground/go-rest-sample/service"
	"net/http"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	server.Run(":8080")
}
