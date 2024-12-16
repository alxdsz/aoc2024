package day16

import (
	"github.com/alxdsz/aoc2024/internal/input"
	"math"
)

type XY struct {
	x, y int
}

type Board map[XY]rune

func (b Board) get(x, y int) rune {
	if val, exists := b[XY{x, y}]; exists {
		return val
	}
	return '#'
}

type Solver struct {
	board      Board
	start, end XY
}

func NewSolver(inputPath string) *Solver {
	board := make(Board)
	var start, end XY
	for y, row := range input.ReadFileUnsafe(inputPath).Lines() {
		for x, cell := range row {
			board[XY{x, y}] = cell
			if cell == 'S' {
				start = XY{x, y}
			}
			if cell == 'E' {
				end = XY{x, y}
			}
		}
	}
	return &Solver{board, start, end}
}

const (
	E = iota
	S
	W
	N
)

var directions = []XY{
	{1, 0},  // East
	{0, 1},  // South
	{-1, 0}, // West
	{0, -1}, // North
}

type State struct {
	pos XY
	dir int
}

func (d *Solver) findMinScore() int {
	bestScores := make(map[State]int)
	minScore := math.MaxInt

	var dfs func(state State, score int)
	dfs = func(state State, score int) {
		if state.pos == d.end {
			if score < minScore {
				minScore = score
			}
			return
		}

		if prevScore, exists := bestScores[state]; exists && prevScore <= score {
			return
		}
		bestScores[state] = score

		if score >= minScore {
			return
		}

		dir := directions[state.dir]
		newPos := XY{state.pos.x + dir.x, state.pos.y + dir.y}
		if d.board.get(newPos.x, newPos.y) != '#' {
			dfs(State{newPos, state.dir}, score+1)
		}

		dfs(State{state.pos, (state.dir + 3) % 4}, score+1000)
		dfs(State{state.pos, (state.dir + 1) % 4}, score+1000)
	}

	dfs(State{d.start, E}, 0)
	return minScore
}

func (d *Solver) findOptimalTiles(targetScore int) map[XY]bool {
	optimalTiles := make(map[XY]bool)
	visited := make(map[XY]map[State]bool) // track visited states per position

	var dfs func(state State, score int, path []XY)
	dfs = func(state State, score int, path []XY) {
		// If we're over target score, this path isn't optimal
		if score > targetScore {
			return
		}

		// Initialize visited map for this position if needed
		if _, exists := visited[state.pos]; !exists {
			visited[state.pos] = make(map[State]bool)
		}

		// Check for cycles with same state and score
		if visited[state.pos][state] {
			return
		}
		visited[state.pos][state] = true

		currentPath := append(path, state.pos)

		// If we reached the end with exact target score, record all tiles in path
		if state.pos == d.end && score == targetScore {
			for _, pos := range currentPath {
				optimalTiles[pos] = true
			}
			visited[state.pos][state] = false
			return
		}

		// Try moving forward
		dir := directions[state.dir]
		newPos := XY{state.pos.x + dir.x, state.pos.y + dir.y}
		if d.board.get(newPos.x, newPos.y) != '#' {
			dfs(State{newPos, state.dir}, score+1, currentPath)
		}

		// Try rotations
		dfs(State{state.pos, (state.dir + 3) % 4}, score+1000, currentPath)
		dfs(State{state.pos, (state.dir + 1) % 4}, score+1000, currentPath)

		visited[state.pos][state] = false // backtrack
	}

	dfs(State{d.start, E}, 0, nil)
	return optimalTiles
}

func (d *Solver) SolvePart1() int {
	return d.findMinScore()
}

func (d *Solver) SolvePart2() int {
	minScore := d.findMinScore()
	optimalTiles := d.findOptimalTiles(minScore)
	return len(optimalTiles)
}
