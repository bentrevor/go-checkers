package checkers_test

import (
	. "github.com/bentrevor/checkers"
	"fmt"
	"testing"
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

func TestBoard_KnowsWhereAPieceCanMove(t *testing.T) {
	whitePiece := board.GetPieceAtSpace(NewSpace("g3"))
	blackPiece := board.GetPieceAtSpace(NewSpace("h6"))

	whiteMove1 := NewSpace("f4")
	whiteMove2 := NewSpace("h4")
	blackMove := NewSpace("g5")
	whiteMoves := []Move{
		Move{StartingSpace: NewSpace("g3"), TargetSpace: whiteMove1},
		Move{StartingSpace: NewSpace("g3"), TargetSpace: whiteMove2},
	}

	blackMoves := []Move{
		Move{StartingSpace: NewSpace("h6"), TargetSpace: blackMove},
	}

	assertEquals(t, whiteMoves, board.MovesForPiece(whitePiece))
	assertEquals(t, blackMoves, board.MovesForPiece(blackPiece))
}

func TestBoard_KnowsWhereAPieceCanJump(t *testing.T) {
	d6 := NewSpace("d6")
	d6Piece := board.GetPieceAtSpace(d6)
	f6 := NewSpace("f6")
	f6Piece := board.GetPieceAtSpace(f6)

	emptySpace := NewSpace("e5")
	whitePiece := Piece{Color: "white", Space: emptySpace}

	board.PlacePiece(whitePiece)

	captureMoveForD6 := Move{
		StartingSpace: d6,
		TargetSpace: NewSpace("f4"),
	}

	captureMoveForF6 := Move{
		StartingSpace: f6,
		TargetSpace: NewSpace("d4"),
	}

	d6Moves := board.MovesForPiece(d6Piece)
	f6Moves := board.MovesForPiece(f6Piece)

	assert(t, IncludesMove(d6Moves, captureMoveForD6), "d6 capture move")
	assert(t, IncludesMove(f6Moves, captureMoveForF6), "f6 capture move")
}
