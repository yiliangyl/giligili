package middleware

import (
	"giligili/model"
	"giligili/utils"
	"giligili/utils/errmsg"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取当前用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			if user, err := model.GetUserById(uid); err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, _ := c.Get("user"); user != nil {
			if _, ok := user.(*model.User); ok {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusOK, utils.CommonResponse(errmsg.USER_NON_LOGIN, errmsg.GetErrMsg(errmsg.USER_NON_LOGIN), nil))
		c.Abort()
		return
	}
}
