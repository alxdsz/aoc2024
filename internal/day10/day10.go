package day10

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/input"
)

type BoardPosition struct {
	x, y int
}

// :D
func (p BoardPosition) getAllNextPositions(b Board) []BoardPosition {
	var result []BoardPosition
	right := BoardPosition{x: p.x + 1, y: p.y}
	if b.isValidStep(p, right) {
		result = append(result, right)
	}
	left := BoardPosition{x: p.x - 1, y: p.y}
	if b.isValidStep(p, left) {
		result = append(result, left)
	}
	down := BoardPosition{x: p.x, y: p.y + 1}
	if b.isValidStep(p, down) {
		result = append(result, down)
	}
	up := BoardPosition{x: p.x, y: p.y - 1}
	if b.isValidStep(p, up) {
		result = append(result, up)
	}
	return result
}

type Board [][]int

func (b Board) getValue(p BoardPosition) int {
	return b[p.y][p.x]
}

func (b Board) isPositionWithinBoard(position BoardPosition) bool {
	return len(b) > 0 && position.x >= 0 && position.x < len(b[0]) && position.y >= 0 && position.y < len(b)
}

func (b *Board) isValidStep(prev BoardPosition, next BoardPosition) bool {
	return b.isPositionWithinBoard(next) && b.getValue(next)-b.getValue(prev) == 1
}

type Solver struct {
	board Board
}

func NewSolver(inputPath string) *Solver {
	inpt, _ := input.ReadFile(inputPath)
	board := inpt.As2DIntArray()
	return &Solver{
		board: board,
	}
}

func (s *Solver) findStartingPositions() []BoardPosition {
	var startingPositions []BoardPosition
	for y, row := range s.board {
		for x, cell := range row {
			if cell == 0 {
				startingPositions = append(startingPositions, BoardPosition{x, y})
			}
		}
	}
	return startingPositions
}

func (s *Solver) scoreTrailhead(startPosition BoardPosition) int {
	visited := make(map[BoardPosition]bool)
	trailsReachingNine := make(map[BoardPosition]bool)
	var dfs func(pos BoardPosition)
	dfs = func(pos BoardPosition) {
		if visited[pos] {
			return
		}
		visited[pos] = true
		if s.board.getValue(pos) == 9 {
			trailsReachingNine[pos] = true
			return
		}
		for _, next := range pos.getAllNextPositions(s.board) {
			dfs(next)
		}
	}
	dfs(startPosition)
	return len(trailsReachingNine)
}

func (s *Solver) SolvePart1() int {
	total := 0
	for _, start := range s.findStartingPositions() {
		total += s.scoreTrailhead(start)
	}
	return total
}

func getPathKey(path []BoardPosition) string {
	result := ""
	for _, pos := range path {
		result += fmt.Sprintf("(%d,%d);", pos.x, pos.y)
	}
	return result
}

func (s *Solver) countUniquePaths(start BoardPosition) int {
	uniquePaths := make(map[string]bool)
	currentPath := []BoardPosition{start}

	var dfs func(pos BoardPosition, path []BoardPosition)
	dfs = func(pos BoardPosition, path []BoardPosition) {
		if s.board.getValue(pos) == 9 {
			uniquePaths[getPathKey(path)] = true
			return
		}
		for _, next := range pos.getAllNextPositions(s.board) {
			// MAKE A FREAKING COPY !
			newPath := make([]BoardPosition, len(path))
			copy(newPath, path)
			newPath = append(newPath, next)

			dfs(next, newPath)
		}
	}

	dfs(start, currentPath)
	return len(uniquePaths)
}

func (s *Solver) SolvePart2() int {
	total := 0
	for _, start := range s.findStartingPositions() {
		pathCount := s.countUniquePaths(start)
		total += pathCount
	}
	return total
}
