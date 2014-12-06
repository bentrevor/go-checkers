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
		Expect(board.GetPieceAt(Space{File: "g", Rank: "3"}).Color).To(Equal("white"))
		Expect(board.GetPieceAt(Space{File: "h", Rank: "8"}).Color).To(Equal("black"))
	})
})
