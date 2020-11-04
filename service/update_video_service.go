package service

import (
	"giligili/model"
	"giligili/utils"
	"giligili/utils/errmsg"
)

type UpdateVideoService struct {
	Title string `json:"title"`
	Info  string `json:"info"`
}

func (service *UpdateVideoService) UpdateVideo(id int) utils.Response {
	var video model.Video

	if err := model.DB.Where("id = ?", id).First(&video).Error; err != nil {
		return utils.ErrResponse(errmsg.VIDEO_NON_EXISTED, errmsg.GetErrMsg(errmsg.VIDEO_NON_EXISTED))
	}
	video.Title = service.Title
	video.Info = service.Info

	if err := model.DB.Save(&video).Error; err != nil {
		return utils.ErrResponse(errmsg.VIDEO_SAVE_FAILED, errmsg.GetErrMsg(errmsg.VIDEO_SAVE_FAILED))
	}

	return utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), video)
}
