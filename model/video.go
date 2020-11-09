package model

import (
	"fmt"
	"giligili/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Video struct {
	gorm.Model
	Title string `json:"title"`
	Info  string `json:"info"`
	// 视频地址
	URL string `json:"url"`
	// 封面地址
	Avatar string `json:"avatar"`
	// 视频播放量
	View int `json:"view"`
}

func (video *Video) AddView() {
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))

	view, _ := strconv.Atoi(cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Val())
	video.View = view
	if err := DB.Save(video).Error; err != nil {
		fmt.Printf("数据库操作失败，err: %v", err)
		return
	}

	// 可查看李文周的博客
	// 增加排行榜点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))
}
