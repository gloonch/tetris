package main

import (
	"fmt"
	"time"
)

func clearScreen() {
	fmt.Print("\033[2J\033[H")
}

func main() {
	board := NewBoard(10, 20)
	piece := NewPiece()
	cmds := make(chan string)
	paused := false
	score := 0

	go ListenInput(cmds)

mainloop:
	for {
		select {
		case cmd := <-cmds:
			switch cmd {
			case "left":
				if board.CanMove(piece, -1, 0) {
					piece.X--
				}
			case "right":
				if board.CanMove(piece, 1, 0) {
					piece.X++
				}
			case "down":
				if board.CanMove(piece, 0, 1) {
					piece.Y++
				}
			case "rotate":
				oldShape := piece.Shape
				piece.Rotate()
				if !board.CanMove(piece, 0, 0) {
					piece.Shape = oldShape
				}
			case "pause":
				paused = !paused
			case "quit":
				return
			}
		default:
			if paused {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			clearScreen()
			fmt.Printf("Score: %d\n", score)
			board.Print(piece)
			time.Sleep(500 * time.Millisecond)
			if board.CanMove(piece, 0, 1) {
				piece.Y++
			} else {
				board.PlacePiece(piece)
				cleared := board.ClearFullRows()
				if cleared > 0 {
					score += cleared * (cleared + 1) / 2
				}
				piece = NewPiece()
				if !board.CanMove(piece, 0, 0) {
					break mainloop
				}
			}
		}
	}
	clearScreen()
	fmt.Printf("Score: %d\n", score)
	board.Print(piece)
	println("Game Over!")
}
