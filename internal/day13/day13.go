package day13

import (
	"errors"
	"github.com/alxdsz/aoc2024/internal/input"
	"regexp"
	"strconv"
)

type Machine struct {
	buttonA  struct{ x, y int }
	buttonB  struct{ x, y int }
	prize    struct{ x, y int }
	solution *struct{ a, b int }
}

type Solver struct {
	machines []Machine
	maxPress int
	costA    int
	costB    int
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
	return &Solver{machines: machines, costA: 3, costB: 1}
}

func solveDirectSubstitution(m Machine) (*struct{ a, b int }, error) {
	x0, y0 := m.buttonA.x, m.buttonA.y
	x1, y1 := m.buttonB.x, m.buttonB.y
	cx, cy := m.prize.x, m.prize.y

	Bdividend := (cy*x0 - cx*y0)
	Bdivisor := (y1*x0 - y0*x1)
	if Bdivisor == 0 {
		return nil, errors.New("zero division for B")
	} else if Bdividend%Bdivisor != 0 {
		return nil, errors.New("non integer solution")
	}
	B := Bdividend / Bdivisor

	Adividend := (cx - B*x1)
	Adivisor := x0
	if Adivisor == 0 {
		return nil, errors.New("zero division for A")
	} else if Adividend%Adivisor != 0 {
		return nil, errors.New("non integer solution")
	}
	A := Adividend / Adivisor

	return &struct{ a, b int }{A, B}, nil
}

func (s *Solver) solve() {
	for i := range s.machines {
		if solution, err := solveDirectSubstitution(s.machines[i]); err == nil {
			s.machines[i].solution = solution
		}
	}
}

func (s *Solver) countTokens() int {
	tokens := 0
	for _, machine := range s.machines {
		if machine.solution == nil {
			continue
		}
		if s.maxPress > 0 && (machine.solution.a > s.maxPress || machine.solution.b > s.maxPress) {
			continue
		}
		tokens += machine.solution.a*s.costA + machine.solution.b*s.costB
	}
	return tokens
}

func (s *Solver) SolvePart1() string {
	s.maxPress = 100
	s.solve()
	return strconv.Itoa(s.countTokens())
}

func (s *Solver) SolvePart2() string {
	s.maxPress = -1
	const offset = 10000000000000
	for i := range s.machines {
		s.machines[i].prize.x += offset
		s.machines[i].prize.y += offset
	}
	s.solve()
	return strconv.Itoa(s.countTokens())
}
