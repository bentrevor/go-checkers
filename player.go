package checkers

import (
	"bufio"
	"fmt"
	"os"
)

type Player interface {
	GetMove(*Board) Move
	Color() string
}

type HumanPlayer struct {
	color string
}

func NewPlayer(color string) Player {
	return &HumanPlayer{color: color}
}

func (*HumanPlayer) GetMove(board *Board) Move {
	in := bufio.NewReader(os.Stdin)
	_, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("\n\n\n              error!\n\n", err, "\n\n--------------------\n")
	}

	return Move{}
}

func (player *HumanPlayer) Color() string {
	return player.color
}
