package checkers

type Move struct {
	StartingSpace Space
	TargetSpace Space
}

func (board *Board) MovesForPiece(piece Piece) []Move {
	space := piece.Space

	return board.movesForSpace(space, piece.Color)
}

func (board *Board) movesForSpace(startingSpace Space, color string) []Move {
	moves := []Move{}
	nextRank := 0
	if color == "white" {
		nextRank = startingSpace.Rank + 1
	} else {
		nextRank = startingSpace.Rank - 1
	}

	if notOnLeftEdge(startingSpace) {
		leftFile := decFile(startingSpace.File)
		targetSpace := Space{File: leftFile, Rank: nextRank}
		if nextMove, ok := getNextMove(startingSpace, targetSpace, board); ok {
			moves = append(moves, nextMove)
		}
	}

	if notOnRightEdge(startingSpace) {
		rightFile := incFile(startingSpace.File)
		targetSpace := Space{File: rightFile, Rank: nextRank}

		if nextMove, ok := getNextMove(startingSpace, targetSpace, board); ok {
			moves = append(moves, nextMove)
		}
	}

	return moves
}

func getNextMove(startingSpace Space, targetSpace Space, board *Board) (Move, bool) {
	if board.GetPieceAt(targetSpace).Color == "" {
		move := Move{StartingSpace: startingSpace, TargetSpace: targetSpace}
		return move, true
	} else {
		nextSpace := getNextSpace(startingSpace, targetSpace)

		if board.GetPieceAt(nextSpace).Color == "" {
			move := Move{StartingSpace: startingSpace, TargetSpace: nextSpace}
			return move, true
		} else {
			return Move{}, false
		}
	}
}
