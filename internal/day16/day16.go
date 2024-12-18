package day16

import (
	"container/heap"
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

func (d *Solver) dijkstra(start XY, startDir int, isForward bool) map[State]int {
	scores := make(map[State]int)
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	startState := State{start, startDir}
	scores[startState] = 0
	heap.Push(&pq, &Item{state: startState, score: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item)
		state, score := current.state, current.score

		if score > scores[state] {
			continue
		}

		moveDir := state.dir
		if !isForward {
			moveDir = (state.dir + 2) % 4
		}

		dir := directions[moveDir]
		newPos := XY{state.pos.x + dir.x, state.pos.y + dir.y}
		if d.board.get(newPos.x, newPos.y) != '#' {
			newState := State{newPos, state.dir}
			newScore := score + 1
			if bestScore, exists := scores[newState]; !exists || newScore < bestScore {
				scores[newState] = newScore
				heap.Push(&pq, &Item{state: newState, score: newScore})
			}
		}

		for _, newDir := range []int{(state.dir + 3) % 4, (state.dir + 1) % 4} {
			newState := State{state.pos, newDir}
			newScore := score + 1000
			if bestScore, exists := scores[newState]; !exists || newScore < bestScore {
				scores[newState] = newScore
				heap.Push(&pq, &Item{state: newState, score: newScore})
			}
		}
	}

	return scores
}

func (d *Solver) findOptimalTiles() (int, int) {
	forwardScores := d.dijkstra(d.start, E, true)
	backwardScores := make(map[State]int)
	minScore := math.MaxInt

	for dir := 0; dir < 4; dir++ {
		scores := d.dijkstra(d.end, dir, false)
		for state, score := range scores {
			if existing, exists := backwardScores[state]; !exists || score < existing {
				backwardScores[state] = score
			}
		}
	}

	// Find minimum total score
	for state, fScore := range forwardScores {
		if state.pos == d.end {
			if fScore < minScore {
				minScore = fScore
			}
		}
	}

	optimalTiles := make(map[XY]bool)
	for state, fScore := range forwardScores {
		if bScore, exists := backwardScores[state]; exists {
			if fScore+bScore == minScore {
				optimalTiles[state.pos] = true
			}
		}
	}

	return minScore, len(optimalTiles)
}

func (d *Solver) SolvePart1() int {
	score, _ := d.findOptimalTiles()
	return score
}

func (d *Solver) SolvePart2() int {
	_, count := d.findOptimalTiles()
	return count
}
