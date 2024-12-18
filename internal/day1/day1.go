package day1

import (
	"github.com/alxdsz/aoc2024/internal/input"
	"math"
	"sort"
	"strconv"
)

type Solver struct {
	left  []int
	right []int
}

func NewSolver(inputPath string) *Solver {
	inp, _ := input.ReadFile(inputPath)
	left, right := inp.UnzipWhiteSpaceSeparatedLists()
	sort.Ints(left)
	sort.Ints(right)
	return &Solver{
		left:  left,
		right: right,
	}
}

func (d *Solver) SolvePart1() string {
	result := 0
	for i, leftNumber := range d.left {
		rightNumber := d.right[i]
		result = result + int(math.Abs(float64(rightNumber-leftNumber)))
	}
	return strconv.Itoa(result)
}

func (d *Solver) SolvePart2() string {
	freqMap := make(map[int]int)
	for _, rightNumber := range d.right {
		freqMap[rightNumber]++
	}
	result := 0
	for _, leftNumber := range d.left {
		result = result + freqMap[leftNumber]*leftNumber
	}
	return strconv.Itoa(result)
}
