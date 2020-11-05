package api

import (
	"giligili/model"
	"giligili/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func TestHandler(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		utils.CommonResponse(200, "pong", nil),
	)
}

func CurrentUser(c *gin.Context) *model.User {
	if user, _ := c.Get("user"); user != nil {
		if u, ok := user.(*model.User); ok {
			return u
		}
	}
	return nil
}
