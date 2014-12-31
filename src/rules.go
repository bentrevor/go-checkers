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

func MoveFromString(input string) (Move, string) {
	if len(input) != 7 {
		return Move{}, "wrong length: enter a move like 'c3 - d4'"
	} else {
		startingSpace := NewSpace(input[0:2])
		targetSpace := NewSpace(input[5:7])

		if onBoard(startingSpace) && onBoard(targetSpace) {
			return Move{StartingSpace: startingSpace, TargetSpace: targetSpace}, ""
		} else {
			return Move{}, "enter real moves, dummy"
		}
	}
}

func IsGameOver(board Board) bool {
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
	_, foundPiece := board.GetPieceAtSpace(move.StartingSpace)

	if !foundPiece {
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
	if !onLeftEdge(startingSpace) {
		leftFile := decFile(startingSpace.File)
		targetSpace := Space{File: leftFile, Rank: nextRank}

		if nextMove, ok := getNextMove(startingSpace, targetSpace, board); ok {
			return nextMove, true
		}
	}

	return Move{}, false
}

func tryRightMove(board Board, startingSpace Space, nextRank int) (Move, bool) {
	if !onRightEdge(startingSpace) {
		rightFile := incFile(startingSpace.File)
		targetSpace := Space{File: rightFile, Rank: nextRank}

		if nextMove, ok := getNextMove(startingSpace, targetSpace, board); ok {
			return nextMove, true
		}
	}

	return Move{}, false
}

func getNextMove(startingSpace Space, targetSpace Space, board Board) (Move, bool) {
	_, pieceAtTargetSpace := board.GetPieceAtSpace(targetSpace)

	if !pieceAtTargetSpace {
		move := Move{StartingSpace: startingSpace, TargetSpace: targetSpace}
		return move, true
	} else {
		nextSpace := getNextSpaceInSameDirection(startingSpace, targetSpace)
		_, pieceAtNextSpace := board.GetPieceAtSpace(nextSpace)

		if !pieceAtNextSpace {
			move := Move{StartingSpace: startingSpace, TargetSpace: nextSpace}
			return move, true
		} else {
			return Move{}, false
		}
	}
}
