package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestIsCurrentDay(t *testing.T) {
	fmt.Println(IsCurrentDay(time.Now()))

	fmt.Println(IsCurrentDay(time.Now().AddDate(1, 0, 0)))

	fmt.Println(IsCurrentDay(time.Now().AddDate(0, 1, 0)))

	fmt.Println(IsCurrentDay(time.Now().AddDate(0, 0, 1)))
}

func TestIsCurrentMonth(t *testing.T) {
	fmt.Println(IsCurrentMonth(time.Now()))

	fmt.Println(IsCurrentMonth(time.Now().AddDate(1, 0, 0)))

	fmt.Println(IsCurrentMonth(time.Now().AddDate(0, 1, 0)))

	fmt.Println(IsCurrentMonth(time.Now().AddDate(0, 1, 1)))
}

func TestIsCurrentYear(t *testing.T) {
	fmt.Println(IsCurrentYear(time.Now()))

	fmt.Println(IsCurrentYear(time.Now().AddDate(1, 0, 0)))

	fmt.Println(IsCurrentYear(time.Now().AddDate(0, 1, 0)))

	fmt.Println(IsCurrentYear(time.Now().AddDate(0, 0, 1)))
}
