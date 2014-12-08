package checkers_test

import (
	. "github.com/bentrevor/checkers/src"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
)

var _ = Describe("a board", func() {
	var board = NewGameBoard()
	fmt.Println()

	It("has pieces", func() {
		Expect(len(board.Pieces)).To(Equal(24))
	})

	It("gets the piece at a space", func() {
		whiteSpaces := []Space{
			Space{File: "a", Rank: 1},
			Space{File: "c", Rank: 1},
			Space{File: "e", Rank: 1},
			Space{File: "g", Rank: 1},
			Space{File: "b", Rank: 2},
			Space{File: "d", Rank: 2},
			Space{File: "f", Rank: 2},
			Space{File: "h", Rank: 2},
			Space{File: "a", Rank: 3},
			Space{File: "c", Rank: 3},
			Space{File: "e", Rank: 3},
			Space{File: "g", Rank: 3},
		}

		blackSpaces := []Space{
			Space{File: "b", Rank: 6},
			Space{File: "d", Rank: 6},
			Space{File: "f", Rank: 6},
			Space{File: "h", Rank: 6},
			Space{File: "a", Rank: 7},
			Space{File: "c", Rank: 7},
			Space{File: "e", Rank: 7},
			Space{File: "g", Rank: 7},
			Space{File: "b", Rank: 8},
			Space{File: "d", Rank: 8},
			Space{File: "f", Rank: 8},
			Space{File: "h", Rank: 8},
		}

		for _, space := range whiteSpaces {
			Expect(board.GetPieceAt(space).Color).To(Equal("white"))
		}

		for _, space := range blackSpaces {
			Expect(board.GetPieceAt(space).Color).To(Equal("black"))
		}
	})

	It("can place a piece", func() {
		emptySpace := Space{File: "e", Rank: 5}
		occupiedSpace := Space{File: "e", Rank: 7}

		piece1 := Piece{Color: "white", Space: emptySpace}
		piece2 := Piece{Color: "white", Space: occupiedSpace}

		_, createdPieceAtE5 := board.PlacePiece(piece1)
		_, createdPieceAtE7 := board.PlacePiece(piece2)

		Expect(createdPieceAtE5).To(BeTrue())
		Expect(createdPieceAtE7).To(BeFalse())
	})

	It("knows where a piece can move", func() {
		whitePiece := board.GetPieceAt(Space{File: "g", Rank: 3})
		blackPiece := board.GetPieceAt(Space{File: "h", Rank: 6})

		whiteMove1 := Space{File: "f", Rank: 4}
		whiteMove2 := Space{File: "h", Rank: 4}
		blackMove := Space{File: "g", Rank: 5}

		Expect(board.MovesForPiece(whitePiece)).To(Equal([]Space{whiteMove1, whiteMove2}))
		Expect(board.MovesForPiece(blackPiece)).To(Equal([]Space{blackMove}))
	})

	It("knows where a piece can jump", func() {
		d6 := board.GetPieceAt(Space{File: "d", Rank: 6})
		f6 := board.GetPieceAt(Space{File: "f", Rank: 6})
		emptySpace := Space{File: "e", Rank: 5}
		whitePiece := Piece{Color: "white", Space: emptySpace}

		board.PlacePiece(whitePiece)

		captureMoveForD6 := Space{File: "f", Rank: 4}
		captureMoveForF6 := Space{File: "d", Rank: 4}

		d6Moves := board.MovesForPiece(d6)
		f6Moves := board.MovesForPiece(f6)

		Expect(Includes(d6Moves, captureMoveForD6)).To(BeTrue())
		Expect(Includes(f6Moves, captureMoveForF6)).To(BeTrue())
	})
})
