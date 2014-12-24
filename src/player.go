package checkers

import (
	"bufio"
	"fmt"
	"os"
)

type Player interface {
	GetMove(Board) Move
	Color() string
}

type HumanPlayer struct {
	color string
}

func NewPlayer(color string) Player {
	return &HumanPlayer{color: color}
}

func (hp *HumanPlayer) GetMove(board Board) Move {
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("no")
		return hp.GetMove(board)
	}

	return MoveFromString(input)
}

func MoveFromString(input string) Move {
	return Move{StartingSpace: NewSpace(input[0:2]), TargetSpace: NewSpace(input[5:7])}
}

func (player *HumanPlayer) Color() string {
	return player.color
}
