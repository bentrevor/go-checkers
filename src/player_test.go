package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

func TestPlayer_GetsAMove(t *testing.T) {
	move := MoveFromString("a1 - b2")

	assert(t, SameSpace(A1, move.StartingSpace), "MoveFromString starting space")
	assert(t, SameSpace(B2, move.TargetSpace), "MoveFromString target space")
}
