package model

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	Title  string `json:"title"`
	Info   string `json:"info"`
	// 视频地址
	URL    string `json:"url"`
	// 封面地址
	Avatar string `json:"avatar"`
}
