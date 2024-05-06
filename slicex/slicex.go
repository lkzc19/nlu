package slicex

// Subtract 差集 s1-s2
func Subtract[T comparable](s1, s2 []T) []T {
	m := make(map[T]bool, len(s1))

	for _, it := range s2 {
		m[it] = true
	}

	s := make([]T, 0)
	for _, it := range s1 {
		if !m[it] {
			s = append(s, it)
		}
	}

	return s
}
