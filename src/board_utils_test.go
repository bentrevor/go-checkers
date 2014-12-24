package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers"
)

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
	assertEquals(t, A1, SpaceForIndex(0))
	assertEquals(t, B2, SpaceForIndex(9))
	assertEquals(t, H2, SpaceForIndex(15))
	assertEquals(t, A3, SpaceForIndex(16))
	assertEquals(t, H8, SpaceForIndex(63))
}

func TestBoardUtils_GetsSpacesInADirection(t *testing.T) {
	space := D4

	leftNonCaptureSpace, _ := GetNonCaptureSpaceInDirection(space, "white", "left")
	leftCaptureSpace, _ := GetCaptureSpaceInDirection(space, "white", "left")

	rightNonCaptureSpace, _ := GetNonCaptureSpaceInDirection(space, "white", "right")
	rightCaptureSpace, _ := GetCaptureSpaceInDirection(space, "white", "right")

	assert(t, SameSpace(C5, leftNonCaptureSpace), "white left non capture space")
	assert(t, SameSpace(B6, leftCaptureSpace), "white left capture space")

	assert(t, SameSpace(E5, rightNonCaptureSpace), "white right non capture space")
	assert(t, SameSpace(F6, rightCaptureSpace), "white right capture space")
}

func TestBoardUtils_GetsSpacesInADirectionAtTheEdgeOfTheBoard(t *testing.T) {
	board := NewEmptyBoard()
	space := H2
	piece := Piece{Color: "black", Space: space}

	board.PlacePiece(piece)

	leftNonCaptureSpace, _ := GetNonCaptureSpaceInDirection(space, "black", "left")
	_, leftCaptureSpaceCreated := GetCaptureSpaceInDirection(space, "black", "left")

	_, rightNonCaptureSpaceCreated := GetNonCaptureSpaceInDirection(space, "black", "right")
	_, rightCaptureSpaceCreated := GetCaptureSpaceInDirection(space, "black", "right")

	assert(t, SameSpace(G1, leftNonCaptureSpace), "left non capture space")
	assert(t, !leftCaptureSpaceCreated, "left capture space shouldn't exist")

	assert(t, !rightNonCaptureSpaceCreated, "right non capture space shouldn't exist")
	assert(t, !rightCaptureSpaceCreated, "right capture space shouldn't exist")
}
