package timex

import (
	"fmt"
	"testing"
	"time"
)

func TestIsCurrentMonth(t *testing.T) {
	fmt.Println(IsCurrentMonth(time.Now()))

	fmt.Println(IsCurrentMonth(time.Now().AddDate(1, 0, 0)))

	fmt.Println(IsCurrentMonth(time.Now().AddDate(0, 1, 0)))
}
