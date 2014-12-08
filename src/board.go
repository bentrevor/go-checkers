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
	board := &Board{make([]Piece, 24, 32)}
	pieceIndex := 0

	for i := 0; i < 64; i++ {
		_, created := board.createPieceAtIndex(pieceIndex, i)

		if created && pieceIndex < cap(board.Pieces) {
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

func (board *Board) PlacePiece(piece Piece) (Piece, bool) {
	if board.GetPieceAt(piece.Space).Color == "" {
		board.addPiece(piece)
		return piece, true
	} else {
		return piece, false
	}
}

func (board *Board) MovesForPiece(piece Piece) []Space {
	space := piece.Space

	return board.movesForSpace(space, piece.Color)
}

func Includes(moves []Space, move Space) bool {
	for _, space := range moves {
		if sameSpace(space, move) {
			return true
		}
	}

	return false
}

// private

func (board *Board) addPiece(piece Piece) {
	board.Pieces = append(board.Pieces, piece)
}

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
		targetSpace := Space{File: leftFile, Rank: nextRank}
		if nextMove, ok := board.getNextMove(space, targetSpace); ok {
			moves = append(moves, nextMove)
		}
	}

	if notOnRightEdge(space) {
		rightFile := incFile(space.File)
		targetSpace := Space{File: rightFile, Rank: nextRank}

		if nextMove, ok := board.getNextMove(space, targetSpace); ok {
			moves = append(moves, nextMove)
		}
	}

	return moves
}

func (board *Board) getNextMove(startingSpace Space, targetSpace Space) (Space, bool) {
	// fmt.Println("target space: ", targetSpace, "\npiece: ", board.GetPieceAt(targetSpace), "\n\n")
	if board.GetPieceAt(targetSpace).Color == "" {
		return targetSpace, true
	} else {
		nextSpace := getNextSpace(startingSpace, targetSpace)
		if board.GetPieceAt(nextSpace).Color == "" {
			return nextSpace, true
		} else {
			return Space{}, false
		}
	}
}

func getNextSpace(startingSpace Space, targetSpace Space) Space {
	increasingRank := startingSpace.Rank < targetSpace.Rank
	increasingFile := startingSpace.File[0] < targetSpace.File[0]
	nextRank := 0
	nextFile := ""

	if increasingRank {
		nextRank = targetSpace.Rank + 1
	} else {
		nextRank = targetSpace.Rank - 1
	}

	if increasingFile {
		nextFile = string(targetSpace.File[0] + 1)
	} else {
		nextFile = string(targetSpace.File[0] - 1)
	}

	return Space{File: nextFile, Rank: nextRank}
}

func (board *Board) createPieceAtIndex(pieceIndex int, spaceIndex int) (Piece, bool) {
	if initialPiece, ok := initialPieceAtIndex(spaceIndex); ok {
		board.Pieces[pieceIndex] = initialPiece
		return initialPiece, true
	}

	return Piece{}, false
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
