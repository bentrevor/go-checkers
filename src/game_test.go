package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

const fakeInput = "c3 - d4"

var player1 = NewHumanPlayer(White)
var player2 = NewHumanPlayer(Black)

var game = NewGame(player1, player2, MockOutput{})

type MockPlayer struct {
	color Color
}

type MockInput struct{}
type MockOutput struct{}

func (MockInput) GetInput() (string, error) {
	return fakeInput, nil
}

func (MockOutput) PrintBoard(b Board) {}

func (MockPlayer) GetMove(board Board) Move {
	return MoveFromString(fakeInput)
}

func (MockPlayer) Color() Color {
	return White
}

func NewMockPlayer() Player {
	return &MockPlayer{color: White}
}

func TestGame_WhiteGoesFirst(t *testing.T) {
	assertEquals(t, White, game.CurrentColor())
}

func TestGame_TogglesPlayers(t *testing.T) {
	assertEquals(t, White, game.CurrentColor())

	game.CurrentPlayer = NewMockPlayer()
	game.Board = NewGameBoard()
	game.NextTurn()

	assertEquals(t, Black, game.CurrentColor())
}

func TestGame_MakesMovesOnTheBoard(t *testing.T) {
	game.CurrentPlayer = NewMockPlayer()

	assertEquals(t, White, game.Board.GetPieceAtSpace(D4).Color)
	assertEquals(t, "", game.Board.GetPieceAtSpace(C3).Color)
}

func TestGame_ValidatesMoves(t *testing.T) {
	assert(t, game.InvalidInput(Move{B4, C5}), "invalid move - no piece at starting space")
}
