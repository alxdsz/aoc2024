package day17

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/input"
	"github.com/alxdsz/aoc2024/internal/utils"
	"math"
	"strconv"
	"strings"
)

type Computer struct {
	regA, regB, regC int
	instructionPtr   int
	program          []int
	output           []int
}

type Solver struct {
	computer Computer
}

func NewSolver(inputPath string) *Solver {
	var regA, regB, regC int
	var program []int
	for i, line := range input.ReadFileUnsafe(inputPath).Lines() {
		if i == 0 {
			regA = utils.UnsafeAtoi(strings.TrimPrefix(line, "Register A: "))[0]
			continue
		}
		if i == 1 {
			regB = utils.UnsafeAtoi(strings.TrimPrefix(line, "Register B: "))[0]
			continue
		}
		if i == 2 {
			regC = utils.UnsafeAtoi(strings.TrimPrefix(line, "Register C: "))[0]
			continue
		}
		if i == 3 {
			continue
		}
		program = utils.UnsafeAtoi(strings.Split(strings.TrimPrefix(line, "Program: "), ",")...)
	}
	return &Solver{Computer{regA, regB, regC, 0, program, []int{}}}
}

func (s *Solver) SolvePart1() string {
	return s.computer.run()
}

func (s *Solver) SolvePart2() string {
	return strconv.Itoa(s.findInputProducingOutput())
}

func (c *Computer) getComboForOperand(operand int) int {
	switch operand {
	case 0:
		fallthrough
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		return operand
	case 4:
		return c.regA
	case 5:
		return c.regB
	case 6:
		return c.regC
	}
	panic("invalid operand?")
}

func (c *Computer) getOperationForOpcode(opcode int) func(x int) {
	switch opcode {
	case 0: // adv
		return func(x int) {
			c.regA = c.regA / int(math.Pow(2, float64(c.getComboForOperand(x))))
			c.instructionPtr += 2
		}
	case 1: // bxl
		return func(x int) {
			c.regB = c.regB ^ x
			c.instructionPtr += 2
		}
	case 2: // bst
		return func(x int) {
			c.regB = c.getComboForOperand(x) % 8
			c.instructionPtr += 2
		}
	case 3: // jnz
		return func(x int) {
			if c.regA == 0 {
				c.instructionPtr += 2
				return
			}
			c.instructionPtr = x
		}
	case 4: // bxc
		return func(x int) {
			c.regB = c.regB ^ c.regC
			c.instructionPtr += 2
		}
	case 5: // out
		return func(x int) {
			c.output = append(c.output, c.getComboForOperand(x)%8)
			c.instructionPtr += 2
		}
	case 6: // bdv
		return func(x int) {
			c.regB = c.regA / int(math.Pow(2, float64(c.getComboForOperand(x))))
			c.instructionPtr += 2
		}
	case 7: // cdv
		return func(x int) {
			c.regC = c.regA / int(math.Pow(2, float64(c.getComboForOperand(x))))
			c.instructionPtr += 2
		}
	}
	panic("unknown op")
}

func (c *Computer) run() string {
	for c.instructionPtr < len(c.program) {
		c.getOperationForOpcode(c.program[c.instructionPtr])(c.program[c.instructionPtr+1])
	}
	var outAsStr []string
	for _, out := range c.output {
		outAsStr = append(outAsStr, strconv.Itoa(out))
	}
	// result printed here
	return strings.Join(outAsStr, ",")
}

func (c *Computer) clone() Computer {
	return Computer{
		regA:           c.regA,
		regB:           c.regB,
		regC:           c.regC,
		instructionPtr: 0,
		program:        c.program,
		output:         []int{},
	}
}

// Thank you reddit...
func (s *Solver) findInputProducingOutput() int {
	queue := []int{0, 1, 2, 3, 4, 5, 6, 7}

	for len(queue) > 0 {
		tryA := queue[0]
		queue = queue[1:]

		if len(fmt.Sprintf("%b", tryA))/3+1 >= len(s.computer.program) {
			return tryA
		}

		for i := 0; i < 8; i++ {
			candidateA := (tryA << 3) + i

			testComputer := s.computer.clone()
			testComputer.regA = candidateA
			testComputer.run()

			if len(testComputer.output) > 0 {
				programSlice := s.computer.program[len(s.computer.program)-len(testComputer.output):]
				matches := true
				for i, v := range testComputer.output {
					if v != programSlice[i] {
						matches = false
						break
					}
				}
				if matches {
					queue = append(queue, candidateA)
				}
			}
		}
	}
	panic("no solution?")
}
