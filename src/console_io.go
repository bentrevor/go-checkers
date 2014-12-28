package checkers

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type ConsoleInput struct{}
type ConsoleOutput struct{}

func (c ConsoleInput) GetMove(board Board) Move {
	input, err := c.GetInput()

	if err != nil {
		// error!
	}

	return MoveFromString(input)
}

func (ConsoleInput) GetInput() (string, error) {
	in := bufio.NewReader(os.Stdin)
	return in.ReadString('\n')
}

func (ConsoleOutput) PrintBoard(board Board) {
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
