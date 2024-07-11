package testx

import (
	"fmt"
	"testing"
	"time"
)

func TestMeasureTimeMillis(t *testing.T) {
	millis := MeasureTimeMillis(func() {
		time.Sleep(2000 * time.Millisecond)
	})
	fmt.Printf("执行了: %d 毫秒\n", millis)
}

func TestMeasureTimeSec(t *testing.T) {
	sec := MeasureTimeSec(func() {
		time.Sleep(2 * time.Second)
	})
	fmt.Printf("执行了: %f 秒\n", sec)
}
