package checkers

type Player interface {
	GetMove(Board) Move
	Color() Color
}

type MoveDecider interface {
	GetMove(Board) Move
}

type Color string

const (
	NoColor Color = ""
	White   Color = "white"
	Black   Color = "black"
)

type HumanPlayer struct {
	color       Color
	MoveDecider ConsoleInput
}

func NewHumanPlayer(color Color) Player {
	return &HumanPlayer{color: color, MoveDecider: ConsoleInput{}}
}

func (hp *HumanPlayer) GetMove(board Board) Move {
	return hp.MoveDecider.GetMove(board)
}

func (player *HumanPlayer) Color() Color {
	return player.color
}
