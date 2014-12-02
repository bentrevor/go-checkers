package checkers_test

import (
	. "github.com/bentrevor/checkers/src"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
)

var _ = Describe("a board", func() {
	var board = NewGameBoard()

	It("has pieces", func() {
		fmt.Println()
		Expect(len(board.Pieces)).To(Equal(24))
	})

	It("gets the piece at a space", func() {
		for x,y := range board.Pieces {
			fmt.Println(x, y)
		}

		Expect(board.GetPieceAt(Space{File: "g", Rank: "3"}).Color).To(Equal("black"))
		Expect(board.GetPieceAt(Space{File: "h", Rank: "8"}).Color).To(Equal("white"))
	})

	// It("knows where a piece can move", func() {
	// 	space := Space{File: "g", Rank: "3"}
	// 	moves := board.MovesFor(space)
	// 	Expect(len(moves)).To(Equal(2))

	// 	Expect(moves[0]).To(Equal("f4"))
	// 	Expect(moves[1]).To(Equal("h4"))
	// })
})
