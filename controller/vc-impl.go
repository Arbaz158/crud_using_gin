package controller

import (
	"github.com/crud_using_gin/config"
	"github.com/crud_using_gin/entity"
)

var conf = config.DbConnection()

func (vd *videoController) SaveVideo(video []entity.Video) error {
	conf.Create(&video)
	return nil
}

func (vd *videoController) GetAllVideo() []entity.Video {
	var result []entity.Video
	conf.Find(&result)
	return result
}

func (vd *videoController) UpdateVideo(video []entity.Video) error {
	conf.Save(&video)
	return nil
}

func (vd *videoController) DeleteVideo(video entity.Video) error {
	conf.Where("url=?", video.Url).Delete(&video)
	return nil
}
