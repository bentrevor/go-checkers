package checkers_test

import (
	. "github.com/bentrevor/checkers"
	// "fmt"
	"testing"
)

var game = NewGame()

func TestGame_WhiteGoesFirst(t *testing.T) {
	assertEquals(t, "white", game.CurrentColor)
}

func TestGame_TogglesPlayers(t *testing.T) {
	game.NextTurn()
	assertEquals(t, "black", game.CurrentColor)
}
