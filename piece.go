package main

import (
	"math/rand"
	"time"
)

var shapes = [][][]rune{
	// square
	{
		{'0', '0'},
		{'0', '0'},
	},
	// line
	{
		{'0', '0', '0', '0'},
	},
	// L
	{
		{'0', 0},
		{'0', 0},
		{'0', '0'},
	},
	// L reversed
	{
		{0, '0'},
		{0, '0'},
		{'0', '0'},
	},
	// T
	{
		{'0', '0', '0'},
		{0, '0', 0},
	},
	// S
	{
		{0, '0', '0'},
		{'0', '0', 0},
	},
	// Z
	{
		{'0', '0', 0},
		{0, '0', '0'},
	},
}

type Piece struct {
	Shape [][]rune
	X, Y  int
	Color string // ANSI color code
}

var shapeColors = []string{
	"\033[33m", // Yellow for square
	"\033[36m", // Cyan for line
	"\033[35m", // Magenta for L
	"\033[34m", // Blue for L reversed
	"\033[32m", // Green for T
	"\033[31m", // Red for S
	"\033[37m", // White for Z
}

func NewPiece() *Piece {
	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(shapes))
	shape := shapes[idx]
	x := 4 // center of the screen
	y := 0
	color := shapeColors[idx]
	return &Piece{Shape: shape, X: x, Y: y, Color: color}
}

func (p *Piece) Rotate() {
	old := p.Shape
	h := len(old)
	w := len(old[0])
	newShape := make([][]rune, w)
	for i := range newShape {
		newShape[i] = make([]rune, h)
		for j := range newShape[i] {
			newShape[i][j] = old[h-1-j][i]
		}
	}
	p.Shape = newShape
}
