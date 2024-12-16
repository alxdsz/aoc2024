package input

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Input struct {
	lines []string
}

func ReadFile(path string) (*Input, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return &Input{lines: lines}, scanner.Err()
}

func ReadFileUnsafe(path string) *Input {
	inpt, _ := ReadFile(path)
	return inpt
}

func (inp *Input) Lines() []string {
	return inp.lines
}

func (inp *Input) UnzipWhiteSpaceSeparatedLists() (left []int, right []int) {
	for _, line := range inp.Lines() {
		items := strings.Fields(line)
		leftNum, _ := strconv.Atoi(items[0])
		rightNum, _ := strconv.Atoi(items[1])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	return left, right
}

func (inp *Input) As2DArray() [][]string {
	array := make([][]string, len(inp.lines))
	for y, line := range inp.Lines() {
		array[y] = make([]string, len(line))
		for x, letter := range line {
			array[y][x] = string(letter)
		}
	}
	return array
}

func (inp *Input) As2DIntArray() [][]int {
	array := make([][]int, len(inp.lines))
	for y, line := range inp.Lines() {
		array[y] = make([]int, len(line))
		for x, letter := range line {
			num, _ := strconv.Atoi(string(letter))
			array[y][x] = num
		}
	}
	return array
}

func (inp *Input) As2DRuneArray() [][]rune {
	array := make([][]rune, len(inp.lines))
	for y, line := range inp.Lines() {
		array[y] = make([]rune, len(line))
		for x, letter := range line {
			array[y][x] = letter
		}
	}
	return array
}

func (inp *Input) SplitByEmptyLine() [][]string {
	var result [][]string
	var current []string
	for _, line := range inp.Lines() {
		if line == "" {
			result = append(result, current)
			current = nil
		} else {
			current = append(current, line)
		}
	}
	if len(current) > 0 {
		result = append(result, current)
	}
	return result
}
