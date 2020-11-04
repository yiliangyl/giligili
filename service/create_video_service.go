package service

import (
	"giligili/model"
	"giligili/utils"
	"giligili/utils/errmsg"
)

type CreateVideoService struct {
	Title  string `json:"title" form:"title"`
	Info   string `json:"info" form:"info"`
	URL    string `json:"url" form:"url"`
	Avatar string `json:"avatar" form:"avatar"`
}

func (service *CreateVideoService) CreateVideo() utils.Response {
	video := model.Video{
		Title: service.Title,
		Info: service.Info,
		URL: service.URL,
		Avatar: service.Avatar,
	}

	if err := model.DB.Create(&video).Error; err != nil {
		return utils.ErrResponse(errmsg.VIDEO_SAVE_FAILED, errmsg.GetErrMsg(errmsg.VIDEO_SAVE_FAILED))
	}

	return utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), video)
}
