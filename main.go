package main

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/day1"
	"github.com/alxdsz/aoc2024/internal/day2"
	"github.com/alxdsz/aoc2024/internal/day3"
	"github.com/alxdsz/aoc2024/internal/day4"
	"github.com/alxdsz/aoc2024/internal/day5"
	"github.com/alxdsz/aoc2024/internal/input"
)

func main() {
	// TODO extract inputs from solver

	d1 := day1.NewDay1Solver("./inputs/d1.txt")
	d1p1 := d1.SolvePart1()
	d1p2 := d1.SolvePart2()
	fmt.Printf("d1p1: %d\n", d1p1)
	fmt.Printf("d1p2: %d\n\n", d1p2)

	d2 := day2.NewDay2Solver("./inputs/d2.txt")
	d2p1 := d2.SolvePart1()
	d2p2 := d2.SolvePart2()
	fmt.Printf("d2p1: %d\n", d2p1)
	fmt.Printf("d2p2: %d\n\n", d2p2)

	d3 := day3.NewDay3Solver("./inputs/d3.txt")
	d3p1 := d3.SolvePart1()
	d3p2 := d3.SolvePart2()
	fmt.Printf("d3p1: %d\n", d3p1)
	fmt.Printf("d3p2: %d\n", d3p2)

	inpt, _ := input.ReadFile("./inputs/d4.txt")
	i := inpt.AsArray()
	d4 := day4.NewDay4Solver(i)
	d4p1 := d4.SolvePart1()
	d4p2 := d4.SolvePart2()
	fmt.Printf("d4p1: %d\n", d4p1)
	fmt.Printf("d4p2: %d\n", d4p2)

	inpt2, _ := input.ReadFile("./inputs/d5.txt")
	d5 := day5.NewDay5Solver(inpt2.SplitByEmptyLine())
	d5p1 := d5.SolvePart1()
	d5p2 := d5.SolvePart2()
	fmt.Printf("d5p1: %d\n", d5p1)
	fmt.Printf("d5p2: %d\n", d5p2)
}
