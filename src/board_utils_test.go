package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

func TestBoardUtils_KnowsTheColorOfASpace(t *testing.T) {
	assertEquals(t, SpaceColorForIndex(0), Black)
	assertEquals(t, SpaceColorForIndex(1), White)
	assertEquals(t, SpaceColorForIndex(6), Black)
	assertEquals(t, SpaceColorForIndex(7), White)

	assertEquals(t, SpaceColorForIndex(8), White)
	assertEquals(t, SpaceColorForIndex(9), Black)
	assertEquals(t, SpaceColorForIndex(14), White)
	assertEquals(t, SpaceColorForIndex(15), Black)

	assertEquals(t, SpaceColorForIndex(16), Black)
	assertEquals(t, SpaceColorForIndex(17), White)
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

	leftNonCaptureSpace, _ := GetNonCaptureSpaceInDirection(space, White, "left")
	leftCaptureSpace, _ := GetCaptureSpaceInDirection(space, White, "left")

	rightNonCaptureSpace, _ := GetNonCaptureSpaceInDirection(space, White, "right")
	rightCaptureSpace, _ := GetCaptureSpaceInDirection(space, White, "right")

	assert(t, SameSpace(C5, leftNonCaptureSpace), "white left non capture space")
	assert(t, SameSpace(B6, leftCaptureSpace), "white left capture space")

	assert(t, SameSpace(E5, rightNonCaptureSpace), "white right non capture space")
	assert(t, SameSpace(F6, rightCaptureSpace), "white right capture space")
}

func TestBoardUtils_GetsSpacesInADirectionAtTheEdgeOfTheBoard(t *testing.T) {
	board := NewEmptyBoard()
	space := H2
	piece := Piece{Color: Black, Space: space}

	board.PlacePiece(piece)

	leftNonCaptureSpace, _ := GetNonCaptureSpaceInDirection(space, Black, "left")
	_, leftCaptureSpaceCreated := GetCaptureSpaceInDirection(space, Black, "left")

	_, rightNonCaptureSpaceCreated := GetNonCaptureSpaceInDirection(space, Black, "right")
	_, rightCaptureSpaceCreated := GetCaptureSpaceInDirection(space, Black, "right")

	assert(t, SameSpace(G1, leftNonCaptureSpace), "left non capture space")
	assert(t, !leftCaptureSpaceCreated, "left capture space shouldn't exist")

	assert(t, !rightNonCaptureSpaceCreated, "right non capture space shouldn't exist")
	assert(t, !rightCaptureSpaceCreated, "right capture space shouldn't exist")
}
