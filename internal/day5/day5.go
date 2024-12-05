package day5

import (
	"math"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Day5Solver struct {
	pages []string
	// slice is fast enough :)
	leftRuleMap  map[string][]string
	rightRuleMap map[string][]string
}

func NewDay5Solver(input [][]string) *Day5Solver {
	leftRuleMap := make(map[string][]string)
	rightRuleMap := make(map[string][]string)
	for _, rule := range input[0] {
		splitRule := strings.Split(rule, "|")
		left, right := splitRule[0], splitRule[1]
		leftRuleMap[left] = append(leftRuleMap[left], right)
		rightRuleMap[right] = append(rightRuleMap[right], left)
	}
	return &Day5Solver{
		leftRuleMap:  leftRuleMap,
		rightRuleMap: rightRuleMap,
		pages:        input[1],
	}
}

func (d *Day5Solver) isRowCorrect(splitPageRow []string) bool {
	for i, page := range splitPageRow {
		pagesOnLeft := splitPageRow[:i]
		pagesOnRight := splitPageRow[i+1:]
		for _, leftPage := range pagesOnLeft {
			if !slices.Contains(d.leftRuleMap[leftPage], page) {
				return false
			}
		}
		for _, rightPage := range pagesOnRight {
			if !slices.Contains(d.rightRuleMap[rightPage], page) {
				return false
			}
		}
	}
	return true
}

func (d *Day5Solver) SolvePart1() int {
	result := 0
	for _, pageRow := range d.pages {
		splitPageRow := strings.Split(pageRow, ",")
		middleIndex := int(math.Ceil(float64(len(splitPageRow) / 2)))
		if d.isRowCorrect(splitPageRow) {
			num, _ := strconv.Atoi(splitPageRow[middleIndex])
			result = result + num
		}
	}

	return result
}

func (d *Day5Solver) SolvePart2() int {
	result := 0
	for _, pageRow := range d.pages {
		splitPageRow := strings.Split(pageRow, ",")
		middleIndex := int(math.Ceil(float64(len(splitPageRow) / 2)))
		if !d.isRowCorrect(splitPageRow) {
			sort.Slice(splitPageRow, func(i, j int) bool {
				return slices.Contains(d.leftRuleMap[splitPageRow[i]], splitPageRow[j])
			})
			num, _ := strconv.Atoi(splitPageRow[middleIndex])
			result = result + num
		}
	}
	return result
}
