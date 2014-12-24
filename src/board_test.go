package checkers_test

import (
	"fmt"
	"testing"

	. "github.com/bentrevor/checkers/src"
)

var board = NewGameBoard()

func TestBoard_StartsWith24Pieces(t *testing.T) {
	fmt.Println()
	assertEquals(t, 24, len(board.Pieces))
}

func TestBoard_SetsTheInitialPieceLayout(t *testing.T) {
	whiteSpaces := []Space{A1, C1, E1, G1, B2, D2, F2, H2, A3, C3, E3, G3}
	blackSpaces := []Space{B6, D6, F6, H6, A7, C7, E7, G7, B8, D8, F8, H8}

	for _, space := range whiteSpaces {
		assertEquals(t, "white", board.GetPieceAtSpace(space).Color)
	}

	for _, space := range blackSpaces {
		assertEquals(t, "black", board.GetPieceAtSpace(space).Color)
	}
}

func TestBoard_CanPlaceAPiece(t *testing.T) {
	emptySpace := E5
	occupiedSpace := E7

	piece1 := Piece{Color: "white", Space: emptySpace}
	piece2 := Piece{Color: "white", Space: occupiedSpace}

	_, createdPieceAtE5 := board.PlacePiece(piece1)
	_, createdPieceAtE7 := board.PlacePiece(piece2)

	assert(t, createdPieceAtE5, "should have placed piece at e5")
	assert(t, !createdPieceAtE7, "should not have placed piece at e5")
}

func TestBoard_CanRemovePieces(t *testing.T) {
	board.RemovePieceAtSpace(G3)

	assertEquals(t, "", board.GetPieceAtSpace(G3).Color)
}

func TestBoard_CanMakeMoves(t *testing.T) {
	board := Board{}
	piece := Piece{Color: "white", Space: G3}
	board.PlacePiece(piece)
	move := Move{StartingSpace: G3, TargetSpace: H4}

	board.MakeMove(move)

	assertEquals(t, "white", board.GetPieceAtSpace(H4).Color)
	assertEquals(t, "", board.GetPieceAtSpace(G3).Color)
}
