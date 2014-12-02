package checkers_test

import (
	. "github.com/bentrevor/checkers/src"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("a board", func() {
	var board = NewGameBoard()

	It("has pieces", func() {
		Expect(len(board.Pieces)).To(Equal(24))
	})

	// It("knows where a piece can move", func() {
	// 	space := Space{File: "g", Rank: "3"}
	// 	moves := board.MovesFor(space)
	// 	Expect(len(moves)).To(Equal(2))

	// 	Expect(moves[0]).To(Equal("f4"))
	// 	Expect(moves[1]).To(Equal("h4"))
	// })
})
