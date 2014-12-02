package checkers

import "strconv"

type Space struct {
	Rank string
	File string
}

type Piece struct {
	color string
	space Space
}

type Board struct {
	Pieces [24]Piece
}

func NewGameBoard() *Board {
	board := new(Board)
	index := 0

	for i := 1; i < 9; i++ {
		for j := 97; j < 105; j++ {
			color := ""
			if index < 12 {
				color = "white"
			} else {
				color = "black"
			}

			board.Pieces[index] = Piece{color: color, space: Space{File: string(j), Rank: strconv.Itoa(i)}}
		}
	}

	return board
}

// func (board *Board) MovesFor(space Space) []string {
// 	var strings []string
// 	return strings
// }
