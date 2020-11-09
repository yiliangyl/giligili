package service

import (
	"fmt"
	"giligili/cache"
	"giligili/model"
	"giligili/utils"
	"giligili/utils/errmsg"
	"strings"
)

// 每日排行榜服务
type DailyRankService struct {
}

func (service *DailyRankService) GetRankVideos() utils.Response {
	var videoList []model.Video

	// 从redis读取点击前十的视频
	videos, _ := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 9).Result()

	if len(videos) > 0 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(videos, ","))
		if err := model.DB.Where("id in (?)", videos).Order(order).Find(&videoList).Error; err != nil {
			return utils.ErrResponse(errmsg.DATABASE_OPERATE_FAILED, errmsg.GetErrMsg(errmsg.DATABASE_OPERATE_FAILED))
		}
	}

	return utils.CommonResponse(errmsg.SUCCESS, errmsg.GetErrMsg(errmsg.SUCCESS), videoList)
}
