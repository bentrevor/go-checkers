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

var A1 = NewSpace("a1")
var A2 = NewSpace("a2")
var A3 = NewSpace("a3")
var A4 = NewSpace("a4")
var A5 = NewSpace("a5")
var A6 = NewSpace("a6")
var A7 = NewSpace("a7")
var A8 = NewSpace("a8")
var B1 = NewSpace("b1")
var B2 = NewSpace("b2")
var B3 = NewSpace("b3")
var B4 = NewSpace("b4")
var B5 = NewSpace("b5")
var B6 = NewSpace("b6")
var B7 = NewSpace("b7")
var B8 = NewSpace("b8")
var C1 = NewSpace("c1")
var C2 = NewSpace("c2")
var C3 = NewSpace("c3")
var C4 = NewSpace("c4")
var C5 = NewSpace("c5")
var C6 = NewSpace("c6")
var C7 = NewSpace("c7")
var C8 = NewSpace("c8")
var D1 = NewSpace("d1")
var D2 = NewSpace("d2")
var D3 = NewSpace("d3")
var D4 = NewSpace("d4")
var D5 = NewSpace("d5")
var D6 = NewSpace("d6")
var D7 = NewSpace("d7")
var D8 = NewSpace("d8")
var E1 = NewSpace("e1")
var E2 = NewSpace("e2")
var E3 = NewSpace("e3")
var E4 = NewSpace("e4")
var E5 = NewSpace("e5")
var E6 = NewSpace("e6")
var E7 = NewSpace("e7")
var E8 = NewSpace("e8")
var F1 = NewSpace("f1")
var F2 = NewSpace("f2")
var F3 = NewSpace("f3")
var F4 = NewSpace("f4")
var F5 = NewSpace("f5")
var F6 = NewSpace("f6")
var F7 = NewSpace("f7")
var F8 = NewSpace("f8")
var G1 = NewSpace("g1")
var G2 = NewSpace("g2")
var G3 = NewSpace("g3")
var G4 = NewSpace("g4")
var G5 = NewSpace("g5")
var G6 = NewSpace("g6")
var G7 = NewSpace("g7")
var G8 = NewSpace("g8")
var H1 = NewSpace("h1")
var H2 = NewSpace("h2")
var H3 = NewSpace("h3")
var H4 = NewSpace("h4")
var H5 = NewSpace("h5")
var H6 = NewSpace("h6")
var H7 = NewSpace("h7")
var H8 = NewSpace("h8")

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
