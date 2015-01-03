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

	leftNearSpace, _ := GetNearSpaceInDirection(space, Northwest)
	leftFarSpace, _ := GetFarSpaceInDirection(space, Northwest)

	rightNearSpace, _ := GetNearSpaceInDirection(space, Northeast)
	rightFarSpace, _ := GetFarSpaceInDirection(space, Northeast)

	assert(t, IsSameSpace(C5, leftNearSpace), "white left non capture space")
	assert(t, IsSameSpace(B6, leftFarSpace), "white left capture space")

	assert(t, IsSameSpace(E5, rightNearSpace), "white right non capture space")
	assert(t, IsSameSpace(F6, rightFarSpace), "white right capture space")
}

func TestBoardUtils_GetsSpacesInADirectionAtTheEdgeOfTheBoard(t *testing.T) {
	board := NewEmptyBoard()
	space := H2
	piece := Piece{Color: Black, Space: space}

	board.PlacePiece(piece)

	leftNearSpace, _ := GetNearSpaceInDirection(space, Southwest)
	_, leftFarSpaceCreated := GetFarSpaceInDirection(space, Southwest)

	_, rightNearSpaceCreated := GetNearSpaceInDirection(space, Southeast)
	_, rightFarSpaceCreated := GetFarSpaceInDirection(space, Southeast)

	assert(t, IsSameSpace(G1, leftNearSpace), "left non capture space")
	assert(t, !leftFarSpaceCreated, "left capture space shouldn't exist")

	assert(t, !rightNearSpaceCreated, "right non capture space shouldn't exist")
	assert(t, !rightFarSpaceCreated, "right capture space shouldn't exist")
}
