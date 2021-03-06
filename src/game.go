package checkers

import "fmt"

type Game struct {
	Board         Board
	CurrentPlayer Player
	OtherPlayer   Player
	Output        Output
	Rules         Rules
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
		Rules:         CheckersRules{},
	}
}

func (game *Game) NextTurn() {
	game.Output.PrintBoard(game.Board)

	move := game.CurrentPlayer.GetMove(game.Board, game.Rules)
	for !game.Rules.IsLegalMove(move, game.Board, game.CurrentColor()) {
		move = game.CurrentPlayer.GetMove(game.Board, game.Rules)
	}

	game.Board.MakeMove(move)
	game.togglePlayers()
}

func (game *Game) Start() {
	for !game.Rules.IsGameOver(game.Board) {
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
