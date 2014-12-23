package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers"
)

var game = NewGame()

func TestGame_WhiteGoesFirst(t *testing.T) {
	assertEquals(t, "white", game.CurrentColor)
}

func TestGame_TogglesPlayers(t *testing.T) {
	game.NextTurn()
	assertEquals(t, "black", game.CurrentColor)
}
