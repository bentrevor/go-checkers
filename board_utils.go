package checkers

// import "fmt"

func IncludesMove(moves []Move, move Move) bool {
	for _, any_move := range moves {
		if sameMove(move, any_move) {
			return true
		}
	}

	return false
}

func sameMove(move1 Move, move2 Move) bool {
	return SameSpace(move1.StartingSpace, move2.StartingSpace) && SameSpace(move1.TargetSpace, move2.TargetSpace)
}

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

func SameSpace(pieceSpace Space, targetSpace Space) bool {
	return pieceSpace.Rank == targetSpace.Rank &&
		pieceSpace.File == targetSpace.File
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

func SpaceForIndex(index int) Space {
	rank := (index / 8) + 1
	file := string((index % 8) + 97)

	space := Space{File: file, Rank: rank}
	return space
}

func SpaceColorForIndex(index int) string {
	evenColor := ""
	oddColor := ""
	if ((index / 8) % 2) == 0 {
		evenColor = "black"
		oddColor  = "white"
	} else {
		evenColor = "white"
		oddColor  = "black"
	}

	if index % 2 == 0 {
		return evenColor
	} else {
		return oddColor
	}
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
		// fmt.Println("")//"\n\nspace: ", space, "\nnext space: ", nonCaptureSpace, "\n==\n")
		return GetNonCaptureSpaceInDirection(nonCaptureSpace, color, direction)
	}

	return Space{}, false
}

func onBoard(file string, rank int) bool {
	r := rank > 0 && rank < 8 &&
		file >= "a" && file <= "h"
	// fmt.Println("got", r, "for", file, rank)
	return r
}
