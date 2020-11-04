package api

import (
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
