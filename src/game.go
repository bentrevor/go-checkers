package checkers

import (
	"fmt"
	"strings"
)

type Game struct {
	Board         Board
	CurrentPlayer Player
	OtherPlayer   Player
	Output        Output
}

type Input interface {
	GetInput() (string, error)
}

type Output interface {
	PrintBoard(Board)
}

func NewGame(player1 Player, player2 Player, output Output) Game {
	board := NewGameBoard()

	return Game{
		Board:         board,
		CurrentPlayer: player1,
		OtherPlayer:   player2,
		Output:        output,
	}
}

func (game *Game) InitFromFen(fen string) {
	board := BoardFromFen(fen)
	fenColor := strings.Split(fen, " ")[1]

	var color Color

	if fenColor == "b" {
		color = Black
	} else {
		color = White
	}

	if color != game.CurrentColor() {
		game.togglePlayers()
	}

	game.Board = board
}

func (game *Game) NextTurn() {
	game.Output.PrintBoard(game.Board)

	move := game.CurrentPlayer.GetMove(game.Board)
	for !game.IsValidMove(move) {
		move = game.CurrentPlayer.GetMove(game.Board)
	}

	game.Board.MakeMove(move)
	game.togglePlayers()
}

func (game *Game) IsValidMove(move Move) bool {
	board := game.Board
	color := game.CurrentPlayer.Color()

	return IsLegalMove(move, board, color)
}

func (game *Game) Start() {
	for !game.Board.IsGameOver() {
		game.NextTurn()
	}

	game.Output.PrintBoard(game.Board)
	fmt.Println("game over!")
}

func (game *Game) CurrentColor() Color {
	return game.CurrentPlayer.Color()
}

func (game *Game) togglePlayers() {
	playerThatJustMoved := game.CurrentPlayer
	game.CurrentPlayer = game.OtherPlayer
	game.OtherPlayer = playerThatJustMoved
}
