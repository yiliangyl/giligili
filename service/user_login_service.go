package service

import (
	"giligili/model"
	"giligili/utils"
	"giligili/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 设置session
type UserLoginService struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (service *UserLoginService) setSession(c *gin.Context, user model.User) {
	s := sessions.Default(c)
	s.Clear()
	s.Set("user_id", user.ID)
	s.Save()
}

func (service *UserLoginService) Login(c *gin.Context) utils.Response {
	var user model.User

	if err := model.DB.Where("username = ?", service.Username).First(&user).Error; err != nil {
		return utils.ErrResponse(errmsg.PARAM_INCORRECT, errmsg.GetErrMsg(errmsg.PARAM_INCORRECT))
	}

	if !user.CheckPassword(service.Password) {
		return utils.ErrResponse(errmsg.PARAM_INCORRECT, errmsg.GetErrMsg(errmsg.PARAM_INCORRECT))
	}

	// 设置session
	service.setSession(c, user)

	return utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), user)
}
