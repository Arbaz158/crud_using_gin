package main

import (
	"github.com/crud_using_gin/controller"
	"github.com/crud_using_gin/handler"
	"github.com/gin-gonic/gin"
)

var (
	VideoController controller.VideoController = controller.New()
	VideoHandler    handler.VideoHandler       = handler.New(VideoController)
)

func main() {
	server := gin.Default()
	server.POST("/save-video", func(ctx *gin.Context) {
		ctx.JSON(200, VideoHandler.Save(ctx))
	})

	server.Run(":8081")
}
