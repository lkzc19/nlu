package slicex

import (
	"fmt"
	"testing"
)

func TestSubtract(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{3, 4, 5}

	fmt.Println(Subtract(s1, s2))
}
