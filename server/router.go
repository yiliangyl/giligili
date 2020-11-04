package server

import (
	"giligili/api"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Router()  {
	AppMode := viper.GetString("AppMode")
	gin.SetMode(AppMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/ping", api.TestHandler)
		v1.POST("/register", api.UserRegister)
		v1.POST("/login", api.UserLogin)
	}

	HttpPort := viper.GetString("HttpPort")
	r.Run(HttpPort)
}
