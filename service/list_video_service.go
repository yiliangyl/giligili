package service

import (
	"giligili/model"
	"giligili/utils"
	"giligili/utils/errmsg"
)

type ListVideoService struct {
}

func (service *ListVideoService) ListVideo(pageSize int, pageNum int) utils.Response {
	var videos []model.Video
	//total := 0

	// todo
	//if err := model.DB.Model(model.Video{}).Count(&total).Error; err != nil {
	//	return utils.ErrResponse(errmsg.LIST_DISPLAY_FALIED, errmsg.GetErrMsg(errmsg.LIST_DISPLAY_FALIED))
	//}

	if err := model.DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&videos).Error; err != nil {
		return utils.ErrResponse(errmsg.LIST_DISPLAY_FALIED, errmsg.GetErrMsg(errmsg.LIST_DISPLAY_FALIED))
	}

	//message := "视频总数为：" + string(total)
	return utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), videos)
}
