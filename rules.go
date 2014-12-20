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

	if leftMove, ok := tryLeftMove(board, startingSpace, nextRank); ok {
		moves = append(moves, leftMove)
	}

	if rightMove, ok := tryRightMove(board, startingSpace, nextRank); ok {
		moves = append(moves, rightMove)
	}

	return moves
}

func tryLeftMove(board *Board, startingSpace Space, nextRank int) (Move, bool) {
	if notOnLeftEdge(startingSpace) {
		leftFile := decFile(startingSpace.File)
		targetSpace := Space{File: leftFile, Rank: nextRank}

		if nextMove, ok := getNextMove(startingSpace, targetSpace, board); ok {
			return nextMove, true
		}
	}

	return Move{}, false
}

func tryRightMove(board *Board, startingSpace Space, nextRank int) (Move, bool) {
	if notOnRightEdge(startingSpace) {
		rightFile := incFile(startingSpace.File)
		targetSpace := Space{File: rightFile, Rank: nextRank}

		if nextMove, ok := getNextMove(startingSpace, targetSpace, board); ok {
			return nextMove, true
		}
	}

	return Move{}, false
}

func getNextMove(startingSpace Space, targetSpace Space, board *Board) (Move, bool) {
	if board.GetPieceAtSpace(targetSpace).Color == "" {
		move := Move{StartingSpace: startingSpace, TargetSpace: targetSpace}
		return move, true
	} else {
		nextSpace := getNextSpace(startingSpace, targetSpace)

		if board.GetPieceAtSpace(nextSpace).Color == "" {
			move := Move{StartingSpace: startingSpace, TargetSpace: nextSpace}
			return move, true
		} else {
			return Move{}, false
		}
	}
}
