package checkers

import (
	"bytes"
	"fmt"
	"strconv"
)

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

func NewEmptyBoard() *Board {
	return &Board{make([]Piece, 24, 32)}
}

func NewGameBoard() *Board {
	board := NewEmptyBoard()
	board.createInitialPieces()
	return board
}

func NewSpace(coordinates string) Space {
	file := string(coordinates[0])
	rank, _ := strconv.Atoi(string(coordinates[1]))

	return Space{File: file, Rank: rank}
}

func (board *Board) GetPieceAtSpace(space Space) Piece {
	for _, piece := range board.Pieces {
		if SameSpace(piece.Space, space) {
			return piece
		}
	}

	return Piece{}
}

func (board *Board) PlacePiece(piece Piece) (Piece, bool) {
	if board.GetPieceAtSpace(piece.Space).Color == "" {
		board.addPiece(piece)
		return piece, true
	} else {
		return piece, false
	}
}

func (board *Board) ConsolePrint() {
	fmt.Println()
	spaces := []Space{}
	emptySpace := "|_|"
	var row bytes.Buffer
	rows := []string{}

	for _, piece := range board.Pieces {
		spaces = append(spaces, piece.Space)
	}

	for i := 0; i < 64; i++ {
		space := SpaceForIndex(i)
		piece := board.GetPieceAtSpace(space)

		if piece.Color == "" {
			row.WriteString(emptySpace)
		} else {
			printableSpace := fmt.Sprintf("|%c|", piece.Color[0])
			row.WriteString(printableSpace)
		}
		if i%8 == 7 {
			row.WriteString("\n")
			rows = append(rows, row.String())
			row.Reset()
		}
	}

	fmt.Println(reverseRows(rows))
	fmt.Println("  1  2  3  4  5  6  7  8")
	return
}

// private

func reverseRows(rows []string) []string {
	length := len(rows)
	reversed := make([]string, length)

	for i := 0; i < length; i++ {
		reversed[length-(i+1)] = rows[i]
	}

	return reversed
}

func (board *Board) createInitialPieces() {
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

	if pieceColor == "" || spaceColor == "white" {
		return Piece{}, false
	} else {
		piece := Piece{Color: pieceColor, Space: SpaceForIndex(index)}
		return piece, true
	}
}

func (board *Board) addPiece(piece Piece) {
	board.Pieces = append(board.Pieces, piece)
}
