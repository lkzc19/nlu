package timex

import (
	"time"
)

func IsCurrentDay(t time.Time) bool {
	year, month, day := time.Now().Date()
	return t.Year() == year && t.Month() == month && t.Day() == day
}

func IsCurrentMonth(t time.Time) bool {
	year, month, _ := time.Now().Date()
	return t.Year() == year && t.Month() == month
}

func IsCurrentYear(t time.Time) bool {
	year, _, _ := time.Now().Date()
	return t.Year() == year
}
