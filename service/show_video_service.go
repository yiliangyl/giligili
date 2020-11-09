package service

import (
	"giligili/model"
	"giligili/utils"
	"giligili/utils/errmsg"
)

type ShowVideoService struct {
}

func (service *ShowVideoService) ShowVideo(id int) utils.Response {
	var video model.Video

	if err := model.DB.Where("id = ?", id).First(&video).Error; err != nil {
		return utils.ErrResponse(errmsg.VIDEO_NON_EXISTED, errmsg.GetErrMsg(errmsg.VIDEO_NON_EXISTED))
	}

	video.AddView()

	return utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), video)
}
