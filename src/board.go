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

		if created {
			pieceIndex += 1
		}
	}

	return board
}

func (board *Board) createPieceAtIndex(pieceIndex int, spaceIndex int) bool {
	board.Pieces[pieceIndex] = initialPieceAtIndex(spaceIndex)

	return pieceIndex < 23
}

func (board *Board) GetPieceAt(space Space) Piece {
	for _, piece := range board.Pieces {
		// fmt.Println(piece, piece.Space, space)

		if sameSpace(piece.Space, space) {
			return piece
		}
	}

	return Piece{Color: "NOOOOOOOOOOOOOOOOO!!!", Space: Space{File: "", Rank: ""}}
}

func sameSpace(pieceSpace Space, targetSpace Space) bool {
	// fmt.Println(pieceSpace.File, targetSpace.File)
	if pieceSpace.File == targetSpace.File {
		return pieceSpace.Rank == targetSpace.Rank
	}

	return false
}

// func (board *Board) MovesFor(space Space) []string {
// 	var strings []string
// 	return strings
// }

func initialPieceAtIndex(index int) Piece {
	file := string((index % 8) + 97)
	rank := strconv.Itoa((index / 8) + 1)
	color := "black"

	if index < 24 && index % 2 == 0 {
		color = "white"
	}

	// fmt.Println(index, ":", rank, file, ":", color)
	return Piece{Color: color, Space: Space{File: file, Rank: rank}}
}
