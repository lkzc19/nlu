package slicex

// Subtract 差集
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

// Intersect 交集
func Intersect[T comparable](s1, s2 []T) []T {
	m := make(map[T]bool, len(s1))

	for _, it := range s2 {
		m[it] = true
	}

	s := make([]T, 0)
	for _, it := range s1 {
		if m[it] {
			s = append(s, it)
		}
	}

	return s
}

// Union 并集
func Union[T comparable](s1, s2 []T) []T {
	s := append(s1, s2...)
	return Distinct(s)
}

// Distinct 去重
func Distinct[T comparable](s []T) []T {
	m := make(map[T]bool)

	newS := make([]T, 0)
	for _, it := range s {
		if !m[it] {
			m[it] = true
			newS = append(newS, it)
		}
	}

	return newS
}
