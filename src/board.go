package checkers

import "fmt"
import "strconv"

type Space struct {
	Rank string
	File string
}

type Piece struct {
	Color string
	Space Space
}

type Board struct {
	Pieces []Piece
}

func dummy() {
	fmt.Println()
}

func NewGameBoard() *Board {
	board := &Board{make([]Piece, 24)}
	pieceIndex := 0

	for i := 0; i < 64; i++ {
		_, created := board.createPieceAtIndex(pieceIndex, i)

		if created && pieceIndex < 23 {
			pieceIndex += 1
		}
	}

	return board
}

func (board *Board) createPieceAtIndex(pieceIndex int, spaceIndex int) (Piece, bool) {
	if initialPiece, ok := initialPieceAtIndex(spaceIndex); ok {
		board.Pieces[pieceIndex] = initialPiece
		return initialPiece, true
	}

	return Piece{}, false
}

func (board *Board) GetPieceAt(space Space) Piece {
	for _, piece := range board.Pieces {
		if sameSpace(piece.Space, space) {
			return piece
		}
	}

	return Piece{Color: "NOOOOOOOOOOOOOOOOO!!!", Space: Space{File: "", Rank: ""}}
}

func sameSpace(pieceSpace Space, targetSpace Space) bool {
	return pieceSpace.Rank == targetSpace.Rank &&
		pieceSpace.File == targetSpace.File
}

func initialPieceAtIndex(index int) (Piece, bool) {
	if index % 2 == 1 {
		return Piece{}, false
	} else {
		return Piece{Color: colorFor(index), Space: spaceFor(index)}, true
	}
}

func colorFor(index int) string {
	if index < 24 {
		return "white"
	} else if index > 39 {
		return "black"
	} else {
		return ""
	}
}

func spaceFor(index int) Space {
	rank := (index / 8) + 1
	file := ""

	if rank % 2 == 1 {
		file = string((index % 8) + 97)
	} else {
		file = string((index % 8) + 98)
	}

	return Space{File: file, Rank: strconv.Itoa(rank)}
}
