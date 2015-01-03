package checkers

import (
	"regexp"
	"strings"
)

type Piece struct {
	Color  Color
	Space  Space
	IsKing bool
}

type Board struct {
	Pieces []Piece
}

func NewEmptyBoard() Board {
	return Board{make([]Piece, 24, 32)}
}

func NewGameBoard() Board {
	board := NewEmptyBoard()
	board.placeInitialPieces()
	return board
}

func BoardFromFen(fen string) Board {
	pieces := strings.Split(fen, " ")[0]
	board := NewEmptyBoard()

	for i, fenRow := range strings.Split(pieces, "/") {
		rank := i + 1
		spaces := spacesForRank(rank)

		rowPieces := getRowPieces(fenRow)

		for pieceIndex, fenPiece := range rowPieces {
			var piece Piece
			space := spaces[pieceIndex]

			if fenPiece == "w" {
				piece = Piece{Color: White, Space: space}
			} else if fenPiece == "b" {
				piece = Piece{Color: Black, Space: space}
			}

			board.PlacePiece(piece)
		}
	}
	return board
}

func getRowPieces(fenRow string) []string {
	return strings.Split(ExpandNumbers(fenRow), "")
}

func ExpandNumbers(row string) string {
	matched, err := regexp.MatchString("[1234]", row)

	if err != nil {
		// TODO
	}

	if matched {
		no4s := strings.Replace(row, "4", "1111", 1)
		no3s := strings.Replace(no4s, "3", "111", 1)
		no2s := strings.Replace(no3s, "2", "11", 1)
		return no2s
	} else {
		return row
	}
}

func (board *Board) GetPieceAtSpace(space Space) (Piece, bool) {
	for _, piece := range board.Pieces {
		if IsSameSpace(piece.Space, space) {
			return piece, true
		}
	}

	return Piece{}, false
}

func (board *Board) MakeMove(move Move) {
	piece, foundPiece := board.GetPieceAtSpace(move.StartingSpace)

	if !foundPiece {
		// TODO
	} else {
		rank := move.TargetSpace.Rank
		isKing := rank == 1 || rank == 8
		newPiece := Piece{Color: piece.Color, Space: move.TargetSpace, IsKing: isKing}
		board.PlacePiece(newPiece)
		board.RemovePieceAtSpace(move.StartingSpace)

		if isCaptureMove(move) {
			capturedSpace := capturedSpace(move)

			board.RemovePieceAtSpace(capturedSpace)
		}
	}
}

func (board *Board) RemovePieceAtSpace(space Space) {
	for i, piece := range board.Pieces {
		if IsSameSpace(piece.Space, space) {
			board.Pieces[i] = Piece{}
		}
	}
}

func (board *Board) PlacePiece(piece Piece) bool {
	_, pieceAtSpace := board.GetPieceAtSpace(piece.Space)

	if !pieceAtSpace {
		board.addPiece(piece)
		return true
	} else {
		return false
	}
}

func reverseRows(rows []string) []string {
	length := len(rows)
	reversed := make([]string, length)

	for i := 0; i < length; i++ {
		reversed[length-(i+1)] = rows[i]
	}

	return reversed
}

func (board *Board) placeInitialPieces() {
	pieceIndex := 0

	for i := 0; i < 64; i++ {
		_, created := board.createInitialPieceAtIndex(pieceIndex, i)

		if created && pieceIndex < cap(board.Pieces) {
			pieceIndex += 1
		}
	}
}

func (board *Board) createInitialPieceAtIndex(pieceIndex int, spaceIndex int) (Piece, bool) {
	if initialPiece, ok := initialPieceAtIndex(spaceIndex); ok {
		board.Pieces[pieceIndex] = initialPiece
		return initialPiece, true
	}

	return Piece{}, false
}

func initialPieceAtIndex(index int) (Piece, bool) {
	pieceColor := initialPieceColorFor(index)
	spaceColor := SpaceColorForIndex(index)

	if pieceColor == NoColor || spaceColor == White {
		return Piece{}, false
	} else {
		piece := Piece{Color: pieceColor, Space: SpaceForIndex(index)}
		return piece, true
	}
}

func (board *Board) addPiece(piece Piece) {
	board.Pieces = append(board.Pieces, piece)
}
