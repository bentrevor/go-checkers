package checkers_test

import (
	"testing"

	. "github.com/bentrevor/checkers/src"
)

var fakeInput = "c3 - d4"

var player1 = NewHumanPlayer(White)
var player2 = NewHumanPlayer(Black)

var game = NewGame(player1, player2, MockOutput{})

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
	assertEquals(t, White, game.CurrentColor())
}

func TestGame_TogglesPlayers(t *testing.T) {
	assertEquals(t, White, game.CurrentColor())

	game.CurrentPlayer = NewMockPlayer("c3 - d4")
	game.Board = NewGameBoard()
	game.NextTurn()

	assertEquals(t, Black, game.CurrentColor())
}

func TestGame_MakesMovesOnTheBoard(t *testing.T) {
	game.CurrentPlayer = NewMockPlayer("c3 - d4")

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

func TestGame_PromotesPiecesToKing(t *testing.T) {
	game = NewGame(player1, player2, MockOutput{})
	game.InitFromFen("4/4/4/4/b/4/3w/4 w")

	game.CurrentPlayer = NewMockPlayer("g7 - f8")
	game.NextTurn()

	king, _ := game.Board.GetPieceAtSpace(F8)
	assert(t, king.IsKing, "king is king yo")
}
