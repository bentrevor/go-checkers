package checkers

import "fmt"

type Space struct {
	Rank int
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

func (board *Board) GetPieceAt(space Space) Piece {
	for _, piece := range board.Pieces {
		if sameSpace(piece.Space, space) {
			return piece
		}
	}

	return Piece{}
}

func (board *Board) MovesForPiece(piece Piece) []Space {
	space := piece.Space

	return board.movesForSpace(space, piece.Color)
}

// private

func (board *Board) movesForSpace(space Space, color string) []Space {
	moves := []Space{}
	nextRank := 0
	if color == "white" {
		nextRank = space.Rank + 1
	} else {
		nextRank = space.Rank - 1
	}

	if notOnLeftEdge(space) {
		leftFile := decFile(space.File)
		moves = append(moves, Space{File: leftFile, Rank: nextRank})
	}

	if notOnRightEdge(space) {
		rightFile := incFile(space.File)
		moves = append(moves, Space{File: rightFile, Rank: nextRank})
	}

	return moves
}

func notOnLeftEdge(space Space) bool {
	return space.File != "a"
}

func notOnRightEdge(space Space) bool {
	return space.File != "h"
}

func incFile(file string) string {
	return string(file[0] + 1)
}

func decFile(file string) string {
	return string(file[0] - 1)
}

func (board *Board) createPieceAtIndex(pieceIndex int, spaceIndex int) (Piece, bool) {
	if initialPiece, ok := initialPieceAtIndex(spaceIndex); ok {
		board.Pieces[pieceIndex] = initialPiece
		return initialPiece, true
	}

	return Piece{}, false
}

func sameSpace(pieceSpace Space, targetSpace Space) bool {
	return pieceSpace.Rank == targetSpace.Rank &&
		pieceSpace.File == targetSpace.File
}

func initialPieceAtIndex(index int) (Piece, bool) {
	color := colorFor(index)

	if index % 2 == 1 || color == "" {
		return Piece{}, false
	} else {
		piece := Piece{Color: color, Space: spaceFor(index)}
		return piece, true
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

	return Space{File: file, Rank: rank}
}
