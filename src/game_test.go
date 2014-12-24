package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

const fakeInput = "c3 - d4"

var fakeIO = MockIO{}
var game = NewGame(fakeIO)

type MockPlayer struct {
	color string
}

type MockIO struct{}

func (MockIO) GetInput() (string, error) {
	return fakeInput, nil
}

func (MockIO) PrintBoard(b Board) {}

func (MockPlayer) GetMove(board Board) Move {
	return Move{StartingSpace: C3, TargetSpace: D4}
}

func (MockPlayer) Color() string {
	return "white"
}

func NewMockPlayer() Player {
	return &MockPlayer{color: "white"}
}

func TestGame_WhiteGoesFirst(t *testing.T) {
	assertEquals(t, "white", game.CurrentColor())
}

func TestGame_TogglesPlayers(t *testing.T) {
	game.CurrentPlayer = NewMockPlayer()
	game.Board = NewGameBoard()
	game.NextTurn()

	assertEquals(t, "black", game.CurrentColor())
}

func TestGame_MakesMovesOnTheBoard(t *testing.T) {
	game.CurrentPlayer = NewMockPlayer()

	assertEquals(t, "white", game.Board.GetPieceAtSpace(D4).Color)
	assertEquals(t, "", game.Board.GetPieceAtSpace(C3).Color)
}

func TestGame_ValidatesMoves(t *testing.T) {
	assert(t, game.InvalidInput(Move{B4, C5}), "invalid move - no piece at starting space")
}
