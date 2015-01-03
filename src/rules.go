package checkers

type Move struct {
	StartingSpace Space
	TargetSpace   Space
}

type Direction struct {
	increasingFile bool
	increasingRank bool
}

var (
	Northwest Direction = Direction{increasingFile: false, increasingRank: true}
	Northeast Direction = Direction{increasingFile: true, increasingRank: true}
	Southwest Direction = Direction{increasingFile: false, increasingRank: false}
	Southeast Direction = Direction{increasingFile: true, increasingRank: false}
)

func (board Board) MovesForPiece(piece Piece) []Move {
	space := piece.Space

	moves := board.MovesForSpace(space)
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
	blackPieceCount := countPiecesByColor(Black, board.Pieces)
	whitePieceCount := countPiecesByColor(White, board.Pieces)

	return whitePieceCount == 0 || blackPieceCount == 0
}

func countPiecesByColor(color Color, pieces []Piece) int {
	count := 0

	for _, piece := range pieces {
		if piece.Color == color {
			count += 1
		}
	}

	return count
}

func (board Board) MovesForSpace(startingSpace Space) []Move {
	var moves []Move

	nwMove, nwMoveCreated := board.TryMove(moves, startingSpace, Northwest)
	neMove, neMoveCreated := board.TryMove(moves, startingSpace, Northeast)
	swMove, swMoveCreated := board.TryMove(moves, startingSpace, Southwest)
	seMove, seMoveCreated := board.TryMove(moves, startingSpace, Southeast)

	if nwMoveCreated {
		moves = append(moves, nwMove)
	}

	if neMoveCreated {
		moves = append(moves, neMove)
	}

	if swMoveCreated {
		moves = append(moves, swMove)
	}

	if seMoveCreated {
		moves = append(moves, seMove)
	}

	return moves
}

func isWrongDirection(color Color, direction Direction) bool {
	var correctColor Color

	if direction.increasingRank {
		correctColor = White
	} else {
		correctColor = Black
	}

	return color != correctColor
}

func (board Board) moveInDirection(direction Direction, space Space) (Move, bool) {
	nonCaptureSpace, _ := GetNonCaptureSpaceInDirection(space, direction)
	captureSpace, _ := GetCaptureSpaceInDirection(space, direction)

	movingPiece, _ := board.GetPieceAtSpace(space)
	nonCapturePiece, foundNonCapturePiece := board.GetPieceAtSpace(nonCaptureSpace)
	_, foundCapturePiece := board.GetPieceAtSpace(captureSpace)

	if onBoard(nonCaptureSpace) && !foundNonCapturePiece {
		return Move{StartingSpace: space, TargetSpace: nonCaptureSpace}, true
	} else if nonCapturePiece.Color != movingPiece.Color && onBoard(captureSpace) && !foundCapturePiece {
		return Move{StartingSpace: space, TargetSpace: captureSpace}, true
	} else {
		return Move{}, false
	}
}

func (board Board) TryMove(moves []Move, space Space, direction Direction) (Move, bool) {
	piece, _ := board.GetPieceAtSpace(space)

	if piece.IsKing {
		return board.moveInDirection(direction, space)
	} else {
		if isWrongDirection(piece.Color, direction) {
			return Move{}, false
		} else {
			return board.moveInDirection(direction, space)
		}
	}
}

func movesFor(board Board, startingSpace Space, nextRank int) []Move {
	moves := []Move{}

	if leftMove, ok := tryLeftMove(board, startingSpace, nextRank); ok {
		moves = append(moves, leftMove)
	}

	if rightMove, ok := tryRightMove(board, startingSpace, nextRank); ok {

		moves = append(moves, rightMove)
	}

	return moves
}

func IsLegalMove(move Move, board Board, color Color) bool {
	piece, foundPiece := board.GetPieceAtSpace(move.StartingSpace)

	if !foundPiece || piece.Color != color {
		return false
	} else {
		moves := board.MovesForPiece(piece)

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
