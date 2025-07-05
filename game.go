package main

type Game struct {
	Score  int
	Paused bool
	Over   bool
}

func NewGame() *Game {
	return &Game{Score: 0, Paused: false, Over: false}
}
