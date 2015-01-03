package main

import (
	"flag"

	. "github.com/bentrevor/checkers/src"
)

var (
	flagFen = ""
)

func init() {
	flag.StringVar(&flagFen, "fen", flagFen, "initial board state")

	flag.Parse()
}

func main() {
	player1 := NewHumanPlayer("white")
	player2 := NewHumanPlayer("black")
	game := NewGame(player1, player2, ConsoleOutput{})

	if len(flagFen) > 0 {
		game.InitFromFen(flagFen)
	}

	game.Start()
}
