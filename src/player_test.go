package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

func TestPlayer_GetsAMove(t *testing.T) {
	move, _ := MoveFromString("a1 - b2")

	assert(t, IsSameSpace(A1, move.StartingSpace), "MoveFromString starting space")
	assert(t, IsSameSpace(B2, move.TargetSpace), "MoveFromString target space")
}
