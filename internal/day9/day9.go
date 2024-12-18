package day9

import (
	"github.com/alxdsz/aoc2024/internal/input"
	"math"
	"slices"
	"strconv"
)

type Solver struct {
	diskMap string
}

type File struct {
	id       int
	start    int
	length   int
	hasMoved bool
}

func NewSolver(inputPath string) *Solver {
	inpt, _ := input.ReadFile(inputPath)
	return &Solver{
		diskMap: inpt.Lines()[0],
	}
}

func (s Solver) convertDiskMapToFS() ([]string, []int, int) {
	var fs []string
	var space []int
	fileId := 0
	currentIndex := 0
	numOfFiles := 0
	for i, valueAsRune := range s.diskMap {
		value, _ := strconv.Atoi(string(valueAsRune))
		if math.Mod(float64(i), 2) == 0 {
			for j := 0; j < value; j++ {
				fs = append(fs, strconv.Itoa(fileId))
				currentIndex++
				numOfFiles++
			}
			fileId++
		} else {
			for j := 0; j < value; j++ {
				fs = append(fs, ".")
				space = append(space, currentIndex)
				currentIndex++
			}
		}
	}
	return fs, space, numOfFiles
}

func isFinito(fs []string, numOfFiles int) bool {
	for i := 0; i < numOfFiles; i++ {
		if fs[i] == "." {
			return false
		}
	}
	return true
}

func (s *Solver) SolvePart1() string {
	fs, space, numOfFiles := s.convertDiskMapToFS()
	currentIndx := len(fs) - 1
	for i := 0; i < len(space); i++ {
		for fs[currentIndx] == "." && currentIndx >= 0 {
			currentIndx--
		}
		if currentIndx == 0 {
			break
		}
		fs[space[i]] = fs[currentIndx]
		space[i] = -1
		fs[currentIndx] = "."
		space = append(space, currentIndx)
		slices.Sort(space)
		if isFinito(fs, numOfFiles) {
			break
		}
	}
	result := 0
	for i, val := range fs {
		if val == "." {
			break
		}
		valInt, _ := strconv.Atoi(val)
		result += i * valInt
	}
	return strconv.Itoa(result)
}

func (s Solver) getDiskFiles() []File {
	var files []File
	fileId := 0
	currentIndex := 0

	for i, valueAsRune := range s.diskMap {
		value, _ := strconv.Atoi(string(valueAsRune))
		if math.Mod(float64(i), 2) == 0 {
			files = append(files, File{
				id:       fileId,
				start:    currentIndex,
				length:   value,
				hasMoved: false,
			})
			currentIndex += value
			fileId++
		} else {
			currentIndex += value
		}
	}
	return files
}

func findFreeSpace(fs []string, length int) int {
	count := 0
	startPos := -1

	for i, val := range fs {
		if val == "." {
			if count == 0 {
				startPos = i
			}
			count++
			if count == length {
				return startPos
			}
		} else {
			count = 0
			startPos = -1
		}
	}
	return -1
}

func (s *Solver) SolvePart2() string {
	fs, _, _ := s.convertDiskMapToFS()
	files := s.getDiskFiles()

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		if file.hasMoved {
			continue
		}

		newStart := findFreeSpace(fs, file.length)
		if newStart == -1 || newStart >= file.start {
			continue
		}

		for j := 0; j < file.length; j++ {
			fs[newStart+j] = strconv.Itoa(file.id)
			fs[file.start+j] = "."
		}
		files[i].start = newStart
		files[i].hasMoved = true
	}

	result := 0
	for i, val := range fs {
		if val == "." {
			continue
		}
		valInt, _ := strconv.Atoi(val)
		result += i * valInt
	}
	return strconv.Itoa(result)
}
