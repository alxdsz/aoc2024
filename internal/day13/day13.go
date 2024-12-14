package day13

import (
	"github.com/alxdsz/aoc2024/internal/input"
	"regexp"
	"strconv"
)

type Machine struct {
	buttonA struct{ x, y int }
	buttonB struct{ x, y int }
	prize   struct{ x, y int }
}

type Solver struct {
	machines []Machine
}

func parseMachineDataLine(dataLine string) (int, int) {
	re := regexp.MustCompile(`-?\d+`)
	numbers := re.FindAllString(dataLine, -1)
	number1, _ := strconv.Atoi(numbers[0])
	number2, _ := strconv.Atoi(numbers[1])
	return number1, number2
}

func NewSolver(inputPath string) *Solver {
	inpt, _ := input.ReadFile(inputPath)
	machineDataPacks := inpt.SplitByEmptyLine()
	var machines []Machine
	for _, dataPack := range machineDataPacks {
		buttonAX, buttonAY := parseMachineDataLine(dataPack[0])
		buttonBX, buttonBY := parseMachineDataLine(dataPack[1])
		prizeX, prizeY := parseMachineDataLine(dataPack[2])
		machine := Machine{
			buttonA: struct{ x, y int }{buttonAX, buttonAY},
			buttonB: struct{ x, y int }{buttonBX, buttonBY},
			prize:   struct{ x, y int }{prizeX, prizeY},
		}
		machines = append(machines, machine)
	}
	return &Solver{machines}
}

func (s *Solver) SolvePart1() int {
	totalTokens := 0
	for _, m := range s.machines {
		minTokens := -1
		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100; b++ {
				x := a*m.buttonA.x + b*m.buttonB.x
				y := a*m.buttonA.y + b*m.buttonB.y
				if x == m.prize.x && y == m.prize.y {
					tokens := 3*a + b
					if minTokens == -1 || tokens < minTokens {
						minTokens = tokens
					}
				}
			}
		}
		if minTokens != -1 {
			totalTokens += minTokens
		}
	}
	return totalTokens
}

func (s *Solver) SolvePart2() int {
	totalTokens := 0
	// I wish I had time
	return totalTokens
}
