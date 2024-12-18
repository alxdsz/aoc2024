package day11

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/input"
	"strconv"
	"strings"
)

type Solver struct {
	stones []string
	cache  map[string]int
}

func NewSolver(inputPath string) *Solver {
	inpt, _ := input.ReadFile(inputPath)
	stones := strings.Split(inpt.Lines()[0], " ")
	return &Solver{
		stones: stones,
		cache:  make(map[string]int),
	}
}

func (s *Solver) transformStone(stone int, iteration int, maxIter int) int {
	if iteration == maxIter {
		return 1
	}

	key := fmt.Sprintf("%d%d", stone, iteration)
	if count, exists := s.cache[key]; exists {
		return count
	}

	var result int
	stoneAsString := strconv.Itoa(stone)
	if stone == 0 {
		result = s.transformStone(1, iteration+1, maxIter)
	} else if len(stoneAsString)%2 == 0 {
		mid := len(stoneAsString) / 2
		left, _ := strconv.Atoi(stoneAsString[:mid])
		right, _ := strconv.Atoi(stoneAsString[mid:])
		result = s.transformStone(left, iteration+1, maxIter) + s.transformStone(right, iteration+1, maxIter)
	} else {
		result = s.transformStone(stone*2024, iteration+1, maxIter)
	}

	s.cache[key] = result
	return result
}

func (s *Solver) SolvePart1() string {
	total := 0
	for _, stoneStr := range s.stones {
		stone, _ := strconv.Atoi(stoneStr)
		total += s.transformStone(stone, 0, 25)
	}
	return strconv.Itoa(total)
}

func (s *Solver) SolvePart2() string {
	total := 0
	for _, stoneStr := range s.stones {
		stone, _ := strconv.Atoi(stoneStr)
		total += s.transformStone(stone, 0, 75)
	}
	return strconv.Itoa(total)
}
