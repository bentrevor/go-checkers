package checkers

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

type ConsoleInput struct{}
type ConsoleOutput struct{}

func (c ConsoleInput) GetMove(board Board, rules Rules, color Color) Move {
	input, err := c.getInput()

	if err != nil {
		log.Fatal("there was an error getting the input")
	}

	move, errorMessage := MoveFromString(input[0 : len(input)-1])

	if len(errorMessage) == 0 {
		return move
	} else {
		fmt.Println(errorMessage)
		return c.GetMove(board, rules, color)
	}
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
		piece, foundPiece := board.GetPieceAtSpace(space)

		if !foundPiece {
			row.WriteString(emptySpace)
		} else {
			abbrev := getPieceAbbrev(piece)
			printableSpace := fmt.Sprintf("|%s|", abbrev)
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

func (ConsoleInput) getInput() (string, error) {
	in := bufio.NewReader(os.Stdin)
	return in.ReadString('\n')
}
