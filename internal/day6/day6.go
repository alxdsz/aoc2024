package day6

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/input"
	"image/color"
)

type Solver struct {
	board     [][]string
	starPoint []int
}

func findStartPoint(board [][]string) []int {
	for cy, row := range board {
		for cx, cell := range row {
			if cell != "#" && cell != "." {
				return []int{cx, cy}
			}
		}
	}
	panic(fmt.Errorf("not found"))
}

func NewSolver(inputPath string) *Solver {
	inp, _ := input.ReadFile(inputPath)
	board := inp.As2DArray()
	return &Solver{
		board:     board,
		starPoint: findStartPoint(board),
	}
}

func cell2Color(cell string) color.Color {
	switch cell {
	case ".":
		return color.NRGBA{R: 255, G: 255, B: 0, A: 255}
	case "#":
		return color.NRGBA{R: 0, G: 0, B: 255, A: 255}
	case "X":
		return color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	case "0":
		return color.NRGBA{R: 255, G: 0, B: 255, A: 255}
	default:
		return color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	}
}

const (
	UP    = "^"
	RIGHT = ">"
	DOWN  = "v"
	LEFT  = "<"
)

var dirMap = map[string]string{
	UP:    RIGHT,
	RIGHT: DOWN,
	DOWN:  LEFT,
	LEFT:  UP,
}

func (d *Solver) doStep(cx, cy int, cdirection string) (nx, ny int, ndirection string, newSpotVisited bool, leftBoard bool) {
	nx, ny = cx, cy
	ndirection = cdirection
	switch cdirection {
	case UP:
		ny--
		break
	case DOWN:
		ny++
		break
	case RIGHT:
		nx++
		break
	case LEFT:
		nx--
		break
	}
	if ny >= 0 && ny < len(d.board) && nx >= 0 && nx < len(d.board[0]) {
		if d.board[ny][nx] == "#" {
			ndirection = dirMap[cdirection]
			return cx, cy, ndirection, false, false
		} else {
			newSpotVisited = d.board[ny][nx] != "X" && d.board[ny][nx] != "0"
			if d.board[cy][cx] != "0" {
				d.board[cy][cx] = "X"
			}
			if d.board[ny][nx] != "0" {
				d.board[ny][nx] = ndirection
			}
			return nx, ny, ndirection, newSpotVisited, false
		}

	} else {
		return -1, -1, ndirection, false, true
	}
}

func (d *Solver) SolvePart1() int {
	cx, cy := d.starPoint[0], d.starPoint[1]
	currentDirection := UP
	leftBoard := false
	result := 0
	for !leftBoard {
		// TODO remove visualization or make possible to enable from terminal
		//if math.Mod(float64(result), 5000) == 0 {
		//	vis.Visualize2dArrayInTerminal(&d.board, cell2Color)
		//}
		var newSpotVisited bool
		cx, cy, currentDirection, newSpotVisited, leftBoard = d.doStep(cx, cy, currentDirection)
		if newSpotVisited {
			result++
		}
	}

	return result
}

// TODO ugly af but works
func (d *Solver) isGoodObstaclePosition(cx, cy int, cDirection string) bool {
	boardCopy := make([][]string, len(d.board))
	for i := range d.board {
		boardCopy[i] = make([]string, len(d.board[0]))
		copy(boardCopy[i], d.board[i])
	}

	visited := make(map[string]bool)

	x, y := cx, cy
	dir := cDirection

	for {
		stateKey := fmt.Sprintf("%d:%d:%s", x, y, dir)
		if visited[stateKey] {
			return true // Found a cycle
		}
		visited[stateKey] = true

		nx, ny := x, y
		switch dir {
		case UP:
			ny--
		case DOWN:
			ny++
		case RIGHT:
			nx++
		case LEFT:
			nx--
		}

		if nx < 0 || nx >= len(boardCopy[0]) || ny < 0 || ny >= len(boardCopy) {
			return false
		}

		if boardCopy[ny][nx] == "#" {
			dir = dirMap[dir]
			continue
		}

		x, y = nx, ny
	}
}

func (d *Solver) SolvePart2() int {
	startX, startY := d.starPoint[0], d.starPoint[1]
	startDir := d.board[startY][startX]
	result := 0

	for y := 0; y < len(d.board); y++ {
		for x := 0; x < len(d.board[0]); x++ {
			if d.board[y][x] != "." || (x == startX && y == startY) {
				continue
			}

			d.board[y][x] = "#"

			if d.isGoodObstaclePosition(startX, startY, startDir) {
				result++
			}

			d.board[y][x] = "."
		}
	}

	return result
}
