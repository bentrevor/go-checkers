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
		Expect(board.GetPieceAt(Space{File: "g", Rank: 3}).Color).To(Equal("white"))
		Expect(board.GetPieceAt(Space{File: "h", Rank: 8}).Color).To(Equal("black"))
	})

	It("knows where a piece can move", func() {
		whitePiece := board.GetPieceAt(Space{File: "g", Rank: 3})
		blackPiece := board.GetPieceAt(Space{File: "h", Rank: 6})
		fmt.Println(whitePiece)
		fmt.Println(board)

		whiteMove1 := Space{File: "f", Rank: 4}
		whiteMove2 := Space{File: "h", Rank: 4}
		blackMove := Space{File: "g", Rank: 5}

		Expect(board.MovesForPiece(blackPiece)).To(Equal([]Space{blackMove}))
		Expect(board.MovesForPiece(whitePiece)).To(Equal([]Space{whiteMove1, whiteMove2}))
	})

	It("knows where a piece can jump", func() {
		whitePiece := board.GetPieceAt(Space{File: "g", Rank: 3})
		blackPiece := board.GetPieceAt(Space{File: "h", Rank: 6})

		whiteMove1 := Space{File: "f", Rank: 4}
		whiteMove2 := Space{File: "h", Rank: 4}
		blackMove := Space{File: "g", Rank: 5}

		Expect(board.MovesForPiece(whitePiece)).To(Equal([]Space{whiteMove1, whiteMove2}))
		Expect(board.MovesForPiece(blackPiece)).To(Equal([]Space{blackMove}))
	})
})
