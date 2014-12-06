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
		created := board.createPieceAtIndex(pieceIndex, i)

		if created && pieceIndex < 23 {
			pieceIndex += 1
		}
	}

	return board
}

func (board *Board) createPieceAtIndex(pieceIndex int, spaceIndex int) bool {
	initialPiece, ok := initialPieceAtIndex(spaceIndex)

	if ok {
		board.Pieces[pieceIndex] = initialPiece
	}

	return ok
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
	if pieceSpace.File == targetSpace.File {
		return pieceSpace.Rank == targetSpace.Rank
	}

	return false
}

func initialPieceAtIndex(index int) (Piece, bool) {
	if index % 2 == 1 {
		return Piece{}, false
	} else {
		rank := (index / 8) + 1
		file := ""

		if rank % 2 == 1 {
			file = string((index % 8) + 97)
		} else {
			file = string((index % 8) + 98)
		}

		color := ""

		if index < 24 {
			color = "white"
		} else if index > 39 {
			color = "black"
		} else {
			return Piece{}, false
		}

		return Piece{Color: color, Space: Space{File: file, Rank: strconv.Itoa(rank)}}, true
	}
}
