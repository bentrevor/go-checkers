package checkers

type Game struct {
	Board         Board
	CurrentPlayer Player
	OtherPlayer   Player
	Io            IO
}

type IO interface {
	PrintBoard(Board)
	GetInput() (string, error)
}

func NewGame(io IO) Game {
	return Game{
		Board:         NewGameBoard(),
		CurrentPlayer: NewPlayer("white"),
		OtherPlayer:   NewPlayer("black"),
		Io:            io,
	}
}

func NewGameWithBoard(board Board) Game {
	return Game{
		Board:         board,
		CurrentPlayer: NewPlayer("white"),
		OtherPlayer:   NewPlayer("black"),
	}
}

func (game *Game) NextTurn() {
	game.Io.PrintBoard(game.Board)

	move := game.CurrentPlayer.GetMove(game.Board)
	for game.InvalidInput(move) {
		move = game.CurrentPlayer.GetMove(game.Board)
	}

	game.Board.MakeMove(move)
	game.togglePlayers()
}

func (game *Game) InvalidInput(move Move) bool {
	board := game.Board
	color := game.CurrentPlayer.Color()

	return !IsLegalMove(move, board, color)
}

func (game *Game) Start() {
	for !gameOver(game.Board) {
		game.NextTurn()
	}
}

func (game *Game) CurrentColor() string {
	return game.CurrentPlayer.Color()
}

func (game *Game) togglePlayers() {
	p := game.CurrentPlayer
	game.CurrentPlayer = game.OtherPlayer
	game.OtherPlayer = p
}
