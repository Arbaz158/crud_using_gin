package handler

import (
	"fmt"

	"github.com/crud_using_gin/controller"
	"github.com/crud_using_gin/entity"
	"github.com/gin-gonic/gin"
)

type VideoHandler interface {
	Save(*gin.Context) []entity.Video
}

type videoHandler struct {
	control controller.VideoController
}

func (vd *videoHandler) Save(ctx *gin.Context) []entity.Video {
	var video []entity.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		fmt.Println("error in getting data from body :", err)
	}
	fmt.Println("body data :", video)
	vd.control.SaveVideo(video)
	return video
}

func New(control controller.VideoController) VideoHandler {
	return &videoHandler{
		control: control,
	}
}
