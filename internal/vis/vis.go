package vis

import (
	"fmt"
	"image/color"
	"strings"
	"time"
)

func clearScreen() {
	// ANSI escape codes:
	// \033[2J  - clear entire screen
	// \033[H   - move cursor to top-left corner
	fmt.Print("\033[2J\033[H")
}

func Visualize2dArrayInTerminal2[T any](grid *[][]T, cell2ColorFun func(T) color.Color) {
	clearScreen()
	// Use Unicode block character
	block := "█"

	for i := range *grid {
		for j := range (*grid)[i] {
			c := cell2ColorFun((*grid)[i][j])
			r, g, b, _ := c.RGBA()
			// Convert from 0-65535 to 0-255 range
			r, g, b = r>>8, g>>8, b>>8
			// Print colored block using ANSI escape codes
			fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, block)
		}
		fmt.Println()
	}
}

// Hide cursor
func hideCursor() {
	fmt.Print("\033[?25l")
}

// Show cursor
func showCursor() {
	fmt.Print("\033[?25h")
}

// Move cursor to specific position
func moveCursor(row, col int) {
	fmt.Printf("\033[%d;%dH", row+1, col+1)
}

func Visualize2dArrayInTerminal[T any](grid *[][]T, cell2ColorFun func(T) color.Color) {
	// Build the entire frame as a string first
	var sb strings.Builder

	// Move to top-left corner
	sb.WriteString("\033[H")

	block := "█"
	for i := range *grid {
		for j := range (*grid)[i] {
			c := cell2ColorFun((*grid)[i][j])
			r, g, b, _ := c.RGBA()
			r, g, b = r>>8, g>>8, b>>8
			sb.WriteString(fmt.Sprintf("\033[38;2;%d;%d;%dm%s", r, g, b, block))
		}
		sb.WriteString("\n")
	}

	// Reset color
	sb.WriteString("\033[0m")

	// Print entire frame at once
	fmt.Print(sb.String())

	time.Sleep(50 * time.Millisecond)
}
