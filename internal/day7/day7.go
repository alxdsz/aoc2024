package day7

import (
	"github.com/alxdsz/aoc2024/internal/input"
	"math"
	"strconv"
	"strings"
)

type Equation struct {
	expectedResult int
	components     []int
}

type Solver struct {
	calibrationEquations []Equation
}

func convertInputToEquations(input *input.Input) []Equation {
	var result []Equation
	for _, line := range input.Lines() {
		split := strings.Split(line, ": ")

		expectedResultAsString := split[0]
		expectedResult, _ := strconv.Atoi(expectedResultAsString)

		splitComponents := strings.Split(split[1], " ")
		components := make([]int, len(splitComponents))
		for i, component := range splitComponents {
			components[i], _ = strconv.Atoi(strings.Trim(component, " "))
		}

		result = append(result, Equation{
			expectedResult: expectedResult,
			components:     components,
		})
	}
	return result
}

func NewSolver(inputhPath string) *Solver {
	inpt, _ := input.ReadFile(inputhPath)
	return &Solver{calibrationEquations: convertInputToEquations(inpt)}
}

// We only have two operators, so we can generate operations using bits
func generateTwoOperatorCombinations(operatorCount int) [][]string {
	totalCombinations := int(math.Pow(2, float64(operatorCount)))
	combinations := make([][]string, totalCombinations)
	for i := 0; i < totalCombinations; i++ {
		combination := make([]string, operatorCount)
		for pos := 0; pos < operatorCount; pos++ {
			// Bit == 1 -> "*"
			if (i>>pos)&1 == 1 {
				combination[pos] = "*"
			} else {
				combination[pos] = "+"
			}
		}
		combinations[i] = combination
	}
	return combinations
}

func (s *Solver) SolvePart1() string {
	solution := 0
	for _, eq := range s.calibrationEquations {
		numOfOperators := len(eq.components) - 1
		operatorCombinations := generateTwoOperatorCombinations(numOfOperators)
		for _, combination := range operatorCombinations {
			testedResult := eq.components[0]
			for i, operation := range combination {
				switch operation {
				case "+":
					testedResult += eq.components[i+1]
				case "*":
					testedResult *= eq.components[i+1]
				}
				if testedResult > eq.expectedResult {
					break
				}
			}
			if testedResult == eq.expectedResult {
				solution += testedResult
				break
			}
		}
	}
	return strconv.Itoa(solution)
}

func generateThreeOperatorCombinations(operatorCount int) [][]string {
	totalCombinations := int(math.Pow(3, float64(operatorCount)))
	combinations := make([][]string, totalCombinations)
	for i := 0; i < totalCombinations; i++ {
		combination := make([]string, operatorCount)
		num := i
		for pos := 0; pos < operatorCount; pos++ {
			switch num % 3 {
			case 0:
				combination[pos] = "+"
			case 1:
				combination[pos] = "*"
			case 2:
				combination[pos] = "||"
			}
			num /= 3
		}
		combinations[i] = combination
	}
	return combinations
}

func (s *Solver) SolvePart2() string {
	solution := 0
	for _, eq := range s.calibrationEquations {
		numOfOperators := len(eq.components) - 1
		operatorCombinations := generateThreeOperatorCombinations(numOfOperators)
		for _, combination := range operatorCombinations {
			testedResult := strconv.Itoa(eq.components[0])
			valid := true
			for i, operation := range combination {
				nextNum := eq.components[i+1]
				switch operation {
				case "+":
					currentVal, _ := strconv.Atoi(testedResult)
					testedResult = strconv.Itoa(currentVal + nextNum)
				case "*":
					currentVal, _ := strconv.Atoi(testedResult)
					testedResult = strconv.Itoa(currentVal * nextNum)
				case "||":
					testedResult = testedResult + strconv.Itoa(nextNum)
				}
				if val, _ := strconv.Atoi(testedResult); val > eq.expectedResult {
					valid = false
					break
				}
			}

			finalVal, _ := strconv.Atoi(testedResult)
			if valid && finalVal == eq.expectedResult {
				solution += eq.expectedResult
				break
			}
		}
	}
	return strconv.Itoa(solution)
}
