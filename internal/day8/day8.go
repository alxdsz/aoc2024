package day8

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/input"
	"math"
)

type FreqMap map[string][]Coordinates

type Coordinates struct {
	x, y float64
}

type Solver struct {
	freqs   FreqMap
	xLength int
	yLength int
}

func NewSolver(inputPath string) *Solver {
	freqMap := make(FreqMap)
	inpt, _ := input.ReadFile(inputPath)
	for y, row := range inpt.Lines() {
		for x, cell := range row {
			if cell != '.' {
				cellAsStr := string(cell)
				freqMap[cellAsStr] = append(freqMap[cellAsStr], Coordinates{float64(x), float64(y)})
			}
		}
	}
	return &Solver{
		freqs:   freqMap,
		yLength: len(inpt.Lines()),
		xLength: len(inpt.Lines()[0]),
	}
}

func calculateCoordinatesDiff(posA, posB Coordinates) Coordinates {
	return Coordinates{
		x: posB.x - posA.x,
		y: posB.y - posA.y,
	}
}

func calculateAntiNodePositions(pos, diff Coordinates) (Coordinates, Coordinates) {
	antiNodeA := Coordinates{
		x: pos.x + (diff.x * 2),
		y: pos.y + (diff.y * 2),
	}
	antiNodeB := Coordinates{
		x: pos.x - (diff.x),
		y: pos.y - (diff.y),
	}
	return antiNodeA, antiNodeB
}

func (s *Solver) validateAntiNodePosition(pos Coordinates) bool {
	isOnGrid := math.Mod(pos.x, 1) == 0 && math.Mod(pos.y, 1) == 0
	isWithinBoundaries := pos.x >= 0 && pos.x < float64(s.xLength) && pos.y >= 0 && pos.y < float64(s.yLength)
	return isOnGrid && isWithinBoundaries
}

func (s *Solver) SolvePart1() int {
	uniquePositions := make(map[string]bool)
	for _, freqs := range s.freqs {

		var pairs [][]Coordinates
		for i, posA := range freqs {
			for j, posB := range freqs {
				if i < j { // Only generate each pair once
					pairs = append(pairs, []Coordinates{posA, posB})
				}
			}
		}

		for _, p := range pairs {
			diff := calculateCoordinatesDiff(p[0], p[1])
			anA, anB := calculateAntiNodePositions(p[0], diff)
			if s.validateAntiNodePosition(anA) {
				uniquePositions[fmt.Sprintf("%.0f,%.0f", anA.x, anA.y)] = true
			}
			if s.validateAntiNodePosition(anB) {
				uniquePositions[fmt.Sprintf("%.0f,%.0f", anB.x, anB.y)] = true
			}
		}

	}

	return len(uniquePositions)
}

func isPointOnLine(a, b, p Coordinates) bool {
	crossProduct := (p.y-a.y)*(b.x-a.x) - (p.x-a.x)*(b.y-a.y)
	if math.Abs(crossProduct) > 0.0001 {
		return false
	}
	return true
}

func (s *Solver) SolvePart2() int {
	uniquePositions := make(map[string]bool)

	for _, freqs := range s.freqs {
		if len(freqs) > 1 {
			for _, pos := range freqs {
				uniquePositions[fmt.Sprintf("%.0f,%.0f", pos.x, pos.y)] = true
			}
		}

		for i, posA := range freqs {
			for j, posB := range freqs {
				if i < j {
					for y := 0.0; y < float64(s.yLength); y++ {
						for x := 0.0; x < float64(s.xLength); x++ {
							if isPointOnLine(posA, posB, Coordinates{x: x, y: y}) {
								uniquePositions[fmt.Sprintf("%.0f,%.0f", x, y)] = true
							}
						}
					}
				}
			}
		}
	}

	return len(uniquePositions)
}
