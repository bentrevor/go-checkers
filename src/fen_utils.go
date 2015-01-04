package checkers

import (
	"fmt"
	"regexp"
	"strings"
)

func (game *Game) InitFromFen(fen string) {
	board := BoardFromFen(fen)
	fenColor := strings.Split(fen, " ")[1]

	var color Color

	if fenColor == "b" {
		color = Black
	} else {
		color = White
	}

	if color != game.CurrentColor() {
		game.togglePlayers()
	}

	game.Board = board
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

func (board Board) PiecesToFen() string {
	piecesFen := ""

	rows := BlackSpaces()

	for i, row := range rows {
		for _, space := range row {
			piece, _ := board.GetPieceAtSpace(space)
			abbrev := getPieceAbbrev(piece)
			piecesFen += abbrev
		}

		if i < len(rows)-1 {
			piecesFen += "/"
		}
	}

	return compactSpaces(piecesFen)
}

func ExpandNumbers(row string) string {
	matched, err := regexp.MatchString("[1234]", row)

	if err != nil {
		panic(fmt.Sprintf("couldn't match a string with row: %s", row))
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

func compactSpaces(piecesFen string) string {
	with4s := strings.Replace(piecesFen, "    ", "4", 8)
	with3s := strings.Replace(with4s, "   ", "3", 8)
	with2s := strings.Replace(with3s, "  ", "2", 8)

	return strings.Replace(with2s, " ", "1", 20)
}

func spacesForRank(rank int) []Space {
	return BlackSpaces()[rank-1]
}
