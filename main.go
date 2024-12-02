package main

import (
	"fmt"
	"github.com/alxdsz/aoc2024/internal/day1"
	"github.com/alxdsz/aoc2024/internal/day2"
)

func main() {
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
}
