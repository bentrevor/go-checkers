package checkers

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type Player interface {
	GetMove(Board) Move
	Color() string
}

type ConsoleIO struct{}

type HumanPlayer struct {
	color string
	Io    IO
}

func NewPlayer(color string) Player {
	return &HumanPlayer{color: color}
}

func (ConsoleIO) GetInput() (string, error) {
	in := bufio.NewReader(os.Stdin)
	return in.ReadString('\n')
}

func (ConsoleIO) PrintBoard(board Board) {
	fmt.Println()
	spaces := []Space{}
	emptySpace := "|_|"
	var row bytes.Buffer
	rows := []string{}

	for _, piece := range board.Pieces {
		spaces = append(spaces, piece.Space)
	}

	for i := 0; i < 64; i++ {
		space := SpaceForIndex(i)
		piece := board.GetPieceAtSpace(space)

		if piece.Color == "" {
			row.WriteString(emptySpace)
		} else {
			printableSpace := fmt.Sprintf("|%c|", piece.Color[0])
			row.WriteString(printableSpace)
		}
		if i%8 == 7 {
			row.WriteString(fmt.Sprintf("  %d\n", (i/8)+1))
			rows = append(rows, row.String())
			row.Reset()
		}
	}

	fmt.Println(reverseRows(rows))
	fmt.Println("  a  b  c  d  e  f  g  h")
}

func (hp *HumanPlayer) GetMove(board Board) Move {
	input, err := hp.Io.GetInput()

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
