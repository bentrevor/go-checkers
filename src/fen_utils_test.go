package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

func TestGame_InitFromFen(t *testing.T) {
	game := NewGameWithMockPlayers()
	game.InitFromFen("4/4/4/4/bw2/4/4/4 b")

	numPieces := 0

	for _, piece := range game.Board.Pieces {
		if piece.Color != NoColor {
			numPieces += 1
		}
	}

	assertEquals(t, Black, game.CurrentColor())
	assertEquals(t, 2, numPieces)
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

func TestBoard_CanCreateBlackPiecesFromFen(t *testing.T) {
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

func TestBoard_CanExportPiecesAsFen(t *testing.T) {
	board := NewGameBoard()

	assertEquals(t, board.PiecesToFen(), "wwww/wwww/wwww/4/4/bbbb/bbbb/bbbb")
}
