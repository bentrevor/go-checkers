package checkers

import (
	"fmt"
)

type Game struct {
	Board         *Board
	CurrentPlayer Player
	OtherPlayer   Player
}

func NewGame() Game {
	return Game{
		Board:         NewGameBoard(),
		CurrentPlayer: NewPlayer("white"),
		OtherPlayer:   NewPlayer("black"),
	}
}

func NewGameWithBoard(board *Board) Game {
	return Game{
		Board:         board,
		CurrentPlayer: NewPlayer("white"),
		OtherPlayer:   NewPlayer("black"),
	}
}

func (game *Game) NextTurn() {
	move := game.CurrentPlayer.GetMove(game.Board)
	game.Board.MakeMove(move)
	game.togglePlayers()
}

func (game *Game) Start() {
	for !gameOver(game.Board) {
		fmt.Println("asdf")
		game.NextTurn()
	}
}

func (game *Game) CurrentColor() string {
	return game.CurrentPlayer.Color()
}

func (game *Game) togglePlayers() {
	p := game.CurrentPlayer
	game.CurrentPlayer = game.OtherPlayer
	game.OtherPlayer = p
}
