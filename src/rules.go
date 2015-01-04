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
	return board.MovesForSpace(piece.Space)
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

func (board Board) IsGameOver() bool {
	blackPieceCount := board.countPiecesByColor(Black)
	whitePieceCount := board.countPiecesByColor(White)

	return whitePieceCount == 0 || blackPieceCount == 0
}

func (board Board) MovesForSpace(startingSpace Space) []Move {
	var moves []Move
	directions := []Direction{Northwest, Northeast, Southeast, Southwest}

	for _, direction := range directions {
		if move, moveCreated := board.TryMove(startingSpace, direction); moveCreated {
			moves = append(moves, move)
		}
	}

	return moves
}

func (board Board) TryMove(space Space, direction Direction) (Move, bool) {
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

func (board Board) moveInDirection(direction Direction, space Space) (Move, bool) {
	nearSpace, _ := GetNearSpaceInDirection(space, direction)
	farSpace, _ := GetFarSpaceInDirection(space, direction)

	if board.canMoveToNearSpace(direction, space) {
		return Move{StartingSpace: space, TargetSpace: nearSpace}, true
	} else if board.canCaptureToFarSpace(direction, space) {
		return Move{StartingSpace: space, TargetSpace: farSpace}, true
	} else {
		return Move{}, false
	}
}

func (board Board) canMoveToNearSpace(direction Direction, space Space) bool {
	nearSpace, _ := GetNearSpaceInDirection(space, direction)
	_, foundNearPiece := board.GetPieceAtSpace(nearSpace)

	return onBoard(nearSpace) && !foundNearPiece
}

func (board Board) canCaptureToFarSpace(direction Direction, space Space) bool {
	farSpace, _ := GetFarSpaceInDirection(space, direction)

	return onBoard(farSpace) &&
		board.oppositeColorOnNearSpace(direction, space) &&
		board.isEmpty(farSpace)
}

func (board Board) isEmpty(space Space) bool {
	_, foundPiece := board.GetPieceAtSpace(space)

	return !foundPiece
}

func (board Board) oppositeColorOnNearSpace(direction Direction, space Space) bool {
	nearSpace, _ := GetNearSpaceInDirection(space, direction)
	farSpace, _ := GetFarSpaceInDirection(space, direction)

	nearPiece, _ := board.GetPieceAtSpace(nearSpace)
	movingPiece, _ := board.GetPieceAtSpace(space)
	_, foundFarPiece := board.GetPieceAtSpace(farSpace)

	if nearPiece.Color != movingPiece.Color {
		return onBoard(farSpace) && !foundFarPiece
	}

	return false
}

func (board Board) countPiecesByColor(color Color) int {
	count := 0

	for _, piece := range board.Pieces {
		if piece.Color == color {
			count += 1
		}
	}

	return count
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
