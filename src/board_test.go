package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

var board = NewGameBoard()

func TestBoard_StartsWith24Pieces(t *testing.T) {
	assertEquals(t, 24, len(board.Pieces))
}

func TestBoard_SetsTheInitialPieceLayout(t *testing.T) {
	whiteSpaces := []Space{A1, C1, E1, G1, B2, D2, F2, H2, A3, C3, E3, G3}
	blackSpaces := []Space{B6, D6, F6, H6, A7, C7, E7, G7, B8, D8, F8, H8}

	for _, space := range whiteSpaces {
		piece, _ := board.GetPieceAtSpace(space)
		assertEquals(t, White, piece.Color)
	}

	for _, space := range blackSpaces {
		piece, _ := board.GetPieceAtSpace(space)
		assertEquals(t, Black, piece.Color)
	}
}

func TestBoard_CanPlaceAPiece(t *testing.T) {
	emptySpace := E5
	occupiedSpace := E7

	piece1 := Piece{Color: White, Space: emptySpace}
	piece2 := Piece{Color: White, Space: occupiedSpace}

	createdPieceAtE5 := board.PlacePiece(piece1)
	createdPieceAtE7 := board.PlacePiece(piece2)

	assert(t, createdPieceAtE5, "should have placed piece at e5")
	assert(t, !createdPieceAtE7, "should not have placed piece at e5")
}

func TestBoard_CanRemovePieces(t *testing.T) {
	board.RemovePieceAtSpace(G3)
	g3Piece, _ := board.GetPieceAtSpace(G3)

	assertEquals(t, NoColor, g3Piece.Color)
}

func TestBoard_CanMakeMoves(t *testing.T) {
	board := Board{}
	whitePiece := Piece{Color: White, Space: G3}
	whitePiece2 := Piece{Color: White, Space: C7}
	board.PlacePiece(whitePiece)
	board.PlacePiece(whitePiece2)

	move, _ := MoveFromString("g3 - h4")

	board.MakeMove(move)

	h4Piece, _ := board.GetPieceAtSpace(H4)
	g3Piece, _ := board.GetPieceAtSpace(G3)

	assertEquals(t, White, h4Piece.Color)
	assertEquals(t, NoColor, g3Piece.Color)

	blackPiece := Piece{Color: Black, Space: G5}
	board.PlacePiece(blackPiece)

	captureMove, _ := MoveFromString("h4 - f6")
	board.MakeMove(captureMove)

	g5Piece, _ := board.GetPieceAtSpace(G5)

	assertEquals(t, NoColor, g5Piece.Color)

	promotionMove, _ := MoveFromString("c7 - d8")
	board.MakeMove(promotionMove)
	king, _ := board.GetPieceAtSpace(D8)

	assert(t, king.IsKing, "MakeMove promotion")
}

func TestBoard_CanBeCreatedFromFen(t *testing.T) {
	fen := "w3/4/4/4/4/4/4/4 w"
	board := BoardFromFen(fen)

	a1Piece, foundPieceAtA1 := board.GetPieceAtSpace(A1)
	_, foundPieceAtG3 := board.GetPieceAtSpace(G3)

	assert(t, !foundPieceAtG3, "fen had no piece at g3")
	assert(t, foundPieceAtA1, "fen had a piece at a1")
	assertEquals(t, White, a1Piece.Color)
}

func TestBoard_CanBeCreatedFromFen2(t *testing.T) {
	fen := "bwbw/4/4/wwbb/4/4/4/4 b"
	board := BoardFromFen(fen)

	a1Piece, foundPieceAtA1 := board.GetPieceAtSpace(A1)
	h4Piece, foundPieceAtH4 := board.GetPieceAtSpace(H4)

	assert(t, foundPieceAtH4, "fen had piece at h4")
	assertEquals(t, Black, h4Piece.Color)
	assert(t, foundPieceAtA1, "fen had a piece at a1")
	assertEquals(t, Black, a1Piece.Color)
}

func TestBoard_CanExpandAFenRow(t *testing.T) {
	fenRow := "4"
	expandedRow := ExpandNumbers(fenRow)
	assertEquals(t, len(expandedRow), 4)
	assertEquals(t, expandedRow, "1111")

	fenRow = "3w"
	expandedRow = ExpandNumbers(fenRow)
	assertEquals(t, len(expandedRow), 4)
	assertEquals(t, expandedRow, "111w")
}
