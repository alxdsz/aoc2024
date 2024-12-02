package day2

import (
	"github.com/alxdsz/aoc2024/internal/input"
	"math"
	"strconv"
	"strings"
)

type Day2Solver struct {
	rows []string
}

func NewDay2Solver(inputhPath string) *Day2Solver {
	inpt, _ := input.ReadFile(inputhPath)
	return &Day2Solver{
		rows: inpt.Lines(),
	}
}

func (d *Day2Solver) SolvePart1() int {
	result := 0
	for _, row := range d.rows {
		nums := strings.Split(row, " ")
		if safe := d.isReportSafe(nums, false); safe {
			result += 1
		}
	}
	return result
}

func unsafeAtoi(a, b string) (int, int) {
	ia, _ := strconv.Atoi(a)
	ib, _ := strconv.Atoi(b)
	return ia, ib
}

func (d *Day2Solver) SolvePart2() int {
	result := 0
	for _, row := range d.rows {
		nums := strings.Split(row, " ")
		if safe := d.isReportSafe(nums, true); safe {
			result += 1
		}
	}
	return result
}

func (d *Day2Solver) removeElement(slice []string, s int) []string {
	newSlice := make([]string, 0, len(slice)-1)
	for i, _ := range slice {
		if i != s {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}

func (d *Day2Solver) isSafe(diff int, shouldAscend bool) bool {
	isAsc := diff < 0
	delta := math.Abs(float64(diff))
	return delta >= 1 && delta <= 3 && isAsc == shouldAscend
}

func (d *Day2Solver) isReportSafe(nums []string, enableDampenerFallback bool) bool {
	isSafe := true
	var shouldAscend bool
	for i, _ := range nums[:len(nums)-1] {
		curr, next := unsafeAtoi(nums[i], nums[i+1])
		diff := curr - next
		if i == 0 {
			shouldAscend = diff < 0
		}
		if !d.isSafe(diff, shouldAscend) {
			if enableDampenerFallback {
				if ok := d.isReportSafe(d.removeElement(nums, i), false); ok {
					isSafe = true
					break
				}
				if ok := d.isReportSafe(d.removeElement(nums, i+1), false); ok {
					isSafe = true
					break
				}
				if i > 0 {
					if ok := d.isReportSafe(d.removeElement(nums, i-1), false); ok {
						isSafe = true
						break
					}
				}
			}
			isSafe = false
			break
		}
	}
	return isSafe
}
