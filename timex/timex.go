package timex

import (
	"time"
)

func IsCurrentMonth(t time.Time) bool {
	now := time.Now()
	year, month, _ := now.Date()
	// 获取当月的第一天
	firstDay := time.Date(year, month, 1, 0, 0, 0, 0, now.Location())
	// 获取当月的最后一天
	lastDay := firstDay.AddDate(0, 1, -1)
	// 判断当前时间是否在当月
	return t.After(firstDay) && t.Before(lastDay)
}
