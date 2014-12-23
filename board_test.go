package checkers_test

import (
	"fmt"
	"testing"

	. "github.com/bentrevor/checkers"
)

var board = NewGameBoard()

func TestBoard_StartsWith24Pieces(t *testing.T) {
	fmt.Println()
	assertEquals(t, 24, len(board.Pieces))
}

func TestBoard_SetsTheInitialPieceLayout(t *testing.T) {
	whiteSpaces := []Space{
		NewSpace("a1"),
		NewSpace("c1"),
		NewSpace("e1"),
		NewSpace("g1"),
		NewSpace("b2"),
		NewSpace("d2"),
		NewSpace("f2"),
		NewSpace("h2"),
		NewSpace("a3"),
		NewSpace("c3"),
		NewSpace("e3"),
		NewSpace("g3"),
	}

	blackSpaces := []Space{
		NewSpace("b6"),
		NewSpace("d6"),
		NewSpace("f6"),
		NewSpace("h6"),
		NewSpace("a7"),
		NewSpace("c7"),
		NewSpace("e7"),
		NewSpace("g7"),
		NewSpace("b8"),
		NewSpace("d8"),
		NewSpace("f8"),
		NewSpace("h8"),
	}

	for _, space := range whiteSpaces {
		assertEquals(t, "white", board.GetPieceAtSpace(space).Color)
	}

	for _, space := range blackSpaces {
		assertEquals(t, "black", board.GetPieceAtSpace(space).Color)
	}
}

func TestBoard_CanPlaceAPiece(t *testing.T) {
	emptySpace := NewSpace("e5")
	occupiedSpace := NewSpace("e7")

	piece1 := Piece{Color: "white", Space: emptySpace}
	piece2 := Piece{Color: "white", Space: occupiedSpace}

	_, createdPieceAtE5 := board.PlacePiece(piece1)
	_, createdPieceAtE7 := board.PlacePiece(piece2)

	assert(t, createdPieceAtE5, "should have placed piece at e5")
	assert(t, !createdPieceAtE7, "should not have placed piece at e5")
}
