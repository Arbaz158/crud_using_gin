package controller

import "github.com/crud_using_gin/entity"

type VideoController interface {
	SaveVideo([]entity.Video) error
	GetAllVideo() []entity.Video
	UpdateVideo([]entity.Video) error
	DeleteVideo(entity.Video) error
}

type videoController struct {
	videos []entity.Video
}

func New() VideoController {
	return &videoController{}
}
