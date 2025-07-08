package main

import "fmt"

type Cell struct {
	Value rune
	Color string
}

type Board struct {
	Width  int
	Height int
	Cells  [][]Cell
}

func NewBoard(width, height int) *Board {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}
	return &Board{Width: width, Height: height, Cells: cells}
}

func (b *Board) Print(active *Piece) {
	reset := "\033[0m"
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			cell := b.Cells[y][x]
			printed := false
			if active != nil && y >= active.Y && y < active.Y+len(active.Shape) && x >= active.X && x < active.X+len(active.Shape[0]) {
				if active.Shape[y-active.Y][x-active.X] != 0 {
					// Print active piece with its color
					fmt.Print(active.Color + string(active.Shape[y-active.Y][x-active.X]) + reset)
					printed = true
				}
			}
			if !printed {
				if cell.Value != 0 {
					fmt.Print(cell.Color + string(cell.Value) + reset)
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
}

func (b *Board) CanMove(p *Piece, dx, dy int) bool {
	for y := 0; y < len(p.Shape); y++ {
		for x := 0; x < len(p.Shape[0]); x++ {
			if p.Shape[y][x] == 0 {
				continue
			}
			nx := p.X + x + dx
			ny := p.Y + y + dy
			if nx < 0 || nx >= b.Width || ny < 0 || ny >= b.Height {
				return false
			}
			if b.Cells[ny][nx].Value != 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) PlacePiece(p *Piece) {
	for y := 0; y < len(p.Shape); y++ {
		for x := 0; x < len(p.Shape[0]); x++ {
			if p.Shape[y][x] != 0 {
				b.Cells[p.Y+y][p.X+x] = Cell{Value: p.Shape[y][x], Color: p.Color}
			}
		}
	}
}

func (b *Board) ClearFullRows() int {
	newCells := make([][]Cell, 0, b.Height)
	cleared := 0
	for y := b.Height - 1; y >= 0; y-- {
		full := true
		for x := 0; x < b.Width; x++ {
			if b.Cells[y][x].Value == 0 {
				full = false
				break
			}
		}
		if !full {
			row := make([]Cell, b.Width)
			copy(row, b.Cells[y])
			newCells = append([][]Cell{row}, newCells...)
		} else {
			cleared++
		}
	}
	for len(newCells) < b.Height {
		newCells = append([][]Cell{make([]Cell, b.Width)}, newCells...)
	}
	b.Cells = newCells
	return cleared
}
