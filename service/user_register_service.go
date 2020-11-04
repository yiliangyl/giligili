package service

import (
	"giligili/model"
	"giligili/utils"
	"giligili/utils/errmsg"
)

type UserRegisterService struct {
	Username        string `json:"username" form:"username"`
	Nickname        string `json:"nickname" form:"nickname"`
	Password        string `json:"password" form:"password"`
	PasswordConfirm string `json:"password_confirm" form:"password_confirm"`
	Avatar          string `json:"avatar" form:"avatar"`
}

func (service *UserRegisterService) valid() utils.Response {
	if service.Password != service.PasswordConfirm {
		return utils.ErrResponse(errmsg.PASSWORD_INCONSISTENCY, errmsg.GetErrMsg(errmsg.PASSWORD_INCONSISTENCY))
	}

	count := 0
	model.DB.Model(&model.User{}).Where("username = ?", service.Username).Count(&count)
	if count > 0 {
		return utils.ErrResponse(errmsg.USER_IS_EXISTED, errmsg.GetErrMsg(errmsg.USER_IS_EXISTED))
	}

	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return utils.ErrResponse(errmsg.NAME_IS_EXISTED, errmsg.GetErrMsg(errmsg.NAME_IS_EXISTED))
	}
	return utils.Response{}
}

func (service *UserRegisterService) Register() utils.Response {
	user := model.User{
		Username: service.Username,
		Nickname: service.Nickname,
		Avatar:   service.Avatar,
	}

	if errResponse := service.valid(); errResponse.Code != 0 {
		return errResponse
	}

	if err := user.SetPassword(service.Password); err != nil {
		return utils.ErrResponse(errmsg.PASSWORD_ENCRYPTION_ERROR, errmsg.GetErrMsg(errmsg.PASSWORD_ENCRYPTION_ERROR))
	}

	if err := model.DB.Create(&user).Error; err != nil {
		return utils.ErrResponse(errmsg.USER_REGISTER_FAILED, errmsg.GetErrMsg(errmsg.USER_REGISTER_FAILED))
	}

	return utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), user)
}
