package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

var fakeInput = "c3 - d4"

var player1 = NewHumanPlayer(White)
var player2 = NewHumanPlayer(Black)

func NewGameWithMockPlayers() Game {
	return NewGame(player1, player2, MockOutput{})
}

type MockPlayer struct {
	color     Color
	FakeInput string
}

type MockInput struct{}
type MockOutput struct{}

func (MockInput) GetInput() (string, error) {
	return fakeInput, nil
}

func (MockOutput) PrintBoard(b Board) {}

func (p MockPlayer) GetMove(board Board) Move {
	move, _ := MoveFromString(p.FakeInput)
	return move
}

func (MockPlayer) Color() Color {
	return White
}

func NewMockPlayer(input string) Player {
	return &MockPlayer{color: White, FakeInput: input}
}

func TestGame_WhiteGoesFirst(t *testing.T) {
	game := NewGameWithMockPlayers()
	assertEquals(t, White, game.CurrentColor())
}

func TestGame_TogglesPlayers(t *testing.T) {
	game := NewGameWithMockPlayers()

	game.CurrentPlayer = NewMockPlayer("c3 - d4")
	game.Board = NewGameBoard()
	game.NextTurn()

	assertEquals(t, Black, game.CurrentColor())
}

func TestGame_MakesMovesOnTheBoard(t *testing.T) {
	game := NewGameWithMockPlayers()
	game.CurrentPlayer = NewMockPlayer("c3 - d4")

	d4Piece, _ := game.Board.GetPieceAtSpace(D4)
	c3Piece, _ := game.Board.GetPieceAtSpace(C3)

	assertEquals(t, NoColor, d4Piece.Color)
	assertEquals(t, White, c3Piece.Color)
}

func TestGame_ValidatesMoves(t *testing.T) {
	game := NewGameWithMockPlayers()
	assert(t, !game.IsValidMove(Move{B4, C5}), "invalid move - no piece at starting space")
}

func TestGame_PromotesPiecesToKing(t *testing.T) {
	game := NewGameWithMockPlayers()
	game.InitFromFen("4/4/4/4/b/4/3w/4 w")

	game.CurrentPlayer = NewMockPlayer("g7 - f8")
	game.NextTurn()

	king, _ := game.Board.GetPieceAtSpace(F8)
	assert(t, king.IsKing, "king is king yo")
}
