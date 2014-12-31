package checkers

type Move struct {
	StartingSpace Space
	TargetSpace   Space
}

func MovesForPiece(piece Piece, board Board) []Move {
	space := piece.Space

	moves := board.MovesForSpace(space, piece.Color)
	return moves
}

func gameOver(board Board) bool {
	return false
}

func (board Board) MovesForSpace(startingSpace Space, color Color) []Move {
	moves := []Move{}

	nextRank := 0

	if color == White {
		nextRank = startingSpace.Rank + 1
	} else {
		nextRank = startingSpace.Rank - 1
	}

	if leftMove, ok := tryLeftMove(board, startingSpace, nextRank); ok {
		moves = append(moves, leftMove)
	}

	if rightMove, ok := tryRightMove(board, startingSpace, nextRank); ok {
		moves = append(moves, rightMove)
	}

	return moves
}

func IsLegalMove(move Move, board Board, color Color) bool {
	if color != board.GetPieceAtSpace(move.StartingSpace).Color {
		return false
	} else {
		moves := board.MovesForSpace(move.StartingSpace, color)

		if IncludesMove(moves, move) {
			return true
		} else {
			return false
		}
	}
}

func tryLeftMove(board Board, startingSpace Space, nextRank int) (Move, bool) {
	if notOnLeftEdge(startingSpace) {
		leftFile := decFile(startingSpace.File)
		targetSpace := Space{File: leftFile, Rank: nextRank}

		if nextMove, ok := getNextMove(startingSpace, targetSpace, board); ok {
			return nextMove, true
		}
	}

	return Move{}, false
}

func tryRightMove(board Board, startingSpace Space, nextRank int) (Move, bool) {
	if notOnRightEdge(startingSpace) {
		rightFile := incFile(startingSpace.File)
		targetSpace := Space{File: rightFile, Rank: nextRank}

		if nextMove, ok := getNextMove(startingSpace, targetSpace, board); ok {
			return nextMove, true
		}
	}

	return Move{}, false
}

func getNextMove(startingSpace Space, targetSpace Space, board Board) (Move, bool) {
	if board.GetPieceAtSpace(targetSpace).Color == NoColor {
		move := Move{StartingSpace: startingSpace, TargetSpace: targetSpace}
		return move, true
	} else {
		nextSpace := getNextSpace(startingSpace, targetSpace)

		if board.GetPieceAtSpace(nextSpace).Color == NoColor {
			move := Move{StartingSpace: startingSpace, TargetSpace: nextSpace}
			return move, true
		} else {
			return Move{}, false
		}
	}
}
