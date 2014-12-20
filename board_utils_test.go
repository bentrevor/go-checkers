package checkers_test

import (
	. "github.com/bentrevor/checkers"
	"testing"
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
	assertEquals(t, Space{File: "a", Rank: 1}, SpaceForIndex(0))
	assertEquals(t, Space{File: "b", Rank: 1}, SpaceForIndex(1))
	assertEquals(t, Space{File: "h", Rank: 1}, SpaceForIndex(7))
	assertEquals(t, Space{File: "a", Rank: 2}, SpaceForIndex(8))
	assertEquals(t, Space{File: "b", Rank: 2}, SpaceForIndex(9))
	assertEquals(t, Space{File: "h", Rank: 2}, SpaceForIndex(15))
	assertEquals(t, Space{File: "a", Rank: 3}, SpaceForIndex(16))
	assertEquals(t, Space{File: "h", Rank: 8}, SpaceForIndex(63))
}
