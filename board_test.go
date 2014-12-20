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
		Space{File: "a", Rank: 1},
		Space{File: "c", Rank: 1},
		Space{File: "e", Rank: 1},
		Space{File: "g", Rank: 1},
		Space{File: "b", Rank: 2},
		Space{File: "d", Rank: 2},
		Space{File: "f", Rank: 2},
		Space{File: "h", Rank: 2},
		Space{File: "a", Rank: 3},
		Space{File: "c", Rank: 3},
		Space{File: "e", Rank: 3},
		Space{File: "g", Rank: 3},
	}

	blackSpaces := []Space{
		Space{File: "b", Rank: 6},
		Space{File: "d", Rank: 6},
		Space{File: "f", Rank: 6},
		Space{File: "h", Rank: 6},
		Space{File: "a", Rank: 7},
		Space{File: "c", Rank: 7},
		Space{File: "e", Rank: 7},
		Space{File: "g", Rank: 7},
		Space{File: "b", Rank: 8},
		Space{File: "d", Rank: 8},
		Space{File: "f", Rank: 8},
		Space{File: "h", Rank: 8},
	}

	for _, space := range whiteSpaces {
		assertEquals(t, "white", board.GetPieceAt(space).Color)
	}

	for _, space := range blackSpaces {
		assertEquals(t, "black", board.GetPieceAt(space).Color)
	}
}


func TestBoard_CanPlaceAPiece(t *testing.T) {
	emptySpace := Space{File: "e", Rank: 5}
	occupiedSpace := Space{File: "e", Rank: 7}

	piece1 := Piece{Color: "white", Space: emptySpace}
	piece2 := Piece{Color: "white", Space: occupiedSpace}

	_, createdPieceAtE5 := board.PlacePiece(piece1)
	_, createdPieceAtE7 := board.PlacePiece(piece2)

	assert(t, createdPieceAtE5, "should have placed piece at e5")
	assert(t, !createdPieceAtE7, "should not have placed piece at e5")
}

func TestBoard_KnowsWhereAPieceCanMove(t *testing.T) {
	whitePiece := board.GetPieceAt(Space{File: "g", Rank: 3})
	blackPiece := board.GetPieceAt(Space{File: "h", Rank: 6})

	whiteMove1 := Space{File: "f", Rank: 4}
	whiteMove2 := Space{File: "h", Rank: 4}
	blackMove := Space{File: "g", Rank: 5}
	whiteMoves := []Move{
		Move{StartingSpace: Space{File: "g", Rank: 3}, TargetSpace: whiteMove1},
		Move{StartingSpace: Space{File: "g", Rank: 3}, TargetSpace: whiteMove2},
	}

	blackMoves := []Move{
		Move{StartingSpace: Space{File: "h", Rank: 6}, TargetSpace: blackMove},
	}

	assertEquals(t, whiteMoves, board.MovesForPiece(whitePiece))
	assertEquals(t, blackMoves, board.MovesForPiece(blackPiece))
}

func TestBoard_KnowsWhereAPieceCanJump(t *testing.T) {
	d6 := Space{File: "d", Rank: 6}
	d6Piece := board.GetPieceAt(d6)
	f6 := Space{File: "f", Rank: 6}
	f6Piece := board.GetPieceAt(f6)

	emptySpace := Space{File: "e", Rank: 5}
	whitePiece := Piece{Color: "white", Space: emptySpace}

	board.PlacePiece(whitePiece)

	captureMoveForD6 := Move{
		StartingSpace: d6,
		TargetSpace: Space{File: "f", Rank: 4},
	}

	captureMoveForF6 := Move{
		StartingSpace: f6,
		TargetSpace: Space{File: "d", Rank: 4},
	}

	d6Moves := board.MovesForPiece(d6Piece)
	f6Moves := board.MovesForPiece(f6Piece)

	assert(t, IncludesMove(d6Moves, captureMoveForD6), "d6 capture move")
	assert(t, IncludesMove(f6Moves, captureMoveForF6), "f6 capture move")
}
