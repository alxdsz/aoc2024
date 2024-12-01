package main

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/day1"
)

func main() {
	d1 := day1.NewDay1Solver("./inputs/d1.txt")
	d1p1 := d1.SolvePart1()
	d1p2 := d1.SolvePart2()
	fmt.Printf("d1p1: %d\n", d1p1)
	fmt.Printf("d1p2: %d\n", d1p2)
}
