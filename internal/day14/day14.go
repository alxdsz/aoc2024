package day14

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/input"
	"github.com/alxdsz/aoc2024/internal/utils"
	"github.com/alxdsz/aoc2024/internal/vis"
	"image/color"
	"math"
	"strings"
)

type Robot struct {
	x, y   int
	vx, vy int
}

func (r *Robot) doStep(Xmax, Ymax int) {
	newX := r.x + r.vx
	newY := r.y + r.vy

	if newX > Xmax-1 {
		newX = newX - Xmax
	}
	if newX < 0 {
		newX = newX + Xmax
	}
	if newY > Ymax-1 {
		newY = newY - Ymax
	}
	if newY < 0 {
		newY = newY + Ymax
	}

	r.x = newX
	r.y = newY
}

func (r *Robot) isSurrounded(board *[][]bool) bool {
	directions := []struct{ x, y int }{
		{0, -1},  // up
		{0, 1},   // down
		{-1, 0},  // left
		{1, 0},   // right
		{-1, -1}, // up-left
		{1, -1},  // up-right
		{-1, 1},  // down-left
		{1, 1},   // down-right
	}

	for _, dir := range directions {
		dx := r.x + dir.x
		dy := r.y + dir.y
		if dx >= 0 && dx < len((*board)[0]) && dy >= 0 && dy < len(*board) {
			if !(*board)[dy][dx] {
				return false
			}
		}
	}

	return true
}

type Solver struct {
	Xmax, Ymax int
	robots     []*Robot
	inputPath  string
}

func NewSolver(inputPath string) *Solver {
	inpt, _ := input.ReadFile(inputPath)
	lines := inpt.Lines()
	robots, xMax, yMax := readInput(lines)
	return &Solver{xMax, yMax, robots, inputPath}
}

func readInput(lines []string) ([]*Robot, int, int) {
	var robots []*Robot
	xMax, yMax := 101, 103
	for _, line := range lines {
		mainSplit := strings.Split(line, " ")
		position, velocity := mainSplit[0], mainSplit[1]
		positionValues := utils.UnsafeAtoi(strings.Split(strings.Split(position, "=")[1], ",")...)
		velocityValues := utils.UnsafeAtoi(strings.Split(strings.Split(velocity, "=")[1], ",")...)
		robots = append(robots, &Robot{positionValues[0], positionValues[1], velocityValues[0], velocityValues[1]})

	}
	return robots, xMax, yMax
}

func (s *Solver) getQuadrant(x, y int) int {
	if x == (s.Xmax-1)/2 || y == (s.Ymax-1)/2 {
		return -1
	}
	if x < (s.Xmax-1)/2 {
		if y < (s.Ymax-1)/2 {
			return 1
		} else {
			return 3
		}
	} else {
		if y < (s.Ymax-1)/2 {
			return 2
		} else {
			return 4
		}
	}
}

// I mean.. why not? :d
// Method with searching for an iteration where at least one robot is fully surrounded by other robots works well too
// But couldn't help myself checking if the an entropy based solution would work - it did, math rules
// The entropy calculation below were generated using AI
func calculateSpatialEntropy(array [][]bool) float64 {
	metrics := make(map[string]float64)

	rows := len(array)
	cols := len(array[0])

	// horizontal transition entropy
	horizontalTransitions := 0
	for i := 0; i < rows; i++ {
		for j := 1; j < cols; j++ {
			if array[i][j] != array[i][j-1] {
				horizontalTransitions++
			}
		}
	}

	//  vertical transition entropy
	verticalTransitions := 0
	for j := 0; j < cols; j++ {
		for i := 1; i < rows; i++ {
			if array[i][j] != array[i-1][j] {
				verticalTransitions++
			}
		}
	}

	// 3. block entropy (2x2 patterns)
	blockPatterns := make(map[[4]bool]int)
	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols-1; j++ {
			pattern := [4]bool{
				array[i][j], array[i][j+1],
				array[i+1][j], array[i+1][j+1],
			}
			blockPatterns[pattern]++
		}
	}

	// normalized metrics
	totalBlocks := float64((rows - 1) * (cols - 1))

	// normalize transitions
	horizontalDensity := float64(horizontalTransitions) / float64(rows*(cols-1))
	verticalDensity := float64(verticalTransitions) / float64(cols*(rows-1))

	// block pattern entropy
	blockEntropy := 0.0
	for _, count := range blockPatterns {
		p := float64(count) / totalBlocks
		blockEntropy -= p * math.Log2(p)
	}

	// Store all metrics
	metrics["horizontal_transition_density"] = horizontalDensity
	metrics["vertical_transition_density"] = verticalDensity
	metrics["block_entropy"] = blockEntropy

	// Calculate an "image likelihood score" (lower means more likely to be an image)
	// Images tend to have:
	// 1. Moderate transition densities (not too high or too low)
	// 2. Similar horizontal and vertical transition densities
	// 3. Lower block entropy than random noise
	transitionBalance := math.Abs(horizontalDensity - verticalDensity)
	idealTransitionDensity := math.Abs((horizontalDensity+verticalDensity)/2 - 0.3)

	imageLikelihood := (transitionBalance + idealTransitionDensity + blockEntropy) / 3

	return imageLikelihood
}

func (s *Solver) SolvePart1() int {
	seconds := 100
	quadrants := []int{0, 0, 0, 0}
	for _, robot := range s.robots {
		for i := 0; i < seconds; i++ {
			robot.doStep(s.Xmax, s.Ymax)
		}
		quadrant := s.getQuadrant(robot.x, robot.y)
		if quadrant != -1 {
			quadrants[quadrant-1]++
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func (s *Solver) SolvePart2() int {
	var board [][]bool
	for y := 0; y < s.Ymax; y++ {
		board = append(board, []bool{})
		for x := 0; x < s.Xmax; x++ {
			board[y] = append(board[y], false)
		}
	}

	i := 0
	isClose := false
	for {
		i++
		for _, robot := range s.robots {
			board[robot.y][robot.x] = false
			robot.doStep(s.Xmax, s.Ymax)
			board[robot.y][robot.x] = true
		}
		entropy := calculateSpatialEntropy(board)
		if entropy < 0.4 || isClose {
			isClose = true
			vis.Visualize2dArrayInTerminal(&board, func(b bool) color.Color {
				if b {
					return color.Black
				} else {
					return color.White
				}
			})
			var in string
			_, _ = fmt.Scanln(&in)
			if in == "f" {
				return i
			}
		}
		if i > 10000 {
			return -1
		}
	}
}
