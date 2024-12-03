package utils

import "strconv"

func UnsafeAtoi(numbers ...string) []int {
	result := make([]int, len(numbers))
	for i, n := range numbers {
		result[i], _ = strconv.Atoi(n)
	}
	return result
}
