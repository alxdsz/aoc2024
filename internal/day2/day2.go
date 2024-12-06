package day2

import (
	"github.com/alxdsz/aoc2024/internal/input"
	"github.com/alxdsz/aoc2024/internal/utils"
	"math"
	"strings"
)

type Solver struct {
	rows []string
}

func NewSolver(inputhPath string) *Solver {
	inpt, _ := input.ReadFile(inputhPath)
	return &Solver{
		rows: inpt.Lines(),
	}
}

func (d *Solver) SolvePart1() int {
	result := 0
	for _, row := range d.rows {
		nums := strings.Split(row, " ")
		if safe := d.isReportSafe(nums, false); safe {
			result += 1
		}
	}
	return result
}

func (d *Solver) SolvePart2() int {
	result := 0
	for _, row := range d.rows {
		nums := strings.Split(row, " ")
		if safe := d.isReportSafe(nums, true); safe {
			result += 1
		}
	}
	return result
}

func (d *Solver) removeElement(slice []string, s int) []string {
	newSlice := make([]string, 0, len(slice)-1)
	for i, _ := range slice {
		if i != s {
			newSlice = append(newSlice, slice[i])
		}
	}
	return newSlice
}

func (d *Solver) isSafe(diff int, shouldAscend bool) bool {
	isAsc := diff < 0
	delta := math.Abs(float64(diff))
	return delta >= 1 && delta <= 3 && isAsc == shouldAscend
}

func (d *Solver) isReportSafe(nums []string, enableDampenerFallback bool) bool {
	isSafe := true
	var shouldAscend bool
	for i, _ := range nums[:len(nums)-1] {
		n := utils.UnsafeAtoi(nums[i], nums[i+1])
		diff := n[0] - n[1]
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
