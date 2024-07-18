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

func TestIntersect(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{3, 4, 5}

	fmt.Println(Intersect(s1, s2))
}

func TestUnion(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{3, 4, 5}

	fmt.Println(Union(s1, s2))
}

func TestDistinct(t *testing.T) {
	s := []int{1, 2, 3, 3}

	fmt.Println(Distinct(s))
}

func TestContains(t *testing.T) {
	var s []string
	fmt.Println(Contains(s, "foo"))
	s = append(s, "foo")
	fmt.Println(Contains(s, "foo"))
}

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(Reverse(s))
	s = []int{1, 2, 3, 4, 5}
	fmt.Println(Reverse(s))
	s = []int{}
	fmt.Println(Reverse(s))
}
