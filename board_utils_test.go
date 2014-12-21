package checkers_test

import (
	. "github.com/bentrevor/checkers"
	"testing"
	"fmt"
)

func useFmt() {
	fmt.Print()
}

func TestBoardUtils_KnowsTheColorOfASpace(t *testing.T) {
	assertEquals(t, SpaceColorForIndex(0), "black")
	assertEquals(t, SpaceColorForIndex(1), "white")
	assertEquals(t, SpaceColorForIndex(6), "black")
	assertEquals(t, SpaceColorForIndex(7), "white")

	assertEquals(t, SpaceColorForIndex(8), "white")
	assertEquals(t, SpaceColorForIndex(9), "black")
	assertEquals(t, SpaceColorForIndex(14), "white")
	assertEquals(t, SpaceColorForIndex(15), "black")

	assertEquals(t, SpaceColorForIndex(16), "black")
	assertEquals(t, SpaceColorForIndex(17), "white")
}

func TestBoardUtils_KnowsTheSpaceForAnIndex(t *testing.T) {
	assertEquals(t, NewSpace("a1"), SpaceForIndex(0))
	assertEquals(t, NewSpace("b1"), SpaceForIndex(1))
	assertEquals(t, NewSpace("h1"), SpaceForIndex(7))
	assertEquals(t, NewSpace("a2"), SpaceForIndex(8))
	assertEquals(t, NewSpace("b2"), SpaceForIndex(9))
	assertEquals(t, NewSpace("h2"), SpaceForIndex(15))
	assertEquals(t, NewSpace("a3"), SpaceForIndex(16))
	assertEquals(t, NewSpace("h8"), SpaceForIndex(63))
}

func TestBoardUtils_GetsSpacesInADirection(t *testing.T) {
	space := NewSpace("d3")

	leftNonCaptureSpace, _ := GetNonCaptureSpaceInDirection(space, "white", "left")
	leftCaptureSpace, _    := GetCaptureSpaceInDirection(space, "white", "left")

	rightNonCaptureSpace, _ := GetNonCaptureSpaceInDirection(space, "white", "right")
	rightCaptureSpace, _    := GetCaptureSpaceInDirection(space, "white", "right")

	// fmt.Println("\n\nasdf\n", leftNonCaptureSpace)
	assert(t, SameSpace(NewSpace("c4"), leftNonCaptureSpace), "white left non capture space")
	assert(t, SameSpace(NewSpace("b5"), leftCaptureSpace), "white left capture space")

	assert(t, SameSpace(NewSpace("e4"), rightNonCaptureSpace), "white right non capture space")
	assert(t, SameSpace(NewSpace("f5"), rightCaptureSpace), "white right capture space")
}

func TestBoardUtils_GetsSpacesInADirectionAtTheEdgeOfTheBoard(t *testing.T) {
	board := NewEmptyBoard()
	space := NewSpace("h2")
	piece := Piece{Color: "black", Space: space}

	board.PlacePiece(piece)

	leftNonCaptureSpace, _     := GetNonCaptureSpaceInDirection(space, "black", "left")
	_, leftCaptureSpaceCreated := GetCaptureSpaceInDirection(space, "black", "left")

	_, rightNonCaptureSpaceCreated := GetNonCaptureSpaceInDirection(space, "black", "right")
	_, rightCaptureSpaceCreated    := GetCaptureSpaceInDirection(space, "black", "right")

	assert(t, SameSpace(NewSpace("g1"), leftNonCaptureSpace), "left non capture space")
	assert(t, !leftCaptureSpaceCreated, "left capture space shouldn't exist")

	assert(t, !rightNonCaptureSpaceCreated, "right non capture space shouldn't exist")
	assert(t, !rightCaptureSpaceCreated, "right capture space shouldn't exist")
}
