package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers"
)

var game = NewGame()

type MockPlayer struct {
	color string
}

func (*MockPlayer) GetMove(board *Board) Move {
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
	game.NextTurn()

	assertEquals(t, "black", game.CurrentColor())
}

func TestGame_MakesMovesOnTheBoard(t *testing.T) {
	game.CurrentPlayer = NewMockPlayer()
	game.NextTurn()

	assertEquals(t, "white", game.Board.GetPieceAtSpace(D4).Color)
	assertEquals(t, "", game.Board.GetPieceAtSpace(C3).Color)
}
