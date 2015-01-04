package checkers

import "time"

type Player interface {
	GetMove(Board, Rules) Move
	Color() Color
}

type MoveDecider interface {
	GetMove(Board, Rules, Color) Move
}

type Color string

const (
	NoColor Color = ""
	White   Color = "white"
	Black   Color = "black"
)

type HumanPlayer struct {
	color       Color
	MoveDecider ConsoleInput
}

type RandomPlayer struct {
	color       Color
	MoveDecider RandomMove
}

func NewHumanPlayer(color Color) Player {
	return &HumanPlayer{color: color, MoveDecider: ConsoleInput{}}
}

func NewRandomPlayer(color Color) Player {
	return &RandomPlayer{color: color, MoveDecider: RandomMove{}}
}

func (player *HumanPlayer) GetMove(board Board, rules Rules) Move {
	return player.MoveDecider.GetMove(board, rules, player.Color())
}

func (player *HumanPlayer) Color() Color {
	return player.color
}

// TODO this is dumb
func (player *RandomPlayer) GetMove(board Board, rules Rules) Move {
	time.Sleep(2 * time.Second)
	return player.MoveDecider.GetMove(board, rules, player.Color())
}

func (player *RandomPlayer) Color() Color {
	return player.color
}
