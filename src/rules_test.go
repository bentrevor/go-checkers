package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

func TestRules_KnowsWhereAPieceCanMove(t *testing.T) {
	board := NewGameBoard()
	whitePiece, _ := board.GetPieceAtSpace(G3)
	blackPiece, _ := board.GetPieceAtSpace(H6)

	whiteMoves := []Move{
		Move{StartingSpace: G3, TargetSpace: F4},
		Move{StartingSpace: G3, TargetSpace: H4},
	}

	blackMoves := []Move{
		Move{StartingSpace: H6, TargetSpace: G5},
	}

	assertEquals(t, whiteMoves, MovesForPiece(whitePiece, board))
	assertEquals(t, blackMoves, MovesForPiece(blackPiece, board))
}

func TestRules_KnowsWhereAPieceCanJump(t *testing.T) {
	d6Piece, _ := board.GetPieceAtSpace(D6)
	f6Piece, _ := board.GetPieceAtSpace(F6)

	board.PlacePiece(Piece{Color: White, Space: E5})

	d6Moves := MovesForPiece(d6Piece, board)
	f6Moves := MovesForPiece(f6Piece, board)

	captureMoveForD6 := Move{StartingSpace: D6, TargetSpace: F4}
	captureMoveForF6 := Move{StartingSpace: F6, TargetSpace: D4}

	assert(t, IncludesMove(d6Moves, captureMoveForD6), "d6 capture move")
	assert(t, IncludesMove(f6Moves, captureMoveForF6), "f6 capture move")
}

func TestRules_KnowsWhenAMoveIsValid(t *testing.T) {
	move := Move{StartingSpace: A3, TargetSpace: B4}
	assert(t, IsLegalMove(move, board, White), "valid white noncapture move")

	move = Move{StartingSpace: B6, TargetSpace: C5}
	assert(t, IsLegalMove(move, board, Black), "valid black noncapture move")

	board.PlacePiece(Piece{Color: Black, Space: F4})
	move = Move{StartingSpace: E3, TargetSpace: G5}
	assert(t, IsLegalMove(move, board, White), "valid white capture move")
}

func TestRules_KnowsWhenAMoveIsInvalid(t *testing.T) {
	move := Move{StartingSpace: A3, TargetSpace: B4}
	assert(t, !IsLegalMove(move, board, Black), "invalid black move - wrong color")

	move = Move{StartingSpace: A3, TargetSpace: A3}
	assert(t, !IsLegalMove(move, board, White), "invalid white move - not a real move")

	board.PlacePiece(Piece{Color: Black, Space: F4})
	move = Move{StartingSpace: E3, TargetSpace: F4}
	assert(t, !IsLegalMove(move, board, White), "invalid white move - moving into an occupied square")

	move = Move{StartingSpace: F4, TargetSpace: E3}
	assert(t, !IsLegalMove(move, board, Black), "invalid black move - moving into an occupied square")
}

func TestRules_ConvertsStringToMove(t *testing.T) {
	input := "a3 - b4"
	move := Move{StartingSpace: A3, TargetSpace: B4}

	parsedMove, _ := MoveFromString(input)
	assert(t, IsSameMove(move, parsedMove), "valid move from string")
}

func TestRules_ValidatesInputFormat(t *testing.T) {
	_, wrongLengthErrorMessage := MoveFromString("a3-b4")
	assert(t, (len(wrongLengthErrorMessage) > 0), "invalid move from string - wrong input length")

	_, nonsenseInputErrorMessage := MoveFromString("xx - yy")
	assert(t, (len(nonsenseInputErrorMessage) > 0), "invalid move from string - garbage input")
}
