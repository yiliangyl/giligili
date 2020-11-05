package api

import (
	"giligili/service"
	"giligili/utils"
	"giligili/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var service service.UserRegisterService

	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, utils.ErrResponse(errmsg.FAILED, errmsg.GetErrMsg(errmsg.FAILED)))
		return
	}
	response := service.Register()
	c.JSON(http.StatusOK, response)
}

func UserLogin(c *gin.Context) {
	var service service.UserLoginService

	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, utils.ErrResponse(errmsg.FAILED, errmsg.GetErrMsg(errmsg.FAILED)))
		return
	}
	response := service.Login(c)
	c.JSON(http.StatusOK, response)
}

func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	if user == nil {
		c.JSON(http.StatusOK, utils.ErrResponse(errmsg.FAILED, errmsg.GetErrMsg(errmsg.FAILED)))
		return
	}
	response := utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), user)
	c.JSON(http.StatusOK, response)
}

func Logout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	response := utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), nil)
	c.JSON(http.StatusOK, response)
}
