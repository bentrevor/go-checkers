package checkers

import "strings"

type Piece struct {
	Color Color
	Space Space
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

		if fenRow == "4" {
			// no pieces
		} else {
			pieceIndex := strings.Index(fenRow, "w")
			color := White

			if pieceIndex == -1 {
				pieceIndex = strings.Index(fenRow, "b")
				color = Black

				if pieceIndex == -1 {
					// TODO
				}
			}

			space := spaces[pieceIndex]

			piece := Piece{Space: space, Color: color}
			board.PlacePiece(piece)
		}
	}
	return board
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
		newPiece := Piece{Color: piece.Color, Space: move.TargetSpace}
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
