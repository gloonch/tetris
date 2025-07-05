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
}

func NewPiece() *Piece {
	rand.Seed(time.Now().UnixNano())
	shape := shapes[rand.Intn(len(shapes))]
	x := 4 // center of the screen
	y := 0
	return &Piece{Shape: shape, X: x, Y: y}
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
