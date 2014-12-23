package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers"
)

func TestRules_KnowsWhereAPieceCanMove(t *testing.T) {
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

	assertEquals(t, whiteMoves, MovesForPiece(whitePiece, board))
	assertEquals(t, blackMoves, MovesForPiece(blackPiece, board))
}

func TestRules_KnowsWhereAPieceCanJump(t *testing.T) {
	d6 := NewSpace("d6")
	d6Piece := board.GetPieceAtSpace(d6)
	f6 := NewSpace("f6")
	f6Piece := board.GetPieceAtSpace(f6)

	emptySpace := NewSpace("e5")
	whitePiece := Piece{Color: "white", Space: emptySpace}

	board.PlacePiece(whitePiece)

	captureMoveForD6 := Move{
		StartingSpace: d6,
		TargetSpace:   NewSpace("f4"),
	}

	captureMoveForF6 := Move{
		StartingSpace: f6,
		TargetSpace:   NewSpace("d4"),
	}

	d6Moves := MovesForPiece(d6Piece, board)
	f6Moves := MovesForPiece(f6Piece, board)

	assert(t, IncludesMove(d6Moves, captureMoveForD6), "d6 capture move")
	assert(t, IncludesMove(f6Moves, captureMoveForF6), "f6 capture move")
}
