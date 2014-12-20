package checkers

import "fmt"

type Space struct {
	Rank int
	File string
}

type Piece struct {
	Color string
	Space Space
}

type Board struct {
	Pieces []Piece
}

type Move struct {
	StartingSpace Space
	TargetSpace Space
}

func dummy() {
	fmt.Println()
}

func NewGameBoard() *Board {
	board := &Board{make([]Piece, 24, 80)}
	pieceIndex := 0

	for i := 0; i < 64; i++ {
		_, created := board.createPieceAtIndex(pieceIndex, i)

		if created && pieceIndex < cap(board.Pieces) {
			pieceIndex += 1
		}
	}

	return board
}

func (board *Board) getNextMove(startingSpace Space, targetSpace Space) (Move, bool) {
	if board.GetPieceAt(targetSpace).Color == "" {
		move := Move{StartingSpace: startingSpace, TargetSpace: targetSpace}
		return move, true
	} else {
		nextSpace := getNextSpace(startingSpace, targetSpace)

		if board.GetPieceAt(nextSpace).Color == "" {
			move := Move{StartingSpace: startingSpace, TargetSpace: nextSpace}
			return move, true
		} else {
			return Move{}, false
		}
	}
}

func (board *Board) GetPieceAt(space Space) Piece {
	for _, piece := range board.Pieces {
		if sameSpace(piece.Space, space) {
			return piece
		}
	}

	return Piece{}
}

func (board *Board) PlacePiece(piece Piece) (Piece, bool) {
	if board.GetPieceAt(piece.Space).Color == "" {
		board.addPiece(piece)
		return piece, true
	} else {
		return piece, false
	}
}

func (board *Board) MovesForPiece(piece Piece) []Move {
	space := piece.Space

	return board.movesForSpace(space, piece.Color)
}

func (board *Board) ConsolePrint() {
	fmt.Println()
	spaces := []Space{}

	for _, piece := range board.Pieces {
		spaces = append(spaces, piece.Space)
	}

	for i := 0; i < 64; i++ {
		space := SpaceForIndex(i)
		piece := board.GetPieceAt(space)

		if piece.Color == "" {
			fmt.Print("|_")
		} else {
			fmt.Printf("|%c", piece.Color[0])
			// fmt.Print(piece.Space.Rank)
		}
		if i % 8 == 7 {
			fmt.Println("|")
			// fmt.Println("==done with row \t%s==", space)
		}
	}
	fmt.Println()
	return
}

// private

func (board *Board) addPiece(piece Piece) {
	board.Pieces = append(board.Pieces, piece)
}

func (board *Board) movesForSpace(startingSpace Space, color string) []Move {
	moves := []Move{}
	nextRank := 0
	if color == "white" {
		nextRank = startingSpace.Rank + 1
	} else {
		nextRank = startingSpace.Rank - 1
	}

	if notOnLeftEdge(startingSpace) {
		leftFile := decFile(startingSpace.File)
		targetSpace := Space{File: leftFile, Rank: nextRank}
		if nextMove, ok := board.getNextMove(startingSpace, targetSpace); ok {
			moves = append(moves, nextMove)
		}
	}

	if notOnRightEdge(startingSpace) {
		rightFile := incFile(startingSpace.File)
		targetSpace := Space{File: rightFile, Rank: nextRank}

		if nextMove, ok := board.getNextMove(startingSpace, targetSpace); ok {
			moves = append(moves, nextMove)
		}
	}

	return moves
}

func (board *Board) createPieceAtIndex(pieceIndex int, spaceIndex int) (Piece, bool) {
	if initialPiece, ok := initialPieceAtIndex(spaceIndex); ok {
		board.Pieces[pieceIndex] = initialPiece
		return initialPiece, true
	}

	return Piece{}, false
}

func initialPieceAtIndex(index int) (Piece, bool) {
	pieceColor := initialPieceColorFor(index)
	spaceColor := SpaceColorForIndex(index)

	if pieceColor == "" || spaceColor == "white" {
		return Piece{}, false
	} else {
		piece := Piece{Color: pieceColor, Space: SpaceForIndex(index)}
		return piece, true
	}
}
