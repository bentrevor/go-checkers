package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

var game = NewGame()

type MockBoard struct{}
type MockPlayer struct {
	color string
}

func (MockBoard) ConsolePrint() {}

func (MockPlayer) GetMove(board Board) Move {
	return Move{StartingSpace: C3, TargetSpace: D4}
}

func (MockPlayer) Color() string {
	return "white"
}

func NewMockBoard() Board {
	return &MockBoard{}
}

func NewMockPlayer() Player {
	return &MockPlayer{color: "white"}
}

func TestGame_WhiteGoesFirst(t *testing.T) {
	assertEquals(t, "white", game.CurrentColor())
}

func TestGame_TogglesPlayers(t *testing.T) {
	game.CurrentPlayer = NewMockPlayer()
	game.Board = NewMockBoard()
	game.NextTurn()

	assertEquals(t, "black", game.CurrentColor())
}

func TestGame_MakesMovesOnTheBoard(t *testing.T) {
	game.CurrentPlayer = NewMockPlayer()

	assertEquals(t, "white", game.Board.GetPieceAtSpace(D4).Color)
	assertEquals(t, "", game.Board.GetPieceAtSpace(C3).Color)
}

func TestGame_ValidatesMoves(t *testing.T) {
	assert(t, game.InvalidInput(Move{B4, C5}), "invalid move - no piece at d4")
}
