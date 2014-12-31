package main

import . "github.com/bentrevor/checkers/src"

func main() {
	player1 := NewHumanPlayer("white")
	player2 := NewHumanPlayer("black")
	game := NewGame(player1, player2, ConsoleOutput{})
	game.Start()
}
