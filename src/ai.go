package checkers

import "math/rand"

type RandomMove struct{}

func (RandomMove) GetMove(board Board, rules Rules, color Color) Move {
	moves := []Move{}

	for _, piece := range board.Pieces {
		if piece.Color == color {
			moves = append(moves, rules.MovesForPiece(piece, board)...)
		}
	}

	return moves[rand.Intn(len(moves))]
}
