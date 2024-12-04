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

func (inp *Input) AsArray() [][]string {
	array := make([][]string, len(inp.lines))
	for y, line := range inp.Lines() {
		array[y] = make([]string, len(line))
		for x, letter := range line {
			array[y][x] = string(letter)
		}
	}
	return array
}
