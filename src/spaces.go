package checkers

import (
	"strconv"
)

type Space struct {
	Rank int
	File string
}

func NewSpace(coordinates string) Space {
	file := string(coordinates[0])
	rank, _ := strconv.Atoi(string(coordinates[1]))

	return Space{File: file, Rank: rank}
}

// black squares
var A1 = NewSpace("a1")
var A3 = NewSpace("a3")
var A5 = NewSpace("a5")
var A7 = NewSpace("a7")
var B2 = NewSpace("b2")
var B4 = NewSpace("b4")
var B6 = NewSpace("b6")
var B8 = NewSpace("b8")
var C1 = NewSpace("c1")
var C3 = NewSpace("c3")
var C5 = NewSpace("c5")
var C7 = NewSpace("c7")
var D2 = NewSpace("d2")
var D4 = NewSpace("d4")
var D6 = NewSpace("d6")
var D8 = NewSpace("d8")
var E1 = NewSpace("e1")
var E3 = NewSpace("e3")
var E5 = NewSpace("e5")
var E7 = NewSpace("e7")
var F2 = NewSpace("f2")
var F4 = NewSpace("f4")
var F6 = NewSpace("f6")
var F8 = NewSpace("f8")
var G1 = NewSpace("g1")
var G3 = NewSpace("g3")
var G5 = NewSpace("g5")
var G7 = NewSpace("g7")
var H2 = NewSpace("h2")
var H4 = NewSpace("h4")
var H6 = NewSpace("h6")
var H8 = NewSpace("h8")

// white squares
// var A2 = NewSpace("a2")
// var A4 = NewSpace("a4")
// var A6 = NewSpace("a6")
// var A8 = NewSpace("a8")
// var B1 = NewSpace("b1")
// var B3 = NewSpace("b3")
// var B5 = NewSpace("b5")
// var B7 = NewSpace("b7")
// var C2 = NewSpace("c2")
// var C4 = NewSpace("c4")
// var C6 = NewSpace("c6")
// var C8 = NewSpace("c8")
// var D1 = NewSpace("d1")
// var D3 = NewSpace("d3")
// var D5 = NewSpace("d5")
// var D7 = NewSpace("d7")
// var E2 = NewSpace("e2")
// var E4 = NewSpace("e4")
// var E6 = NewSpace("e6")
// var E8 = NewSpace("e8")
// var F1 = NewSpace("f1")
// var F3 = NewSpace("f3")
// var F5 = NewSpace("f5")
// var F7 = NewSpace("f7")
// var G2 = NewSpace("g2")
// var G4 = NewSpace("g4")
// var G6 = NewSpace("g6")
// var G8 = NewSpace("g8")
// var H1 = NewSpace("h1")
// var H3 = NewSpace("h3")
// var H5 = NewSpace("h5")
// var H7 = NewSpace("h7")
