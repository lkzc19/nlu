package testx

import (
	"time"
)

// MeasureTimeMillis 计算方法执行时间 单位: 毫秒
func MeasureTimeMillis(f func()) int64 {
	start := time.Now()
	f()
	end := time.Now()
	elapsed := end.Sub(start)
	return elapsed.Milliseconds()
}

// MeasureTimeSec 计算方法执行时间 单位: 秒
func MeasureTimeSec(f func()) float64 {
	start := time.Now()
	f()
	end := time.Now()
	elapsed := end.Sub(start)
	return elapsed.Seconds()
}
