package api

import (
	"giligili/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VideoRank(c *gin.Context) {
	var service service.DailyRankService
	response := service.GetRankVideos()
	c.JSON(http.StatusOK, response)
}
