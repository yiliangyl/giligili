package server

import (
	"giligili/api"
	"giligili/middleware"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Router() {
	AppMode := viper.GetString("AppMode")
	gin.SetMode(AppMode)
	r := gin.Default()

	secret := viper.GetString("SessionSecret")
	r.Use(middleware.Session(secret))
	r.Use(middleware.CurrentUser())

	v1 := r.Group("/api/v1")
	{
		v1.POST("/ping", api.TestHandler)

		user := v1.Group("/user")
		{
			user.POST("/register", api.UserRegister)
			user.POST("/login", api.UserLogin)
			auth := user.Group("/")
			auth.Use(middleware.AuthRequired())
			{
				auth.GET("/me", api.UserMe)
				auth.DELETE("/logout", api.Logout)
			}
		}

		video := v1.Group("/video")
		{
			video.GET("/show/:id", api.ShowVideo)
			video.GET("/list", api.ListVideo)
			video.GET("/rank", api.VideoRank)
			auth := video.Group("/")
			auth.Use(middleware.AuthRequired())
			{
				auth.POST("/create", api.CreateVideo)
				auth.PUT("/update/:id", api.UpdateVideo)
				auth.DELETE("/delete/:id", api.DeleteVideo)
			}
		}
	}

	HttpPort := viper.GetString("HttpPort")
	r.Run(HttpPort)
}
