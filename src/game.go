package checkers

import "fmt"

type Game struct {
	Board         Board
	CurrentPlayer Player
	OtherPlayer   Player
	Output        Output
}

type IO interface {
	PrintBoard(Board)
	GetInput() (string, error)
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
	for !IsGameOver(game.Board) {
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
