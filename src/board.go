package checkers

type Piece struct {
	Color Color
	Space Space
}

type Board struct {
	Pieces []Piece
}

func NewEmptyBoard() Board {
	return Board{make([]Piece, 24, 32)}
}

func NewGameBoard() Board {
	board := NewEmptyBoard()
	board.placeInitialPieces()
	return board
}

func (board *Board) GetPieceAtSpace(space Space) Piece {
	for _, piece := range board.Pieces {
		if SameSpace(piece.Space, space) {
			return piece
		}
	}

	return Piece{}
}

func (board *Board) MakeMove(move Move) {
	color := board.GetPieceAtSpace(move.StartingSpace).Color

	if color == NoColor {
		// error!
	} else {
		piece := Piece{Color: color, Space: move.TargetSpace}
		board.PlacePiece(piece)
		board.RemovePieceAtSpace(move.StartingSpace)
	}
}

func (board *Board) RemovePieceAtSpace(space Space) {
	for i, piece := range board.Pieces {
		if SameSpace(piece.Space, space) {
			board.Pieces[i] = Piece{}
		}
	}
}

func (board *Board) PlacePiece(piece Piece) (Piece, bool) {
	if board.GetPieceAtSpace(piece.Space).Color == NoColor {
		board.addPiece(piece)
		return piece, true
	} else {
		return piece, false
	}
}

func reverseRows(rows []string) []string {
	length := len(rows)
	reversed := make([]string, length)

	for i := 0; i < length; i++ {
		reversed[length-(i+1)] = rows[i]
	}

	return reversed
}

func (board *Board) placeInitialPieces() {
	pieceIndex := 0

	for i := 0; i < 64; i++ {
		_, created := board.createInitialPieceAtIndex(pieceIndex, i)

		if created && pieceIndex < cap(board.Pieces) {
			pieceIndex += 1
		}
	}
}

func (board *Board) createInitialPieceAtIndex(pieceIndex int, spaceIndex int) (Piece, bool) {
	if initialPiece, ok := initialPieceAtIndex(spaceIndex); ok {
		board.Pieces[pieceIndex] = initialPiece
		return initialPiece, true
	}

	return Piece{}, false
}

func initialPieceAtIndex(index int) (Piece, bool) {
	pieceColor := initialPieceColorFor(index)
	spaceColor := SpaceColorForIndex(index)

	if pieceColor == NoColor || spaceColor == White {
		return Piece{}, false
	} else {
		piece := Piece{Color: pieceColor, Space: SpaceForIndex(index)}
		return piece, true
	}
}

func (board *Board) addPiece(piece Piece) {
	board.Pieces = append(board.Pieces, piece)
}
