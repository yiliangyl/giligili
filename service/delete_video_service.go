package service

import (
	"giligili/model"
	"giligili/utils"
	"giligili/utils/errmsg"
)

type DeleteVideoService struct {
}

func (service *DeleteVideoService) DeleteVideo(id int) utils.Response {
	if err := model.DB.Where("id = ?", id).Delete(&model.Video{}).Error; err != nil {
		return utils.ErrResponse(errmsg.VIDEO_DELETE_FAILED, errmsg.GetErrMsg(errmsg.VIDEO_DELETE_FAILED))
	}

	return utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), nil)
}
