package day15

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/input"
	"github.com/alxdsz/aoc2024/internal/vis"
	"image/color"
	"strconv"
)

type Solver struct {
	board     *[][]rune
	bigBoard  *[][]rune
	movements string
}

func enlargeSymbol(symbol rune) []rune {
	switch symbol {
	case '#':
		return []rune{'#', '#'}
	case '@':
		return []rune{'@', '.'}
	case 'O':
		return []rune{'[', ']'}
	default:
		return []rune{'.', '.'}
	}
}

func NewSolver(inputPath string) *Solver {
	inpt, _ := input.ReadFile(inputPath)
	split := inpt.SplitByEmptyLine()
	boardLines, movementLines := split[0], split[1]

	var board [][]rune
	var bigBoard [][]rune

	for y, boardLine := range boardLines {
		board = append(board, []rune{})
		bigBoard = append(bigBoard, []rune{})
		for _, symbol := range boardLine {
			board[y] = append(board[y], symbol)
			bigBoard[y] = append(bigBoard[y], enlargeSymbol(symbol)...)
		}
	}

	movements := ""
	for _, movementLine := range movementLines {
		movements += movementLine
	}

	return &Solver{&board, &bigBoard, movements}
}

func (s *Solver) SolvePart1() string {
	x, y := findRobotPosition(s.board)
	for _, movement := range s.movements {
		visualizeBoard(s.board)
		dx, dy := s.getDirection(movement)
		moved := s.tryMoveElement(x, y, dx, dy, s.board)
		if moved {
			x, y = x+dx, y+dy
		}
	}
	result := 0
	for y, row := range *s.board {
		for x, cell := range row {
			if cell == 'O' {
				result += 100*y + x
			}
		}
	}
	return strconv.Itoa(result)
}

func (s *Solver) SolvePart2() string {
	x, y := findRobotPosition(s.bigBoard)
	for _, movement := range s.movements {
		//visualizeBoard(s.bigBoard)
		dx, dy := s.getDirection(movement)
		moved := s.tryMoveElement2(x, y, dx, dy, s.bigBoard)
		if moved {
			x, y = x+dx, y+dy
		}
	}
	result := 0
	for y, row := range *s.bigBoard {
		for x, cell := range row {
			if cell == '[' {
				result += 100*y + x
			}
		}
	}
	return strconv.Itoa(result)
}

func (s *Solver) getDirection(mvmnt rune) (x int, y int) {
	switch mvmnt {
	case '<':
		return -1, 0
	case '>':
		return 1, 0
	case '^':
		return 0, -1
	case 'v':
		return 0, 1
	}
	panic("unknown movement")
}

func findRobotPosition(board *[][]rune) (x int, y int) {
	for y, row := range *board {
		for x, symbol := range row {
			if symbol == '@' {
				return x, y
			}
		}
	}
	panic("where is the robot!?")
}

func (s *Solver) tryMoveElement(x, y, dx, dy int, board *[][]rune) bool {
	nx, ny := x+dx, y+dy
	if (*board)[ny][nx] == '#' {
		return false
	}
	if (*board)[ny][nx] == '.' {
		(*board)[ny][nx] = (*board)[y][x]
		(*board)[y][x] = '.'
		return true
	}
	moved := s.tryMoveElement(nx, ny, dx, dy, board)
	if moved {
		(*board)[ny][nx] = (*board)[y][x]
		(*board)[y][x] = '.'
	}
	return moved
}

func (s *Solver) tryMoveElement2(x, y, dx, dy int, board *[][]rune) bool {
	nx, ny := x+dx, y+dy
	if (*board)[ny][nx] == '#' {
		return false
	}
	if (*board)[ny][nx] == '.' {
		(*board)[ny][nx] = (*board)[y][x]
		(*board)[y][x] = '.'
		return true
	}
	var partnerX int
	if (*board)[ny][nx] == '[' {
		if (*board)[ny][nx+1] != ']' {
			visualizeBoard(board)
			panic("wrong")
		}
		partnerX = 1
	} else if (*board)[ny][nx] == ']' {
		if (*board)[ny][nx-1] != '[' {
			visualizeBoard(board)
			panic("wrong")
		}
		partnerX = -1
	}

	if partnerX != 0 {
		if dy != 0 {
			// test on copy :x
			boardCopy := copyBoard(board)
			if !s.tryMoveElement2(nx, ny, dx, dy, boardCopy) ||
				!s.tryMoveElement2(nx+partnerX, ny, dx, dy, boardCopy) {
				return false
			}
			// proceed
			s.tryMoveElement2(nx, ny, dx, dy, board)
			s.tryMoveElement2(nx+partnerX, ny, dx, dy, board)
			(*board)[ny][nx] = (*board)[y][x]
			(*board)[y][x] = '.'
			return true
		}
	}

	moved := s.tryMoveElement(nx, ny, dx, dy, board)
	if moved {
		(*board)[ny][nx] = (*board)[y][x]
		(*board)[y][x] = '.'
	}
	return moved
}

func copyBoard(board *[][]rune) *[][]rune {
	newBoard := make([][]rune, len(*board))
	for i, row := range *board {
		newRow := make([]rune, len(row))
		copy(newRow, row)
		newBoard[i] = newRow
	}
	return &newBoard
}

func visualizeBoard(board *[][]rune) {
	vis.Visualize2dArrayInTerminal(board, func(r rune) color.Color {
		neonPink := color.RGBA{R: 255, G: 16, B: 240, A: 255}
		neonBlue := color.RGBA{R: 17, G: 255, B: 253, A: 255}
		neonYellow := color.RGBA{R: 255, G: 255, A: 255}
		switch r {
		case '@':
			return neonBlue
		case '#':
			return neonYellow
		case 'O':
			return neonPink
		case '[':
			return neonPink
		case ']':
			return neonPink
		default:
			return color.Black
		}
		panic(fmt.Errorf("unknown symbol: %s", string(r)))
	})
}
