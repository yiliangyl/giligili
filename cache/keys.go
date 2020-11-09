package cache

import (
	"fmt"
	"strconv"
)

// 每日排行榜的key
const DailyRankKey = "rank:daily"

// 此id的video的点击数的key
func VideoViewKey(id uint) string {
	return fmt.Sprintf("view:video-%s", strconv.Itoa(int(id)))
}
