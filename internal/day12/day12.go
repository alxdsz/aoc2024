package day12

import (
	"github.com/alxdsz/aoc2024/internal/input"
	"github.com/alxdsz/aoc2024/internal/utils"
	"github.com/alxdsz/aoc2024/internal/vis"
	"image/color"
	"math"
	"slices"
)

type Coordinates struct {
	x, y int
}

type Farm struct {
	board        [][]rune
	plantToColor map[rune]color.Color
}

func (f *Farm) inBounds(c Coordinates) bool {
	return c.x >= 0 && c.x < len(f.board[0]) && c.y >= 0 && c.y < len(f.board)
}

func (f *Farm) getAllNeighbours(spot Coordinates) []Coordinates {
	directions := []Coordinates{
		{0, -1}, // up
		{0, 1},  // down
		{-1, 0}, // left
		{1, 0},  // right
	}

	var neighbours []Coordinates

	for _, dir := range directions {
		neighbor := Coordinates{spot.x + dir.x, spot.y + dir.y}
		neighbours = append(neighbours, neighbor)
	}

	return neighbours
}

func (f *Farm) getDirectNeighbours(spot Coordinates) []Coordinates {
	directions := []Coordinates{
		{0, -1}, // up
		{0, 1},  // down
		{-1, 0}, // left
		{1, 0},  // right
	}

	var neighbours []Coordinates
	sourcePlant := f.getPlant(spot)

	for _, dir := range directions {
		neighbor := Coordinates{spot.x + dir.x, spot.y + dir.y}
		if f.inBounds(neighbor) && f.getPlant(neighbor) == sourcePlant {
			neighbours = append(neighbours, neighbor)
		}
	}

	return neighbours
}

func (f *Farm) getDirectNotNeighbours(spot Coordinates) []Coordinates {
	directions := []Coordinates{
		{0, -1}, // up
		{0, 1},  // down
		{-1, 0}, // left
		{1, 0},  // right
	}

	var neighbours []Coordinates

	for _, dir := range directions {
		neighbor := Coordinates{spot.x + dir.x, spot.y + dir.y}
		if !f.inBounds(neighbor) || f.getPlant(neighbor) != f.getPlant(spot) {
			neighbours = append(neighbours, neighbor)
		}
	}

	return neighbours
}

func (f *Farm) getPlant(c Coordinates) rune {
	return f.board[c.y][c.x]
}

func (f *Farm) getArea(spot Coordinates) (map[Coordinates]bool, int) {
	visitedSpots := make(map[Coordinates]bool)
	totalAreaPrice := 0
	perimeter := 0
	var recursiveNeighbourSearch func(spot Coordinates)
	recursiveNeighbourSearch = func(spot Coordinates) {
		visitedSpots[spot] = true
		directNeighbours := f.getDirectNeighbours(spot)
		perimeter += 4 - len(directNeighbours)
		if len(directNeighbours) == 0 {
			return
		}
		for _, neighbour := range directNeighbours {
			if !visitedSpots[neighbour] {
				recursiveNeighbourSearch(neighbour)
			}
		}
	}
	recursiveNeighbourSearch(spot)
	totalAreaPrice += perimeter * len(visitedSpots)
	return visitedSpots, totalAreaPrice
}

func (f *Farm) buildColorMap() {
	colorNo := 0
	for _, row := range f.board {
		for _, plant := range row {
			if f.plantToColor[plant] == nil {
				f.plantToColor[plant] = vis.GenerateUniqueColor(colorNo)
				colorNo++
			}
		}
	}
}

func (f *Farm) visualize() {
	vis.Visualize2dArrayInTerminal(&f.board, func(p rune) color.Color {
		return f.plantToColor[p]
	})
}

type Solver struct {
	farm Farm
}

func NewSolver(inputPath string) *Solver {
	inpt, _ := input.ReadFile(inputPath)
	farm := Farm{inpt.As2DRuneArray(), make(map[rune]color.Color)}
	farm.buildColorMap()
	return &Solver{farm}
}

func (s *Solver) SolvePart1() int {
	visited := make(map[Coordinates]bool)
	totalPrice := 0
	for y, row := range s.farm.board {
		for x, _ := range row {
			if visited[Coordinates{x, y}] {
				continue
			}
			spots, price := s.farm.getArea(Coordinates{x, y})
			for spot := range spots {
				visited[spot] = true
			}
			totalPrice += price
		}
	}
	return totalPrice
}

// patchwork
func (s *Solver) countSides(spots map[Coordinates]bool) int {
	if len(spots) == 1 {
		return 4
	}
	result := 0
	threeCount := 0
	twoCount := 0
	for spot, _ := range spots {
		neighboursOfSpot := s.farm.getDirectNotNeighbours(spot)
		if len(neighboursOfSpot) == 3 {
			result += 2
			threeCount++
		}
		if len(neighboursOfSpot) == 2 {
			if neighboursOfSpot[0].x != neighboursOfSpot[1].x && neighboursOfSpot[0].y != neighboursOfSpot[1].y {
				result += 1
				twoCount++
			}
		}
	}
	commonNeighbourMap := make(map[Coordinates][]Coordinates)
	for spot, _ := range spots {
		neighboursOfSpot := s.farm.getDirectNotNeighbours(spot)
		visitedSpotForNeigbour := make(map[Coordinates]bool)
		for otherSpot, _ := range spots {
			if otherSpot != spot {
				visitedSpotForNeigbour[spot] = true
				neighboursOfOtherSpot := s.farm.getDirectNotNeighbours(otherSpot)
				for _, n := range neighboursOfOtherSpot {
					if slices.Contains(neighboursOfSpot, n) {
						commonNeighbourMap[n] = append(commonNeighbourMap[n], spot)
					}
				}
			}
		}
	}

	for k, v := range commonNeighbourMap {
		for k2, v2 := range commonNeighbourMap {
			if k != k2 {
				if utils.SlicesEqual(v, v2) {
					delete(commonNeighbourMap, k2)
					delete(commonNeighbourMap, k)
				}
			}
		}
	}

	for _, v := range commonNeighbourMap {
		if len(v) == 2 {
			if math.Abs(float64(v[0].x-v[1].x)) > 1 || math.Abs(float64(v[0].y-v[1].y)) > 1 {
				continue
			}
			result += 1
		}
		if len(v) > 2 && len(v) < 12 {
			result += 2
		}
		if len(v) >= 12 {
			if math.Abs(float64(v[0].x-v[1].x)) > 1 || math.Abs(float64(v[0].y-v[1].y)) > 1 {
				continue
			}
			result += 4
		}
	}

	return result

}

func (s *Solver) SolvePart2() int {
	visited := make(map[Coordinates]bool)
	s.farm.visualize()
	totalPrice := 0
	for y, row := range s.farm.board {
		for x, _ := range row {
			if visited[Coordinates{x, y}] {
				continue
			}
			spots, _ := s.farm.getArea(Coordinates{x, y})
			for spot := range spots {
				visited[spot] = true
			}
			sides := s.countSides(spots)
			totalPrice += len(spots) * sides
		}
	}
	return totalPrice
}
