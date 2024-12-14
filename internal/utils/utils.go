package utils

import "strconv"

func UnsafeAtoi(numbers ...string) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i], _ = strconv.Atoi(n)
	}
	return result
}

func SlicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[T]int)
	for _, v := range a {
		counts[v]++
	}

	for _, v := range b {
		counts[v]--
		if counts[v] < 0 {
			return false
		}
	}

	return true
}
