package day3

import (
	"github.com/alxdsz/aoc2024/internal/input"
	"github.com/alxdsz/aoc2024/internal/utils"
	"regexp"
	"strings"
)

type Day3Solver struct {
	inputRows []string
}

func NewDay3Solver(inputPath string) *Day3Solver {
	inpt, _ := input.ReadFile(inputPath)
	return &Day3Solver{
		inputRows: inpt.Lines(),
	}
}

func (d *Day3Solver) SolvePart1() int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	rgx, _ := regexp.Compile(pattern)
	result := 0
	for _, row := range d.inputRows {
		for _, matches := range rgx.FindAllStringSubmatch(row, -1) {
			nums := utils.UnsafeAtoi(matches[1], matches[2])
			result += nums[0] * nums[1]
		}
	}
	return result
}

func (d *Day3Solver) SolvePart2() int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`
	rgx, _ := regexp.Compile(pattern)
	result := 0
	instructionEnabled := true
	for _, row := range d.inputRows {
		for _, matches := range rgx.FindAllStringSubmatch(row, -1) {
			if strings.HasPrefix(matches[0], "do") {
				instructionEnabled = matches[0] == "do()"
			} else if instructionEnabled {
				nums := utils.UnsafeAtoi(matches[1], matches[2])
				result += nums[0] * nums[1]
			}
		}
	}
	return result
}
