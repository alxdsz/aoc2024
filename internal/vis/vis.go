package vis

import (
	"fmt"
	"image/color"
	"math"
	"strings"
	"time"
)

func Visualize2dArrayInTerminal[T any](grid *[][]T, cell2ColorFun func(T) color.Color) {
	var sb strings.Builder

	// Clear screen and scrollback buffer
	sb.WriteString("\033[2J\033[3J")

	// Find maximum width
	maxWidth := 0
	for _, row := range *grid {
		if len(row) > maxWidth {
			maxWidth = len(row)
		}
	}

	// Move to top-left corner
	sb.WriteString("\033[H")

	// Pre-print blank lines to cover previous content
	blankLine := strings.Repeat(" ", maxWidth*2+1) + "\n" // *2 because each cell is "■ "
	for range *grid {
		sb.WriteString(blankLine)
	}

	// Move back to top
	sb.WriteString(fmt.Sprintf("\033[%dA", len(*grid)))

	// Draw the actual grid
	block := "■ "
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

	time.Sleep(300 * time.Millisecond)
}

func GenerateUniqueColor(i int) color.NRGBA {
	// Use prime numbers for better distribution
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}

	// Get unique angles by combining multiple primes
	angle := float64(i*primes[i%len(primes)]) * math.Pi / 50

	// Vary saturation and value slightly for more uniqueness
	s := 0.8 + 0.2*math.Sin(float64(i))
	v := 0.9 + 0.1*math.Cos(float64(i))

	// Convert polar coordinates to RGB
	r := v * (1 + s*math.Cos(angle))
	g := v * (1 + s*math.Cos(angle+2*math.Pi/3))
	b := v * (1 + s*math.Cos(angle+4*math.Pi/3))

	// Normalize and convert to uint8
	max := math.Max(1.0, math.Max(r, math.Max(g, b)))
	return color.NRGBA{
		R: uint8(255 * r / max),
		G: uint8(255 * g / max),
		B: uint8(255 * b / max),
		A: 255,
	}
}
