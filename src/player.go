package checkers

type Player interface {
	GetMove(Board) Move
	Color() string
}

type MoveDecider interface {
	GetMove(Board) Move
}

type HumanPlayer struct {
	color       string
	MoveDecider ConsoleInput
}

func NewHumanPlayer(color string) Player {
	return &HumanPlayer{color: color, MoveDecider: ConsoleInput{}}
}

func MoveFromString(input string) Move {
	return Move{StartingSpace: NewSpace(input[0:2]), TargetSpace: NewSpace(input[5:7])}
}

func (hp *HumanPlayer) GetMove(board Board) Move {
	return hp.MoveDecider.GetMove(board)
}

func (player *HumanPlayer) Color() string {
	return player.color
}
