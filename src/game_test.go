package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

const fakeInput = "c3 - d4"

var player1 = NewHumanPlayer(White)
var player2 = NewHumanPlayer(Black)

var game = NewGame(player1, player2, MockOutput{})

type MockPlayer struct {
	color Color
}

type MockInput struct{}
type MockOutput struct{}

func (MockInput) GetInput() (string, error) {
	return fakeInput, nil
}

func (MockOutput) PrintBoard(b Board) {}

func (MockPlayer) GetMove(board Board) Move {
	move, _ := MoveFromString(fakeInput)
	return move
}

func (MockPlayer) Color() Color {
	return White
}

func NewMockPlayer() Player {
	return &MockPlayer{color: White}
}

func TestGame_WhiteGoesFirst(t *testing.T) {
	assertEquals(t, White, game.CurrentColor())
}

func TestGame_TogglesPlayers(t *testing.T) {
	assertEquals(t, White, game.CurrentColor())

	game.CurrentPlayer = NewMockPlayer()
	game.Board = NewGameBoard()
	game.NextTurn()

	assertEquals(t, Black, game.CurrentColor())
}

func TestGame_MakesMovesOnTheBoard(t *testing.T) {
	game.CurrentPlayer = NewMockPlayer()

	d4Piece, _ := game.Board.GetPieceAtSpace(D4)
	c3Piece, _ := game.Board.GetPieceAtSpace(C3)

	assertEquals(t, White, d4Piece.Color)
	assertEquals(t, NoColor, c3Piece.Color)
}

func TestGame_ValidatesMoves(t *testing.T) {
	assert(t, !game.IsValidMove(Move{B4, C5}), "invalid move - no piece at starting space")
}

func TestGame_InitFromFen(t *testing.T) {
	game = NewGame(player1, player2, MockOutput{})
	game.InitFromFen("4/4/4/4/bw2/4/4/4 b")

	numPieces := 0

	for _, piece := range game.Board.Pieces {
		if piece.Color != NoColor {
			numPieces += 1
		}
	}

	assertEquals(t, Black, game.CurrentColor())
	assertEquals(t, 2, numPieces)
}
