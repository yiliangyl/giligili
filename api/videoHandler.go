package api

import (
	"giligili/service"
	"giligili/utils"
	"giligili/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateVideo(c *gin.Context) {
	var service service.CreateVideoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, utils.ErrResponse(errmsg.FAILED, errmsg.GetErrMsg(errmsg.FAILED)))
		return
	}
	response := service.CreateVideo()
	c.JSON(http.StatusOK, response)
}

func ShowVideo(c *gin.Context) {
	var service service.ShowVideoService
	id, _ := strconv.Atoi(c.Param("id"))
	response := service.ShowVideo(id)
	c.JSON(http.StatusOK, response)
}

func ListVideo(c *gin.Context) {
	var service service.ListVideoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, utils.ErrResponse(errmsg.FAILED, errmsg.GetErrMsg(errmsg.FAILED)))
		return
	}

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	response := service.ListVideo(pageSize, pageNum)
	c.JSON(http.StatusOK, response)
}

func UpdateVideo(c *gin.Context) {
	var service service.UpdateVideoService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusOK, utils.ErrResponse(errmsg.FAILED, errmsg.GetErrMsg(errmsg.FAILED)))
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	response := service.UpdateVideo(id)
	c.JSON(http.StatusOK, response)
}

func DeleteVideo(c *gin.Context) {
	var service service.DeleteVideoService
	id, _ := strconv.Atoi(c.Param("id"))
	response := service.DeleteVideo(id)
	c.JSON(http.StatusOK, response)
}
