package checkers

import (
	"fmt"
	// "strconv"
	// "bytes"
)

type Game struct {
	board        *Board
	CurrentColor string
}

func NewGame() Game {
	return Game{CurrentColor: "white"}
}

func (game *Game) NextTurn() {
	game.toggleCurrentColor()
}

func (game *Game) Start() {
	for !gameOver(game.board) {
		fmt.Println("asdf")
		game.NextTurn()
	}
}

func (game *Game) toggleCurrentColor() {
	if game.CurrentColor == "white" {
		game.CurrentColor = "black"
	} else {
		game.CurrentColor = "white"
	}
}
