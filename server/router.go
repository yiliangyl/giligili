package server

import (
	"giligili/api"
	"giligili/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Router()  {
	AppMode := viper.GetString("AppMode")
	gin.SetMode(AppMode)
	r := gin.Default()

	secret := viper.GetString("SessionSecret")
	r.Use(middleware.Session(secret))

	v1 := r.Group("/api/v1")
	{
		v1.POST("/ping", api.TestHandler)

		user := v1.Group("/user")
		{
			user.POST("/register", api.UserRegister)
			user.POST("/login", api.UserLogin)
		}

		video := v1.Group("/video")
		{
			video.POST("/create", api.CreateVideo)
			video.GET("/show/:id", api.ShowVideo)
			video.GET("/list", api.ListVideo)
			video.PUT("/update/:id", api.UpdateVideo)
			video.DELETE("/delete/:id", api.DeleteVideo)
		}
	}

	HttpPort := viper.GetString("HttpPort")
	r.Run(HttpPort)
}
