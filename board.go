package main

type Board struct {
	Width  int
	Height int
	Cells  [][]rune
}

func NewBoard(width, height int) *Board {
	cells := make([][]rune, height)
	for i := range cells {
		cells[i] = make([]rune, width)
	}
	return &Board{Width: width, Height: height, Cells: cells}
}

func (b *Board) Print(active *Piece) {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			cell := b.Cells[y][x]
			if active != nil && y >= active.Y && y < active.Y+len(active.Shape) && x >= active.X && x < active.X+len(active.Shape[0]) {
				if active.Shape[y-active.Y][x-active.X] != 0 {
					cell = active.Shape[y-active.Y][x-active.X]
				}
			}
			if cell == 0 {
				print(".")
			} else {
				print(string(cell))
			}
		}
		println("")
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
			if b.Cells[ny][nx] != 0 {
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
				b.Cells[p.Y+y][p.X+x] = p.Shape[y][x]
			}
		}
	}
}

func (b *Board) ClearFullRows() int {
	newCells := make([][]rune, 0, b.Height)
	cleared := 0
	for y := b.Height - 1; y >= 0; y-- {
		full := true
		for x := 0; x < b.Width; x++ {
			if b.Cells[y][x] == 0 {
				full = false
				break
			}
		}
		if !full {
			row := make([]rune, b.Width)
			copy(row, b.Cells[y])
			newCells = append([][]rune{row}, newCells...)
		} else {
			cleared++
		}
	}
	for len(newCells) < b.Height {
		newCells = append([][]rune{make([]rune, b.Width)}, newCells...)
	}
	b.Cells = newCells
	return cleared
}
