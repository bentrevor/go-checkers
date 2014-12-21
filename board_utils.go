package checkers

// import "fmt"

func IncludesMove(moves []Move, move Move) bool {
	for _, any_move := range moves {
		if SameMove(move, any_move) {
			return true
		}
	}

	return false
}

func SameSpace(pieceSpace Space, targetSpace Space) bool {
	return pieceSpace.Rank == targetSpace.Rank &&
		pieceSpace.File == targetSpace.File
}

func SameMove(move1 Move, move2 Move) bool {
	return SameSpace(move1.StartingSpace, move2.StartingSpace) && SameSpace(move1.TargetSpace, move2.TargetSpace)
}

func SpaceForIndex(index int) Space {
	rank := (index / 8) + 1
	file := string((index % 8) + 97)

	space := Space{File: file, Rank: rank}
	return space
}

func SpaceColorForIndex(index int) string {
	oddColor, evenColor := getOddAndEvenColor(index)

	if index % 2 == 0 {
		return evenColor
	} else {
		return oddColor
	}
}

func leftTargetSpace(board *Board, space Space) Space {
	color := board.GetPieceAtSpace(space).Color
	nextRank := 0
	nextFile := decFile(space.File)

	if color == "black" {
		nextRank = space.Rank - 1
	} else {
		nextRank = space.Rank + 1
	}

	return Space{File: nextFile, Rank: nextRank}
}

func GetNonCaptureSpaceInDirection(space Space, color string, direction string) (Space, bool) {
	nextRank := 0
	nextFile := ""

	if color == "black" {
		nextRank = space.Rank - 1
	} else {
		nextRank = space.Rank + 1
	}

	if direction == "left" {
		nextFile = decFile(space.File)
	} else {
		nextFile = incFile(space.File)
	}

	if onBoard(nextFile, nextRank) {
		return Space{File: nextFile, Rank: nextRank}, true
	} else {
		return Space{}, false
	}
}

func GetCaptureSpaceInDirection(space Space, color string, direction string) (Space, bool) {
	if nonCaptureSpace, ok := GetNonCaptureSpaceInDirection(space, color, direction); ok {
		return GetNonCaptureSpaceInDirection(nonCaptureSpace, color, direction)
	}

	return Space{}, false
}

// private

func getNextSpace(startingSpace Space, targetSpace Space) Space {
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

func notOnLeftEdge(space Space) bool {
	return space.File != "a"
}

func notOnRightEdge(space Space) bool {
	return space.File != "h"
}

func incFile(file string) string {
	return string(file[0] + 1)
}

func decFile(file string) string {
	return string(file[0] - 1)
}

func initialPieceColorFor(index int) string {
	if index < 24 {
		return "white"
	} else if index > 39 {
		return "black"
	} else {
		return ""
	}
}


func onBoard(file string, rank int) bool {
	r := rank > 0 && rank < 8 &&
		file >= "a" && file <= "h"
	return r
}

func getOddAndEvenColor(index int) (string, string) {
	row := index / 8

	if row % 2 == 0 {
		return "white", "black"
	} else {
		return "black", "white"
	}
}
