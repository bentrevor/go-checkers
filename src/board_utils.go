package checkers

import (
	"math"
	"strings"
)

func BlackSpaces() [][]Space {
	return [][]Space{
		[]Space{A1, C1, E1, G1},
		[]Space{B2, D2, F2, H2},
		[]Space{A3, C3, E3, G3},
		[]Space{B4, D4, F4, H4},
		[]Space{A5, C5, E5, G5},
		[]Space{B6, D6, F6, H6},
		[]Space{A7, C7, E7, G7},
		[]Space{B8, D8, F8, H8},
	}
}

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

func GetNearSpaceInDirection(space Space, direction Direction) (Space, bool) {
	nextRank := nextRankInDirection(space, direction)
	nextFile := nextFileInDirection(space, direction)

	nearSpace := Space{File: nextFile, Rank: nextRank}
	if onBoard(nearSpace) {
		return nearSpace, true
	} else {
		return Space{}, false
	}
}

func GetFarSpaceInDirection(space Space, direction Direction) (Space, bool) {
	if nearSpace, ok := GetNearSpaceInDirection(space, direction); ok {
		return GetNearSpaceInDirection(nearSpace, direction)
	}

	return Space{}, false
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

	return rank > 0 && rank <= 8 &&
		file >= "a" && file <= "h"
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

	// TODO going to have to fix this for double jumps
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

func getPieceAbbrev(piece Piece) string {
	if piece.Color == NoColor {
		return " "
	} else {
		token := string(piece.Color[0])

		if piece.IsKing {
			return string(strings.ToUpper(string(token)))
		} else {
			return string(token)
		}
	}
}

func nextRankInDirection(space Space, direction Direction) int {
	if direction.increasingRank {
		return space.Rank + 1
	} else {
		return space.Rank - 1
	}
}

func nextFileInDirection(space Space, direction Direction) string {
	if direction.increasingFile {
		return incFile(space.File)
	} else {
		return decFile(space.File)
	}
}
