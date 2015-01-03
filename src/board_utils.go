package checkers

import "math"

func IncludesMove(moves []Move, move Move) bool {
	for _, any_move := range moves {
		if IsSameMove(move, any_move) {
			return true
		}
	}

	return false
}

func IsSameMove(move1 Move, move2 Move) bool {
	return IsSameSpace(move1.StartingSpace, move2.StartingSpace) &&
		IsSameSpace(move1.TargetSpace, move2.TargetSpace)
}

func IsSameSpace(pieceSpace Space, targetSpace Space) bool {
	return pieceSpace.Rank == targetSpace.Rank &&
		pieceSpace.File == targetSpace.File
}

func SpaceForIndex(index int) Space {
	rank := (index / 8) + 1
	file := string((index % 8) + 97)

	space := Space{File: file, Rank: rank}
	return space
}

func SpaceColorForIndex(index int) Color {
	oddColor, evenColor := getOddAndEvenColor(index)

	if index%2 == 0 {
		return evenColor
	} else {
		return oddColor
	}
}

func leftTargetSpace(board Board, space Space) Space {
	piece, foundPiece := board.GetPieceAtSpace(space)

	if !foundPiece {
		// TODO
	}

	nextRank := 0
	nextFile := decFile(space.File)

	if piece.Color == Black {
		nextRank = space.Rank - 1
	} else {
		nextRank = space.Rank + 1
	}

	return Space{File: nextFile, Rank: nextRank}
}

func GetNonCaptureSpaceInDirection(space Space, color Color, direction string) (Space, bool) {
	nextRank := 0
	nextFile := ""

	if color == Black {
		nextRank = space.Rank - 1
	} else {
		nextRank = space.Rank + 1
	}

	if direction == "left" {
		nextFile = decFile(space.File)
	} else {
		nextFile = incFile(space.File)
	}

	nonCaptureSpace := Space{File: nextFile, Rank: nextRank}
	if onBoard(nonCaptureSpace) {
		return nonCaptureSpace, true
	} else {
		return Space{}, false
	}
}

func GetCaptureSpaceInDirection(space Space, color Color, direction string) (Space, bool) {
	if nonCaptureSpace, ok := GetNonCaptureSpaceInDirection(space, color, direction); ok {
		return GetNonCaptureSpaceInDirection(nonCaptureSpace, color, direction)
	}

	return Space{}, false
}

func getNextSpaceInSameDirection(startingSpace Space, targetSpace Space) Space {
	increasingRank := startingSpace.Rank < targetSpace.Rank
	increasingFile := startingSpace.File[0] < targetSpace.File[0]
	nextRank := 0
	nextFile := ""

	if increasingRank {
		nextRank = targetSpace.Rank + 1
	} else {
		nextRank = targetSpace.Rank - 1
	}

	if increasingFile {
		nextFile = string(targetSpace.File[0] + 1)
	} else {
		nextFile = string(targetSpace.File[0] - 1)
	}

	return Space{File: nextFile, Rank: nextRank}
}

func onLeftEdge(space Space) bool {
	return space.File == "a"
}

func onRightEdge(space Space) bool {
	return space.File == "h"
}

func incFile(file string) string {
	return string(file[0] + 1)
}

func decFile(file string) string {
	return string(file[0] - 1)
}

func initialPieceColorFor(index int) Color {
	if index < 24 {
		return White
	} else if index > 39 {
		return Black
	} else {
		return NoColor
	}
}

func onBoard(space Space) bool {
	rank := space.Rank
	file := space.File

	r := rank > 0 && rank < 8 &&
		file >= "a" && file <= "h"

	return r
}

func getOddAndEvenColor(index int) (Color, Color) {
	row := index / 8

	if row%2 == 0 {
		return White, Black
	} else {
		return Black, White
	}
}

func isCaptureMove(move Move) bool {
	rankDelta := math.Abs(float64(move.StartingSpace.Rank - move.TargetSpace.Rank))

	// going to have to fix this for double jumps
	return rankDelta > 1
}

func capturedSpace(move Move) Space {
	capturedRank := rankBetween(move.StartingSpace.Rank, move.TargetSpace.Rank)
	capturedFile := fileBetween(move.StartingSpace.File, move.TargetSpace.File)

	return Space{Rank: capturedRank, File: capturedFile}
}

func rankBetween(rank1 int, rank2 int) int {
	return (rank1 + rank2) / 2
}

func fileBetween(file1 string, file2 string) string {
	captureFile := (file1[0] + file2[0]) / 2

	return string(captureFile)
}

func spacesForRank(rank int) []Space {
	spaces := [][]Space{
		[]Space{A1, C1, E1, G1},
		[]Space{B2, D2, F2, H2},
		[]Space{A3, C3, E3, G3},
		[]Space{B4, D4, F4, H4},
		[]Space{A5, C5, E5, G5},
		[]Space{B6, D6, F6, H6},
		[]Space{A7, C7, E7, G7},
		[]Space{B8, D8, F8, H8},
	}

	return spaces[rank-1]
}
