package checkers

type Game struct {
	Board         IBoard
	CurrentPlayer Player
	OtherPlayer   Player
}

func NewGame() Game {
	return Game{
		Board:         NewGameBoard(),
		CurrentPlayer: NewPlayer("white"),
		OtherPlayer:   NewPlayer("black"),
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
	game.Board.ConsolePrint()

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
