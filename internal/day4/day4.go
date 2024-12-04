package day4

import (
	"strings"
)

type Day4Solver struct {
	input [][]string
}

func NewDay4Solver(input [][]string) *Day4Solver {
	hitMatrix := make([][]bool, len(input))
	for i := range hitMatrix {
		hitMatrix[i] = make([]bool, len(input[0]))
	}

	return &Day4Solver{
		input: input,
	}
}

func matchesTarget(s string, tsd TraverseSearchData) bool {
	return s == tsd.searchedWord || s == tsd.searchedWordReversed
}

func (d *Day4Solver) traverseForward(x, y int, tsd TraverseSearchData) bool {
	wordLength := len(tsd.searchedWord)
	currentLetter := d.input[y][x]
	if x <= tsd.rightBoundary && strings.Contains(tsd.searchedWord, currentLetter) {
		testWord := ""
		for i := 0; i < wordLength; i++ {
			testWord += d.input[y][x+i]
		}
		if matchesTarget(testWord, tsd) {
			return true
		}
	}
	return false
}

func (d *Day4Solver) traverseDown(x, y int, tsd TraverseSearchData) bool {
	wordLength := len(tsd.searchedWord)
	currentLetter := d.input[y][x]
	if y <= tsd.bottomBoundary && strings.Contains(tsd.searchedWord, currentLetter) {
		testWord := ""
		for i := 0; i < wordLength; i++ {
			testWord += d.input[y+i][x]
		}
		if matchesTarget(testWord, tsd) {
			return true
		}
	}
	return false
}

func (d *Day4Solver) traverseDiagonalRight(x, y int, tsd TraverseSearchData) bool {
	wordLength := len(tsd.searchedWord)
	currentLetter := d.input[y][x]
	if y <= tsd.bottomBoundary && x <= tsd.rightBoundary && strings.Contains(tsd.searchedWord, currentLetter) {
		testWord := d.input[y][x]
		for i := 1; i < wordLength; i++ {
			testWord += d.input[y+i][x+i]
		}
		if matchesTarget(testWord, tsd) {
			return true
		}
	}
	return false
}

func (d *Day4Solver) traverseDiagonalLeft(x, y int, tsd TraverseSearchData) bool {
	wordLength := len(tsd.searchedWord)
	currentLetter := d.input[y][x]
	if y <= tsd.bottomBoundary && x >= tsd.leftBoundary && strings.Contains(tsd.searchedWord, currentLetter) {
		testWord := d.input[y][x]
		for i := 1; i < wordLength; i++ {
			testWord += d.input[y+i][x-i]
		}
		if matchesTarget(testWord, tsd) {
			return true
		}
	}
	return false
}

type TraverseSearchData struct {
	leftBoundary   int
	rightBoundary  int
	bottomBoundary int

	searchedWord         string
	searchedWordReversed string
}

func (d *Day4Solver) SolvePart1() int {
	searchedWord := "XMAS"
	searchedWordReversed := "SAMX"
	tsd := TraverseSearchData{
		searchedWord:         searchedWord,
		searchedWordReversed: searchedWordReversed,
		leftBoundary:         len(searchedWord) - 1,
		rightBoundary:        len(d.input[0]) - len(searchedWord),
		bottomBoundary:       len(d.input) - len(searchedWord),
	}

	result := 0
	for y := 0; y < len(d.input); y++ {
		for x := 0; x < len(d.input[y]); x++ {
			if d.traverseForward(x, y, tsd) {
				result++
			}
			if d.traverseDown(x, y, tsd) {
				result++
			}
			if d.traverseDiagonalRight(x, y, tsd) {
				result++
			}
			if d.traverseDiagonalLeft(x, y, tsd) {
				result++
			}
		}
	}
	return result
}

func (d *Day4Solver) SolvePart2() int {
	searchedWord := "MAS"
	searchedWordReversed := "SAM"
	tsd := TraverseSearchData{
		searchedWord:         searchedWord,
		searchedWordReversed: searchedWordReversed,
		leftBoundary:         len(searchedWord) - 1,
		rightBoundary:        len(d.input[0]) - len(searchedWord),
		bottomBoundary:       len(d.input) - len(searchedWord),
	}

	result := 0
	for y := 0; y < len(d.input); y++ {
		for x := 0; x < len(d.input[y])-2; x++ {
			if d.traverseDiagonalRight(x, y, tsd) && d.traverseDiagonalLeft(x+2, y, tsd) {
				result++
			}
		}
	}
	return result
}
