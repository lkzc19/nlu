package mapx

import (
	"fmt"
	"testing"
)

func TestKeys(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(Keys(m))
}

func TestValues(t *testing.T) {
	m := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(Values(m))
}
