package main

import (
	"flag"
	"fmt"
	"github.com/alxdsz/aoc2024/internal/day1"
	"github.com/alxdsz/aoc2024/internal/day10"
	"github.com/alxdsz/aoc2024/internal/day11"
	"github.com/alxdsz/aoc2024/internal/day12"
	"github.com/alxdsz/aoc2024/internal/day13"
	"github.com/alxdsz/aoc2024/internal/day14"
	"github.com/alxdsz/aoc2024/internal/day15"
	"github.com/alxdsz/aoc2024/internal/day16"
	"github.com/alxdsz/aoc2024/internal/day2"
	"github.com/alxdsz/aoc2024/internal/day3"
	"github.com/alxdsz/aoc2024/internal/day4"
	"github.com/alxdsz/aoc2024/internal/day5"
	"github.com/alxdsz/aoc2024/internal/day6"
	"github.com/alxdsz/aoc2024/internal/day7"
	"github.com/alxdsz/aoc2024/internal/day8"
	"github.com/alxdsz/aoc2024/internal/day9"
	"time"
)

type Solver interface {
	SolvePart1() int
	SolvePart2() int
}

type SolverFactory func(string) Solver

var solvers = make(map[int]SolverFactory)

func Register(day int, factory SolverFactory) {
	solvers[day] = factory
}

func main() {
	// Command line flags
	day := flag.Int("day", 0, "Day to solve (0 for all days)")
	part := flag.Int("part", 0, "Part to solve (0 for both parts)")
	flag.Parse()

	// Register all solvers
	Register(1, func(input string) Solver { return day1.NewSolver(input) })
	Register(2, func(input string) Solver { return day2.NewSolver(input) })
	Register(3, func(input string) Solver { return day3.NewSolver(input) })
	Register(4, func(input string) Solver { return day4.NewSolver(input) })
	Register(5, func(input string) Solver { return day5.NewSolver(input) })
	Register(6, func(input string) Solver { return day6.NewSolver(input) })
	Register(7, func(input string) Solver { return day7.NewSolver(input) })
	Register(8, func(input string) Solver { return day8.NewSolver(input) })
	Register(9, func(input string) Solver { return day9.NewSolver(input) })
	Register(10, func(input string) Solver { return day10.NewSolver(input) })
	Register(11, func(input string) Solver { return day11.NewSolver(input) })
	Register(12, func(input string) Solver { return day12.NewSolver(input) })
	Register(13, func(input string) Solver { return day13.NewSolver(input) })
	Register(14, func(input string) Solver { return day14.NewSolver(input) })
	Register(15, func(input string) Solver { return day15.NewSolver(input) })
	Register(16, func(input string) Solver { return day16.NewSolver(input) })

	if *day == 0 {
		runAllDays(*part)
	} else {
		runDay(*day, *part)
	}
}

func runDay(day, part int) {
	solver, exists := solvers[day]
	if !exists {
		fmt.Printf("Day %d not implemented yet\n", day)
		return
	}

	inputPath := fmt.Sprintf("./inputs/d%d.txt", day)
	s := solver(inputPath)

	start := time.Now()

	if part == 0 || part == 1 {
		result := s.SolvePart1()
		elapsed := time.Since(start)
		fmt.Printf("Day %d Part 1: %d (took %s)\n", day, result, elapsed)
	}

	if part == 0 || part == 2 {
		result := s.SolvePart2()
		elapsed := time.Since(start)
		fmt.Printf("Day %d Part 2: %d (took %s)\n", day, result, elapsed)
	}
}

func runAllDays(part int) {
	for day := 1; day <= 25; day++ {
		if _, exists := solvers[day]; exists {
			fmt.Printf("\nRunning Day %d:\n", day)
			runDay(day, part)
		}
	}
}
