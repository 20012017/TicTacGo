package ttt

type Game struct {
	playerOne, playerTwo Player
	board                Board
	rules                rules
}

func NewGame(playerOne, playerTwo Player, board Board, rule rules) Game {
	return Game{playerOne, playerTwo, board, rule}
}

func (game Game) play(move int) Board {
	game.board = game.playMove(move)
	return game.board
}

func (game Game) currentMark() string {
	if game.board.countMarks()%2 == 0 {
		return game.playerOne.mark()
	}
	return game.playerTwo.mark()
}

func (game Game) isOver() bool {
	over := game.isAWin(game.data()) ||
		game.rules.isADraw(game.data())
	return over
}

func (game Game) result() (bool, string) {
	if game.isOver() && game.isAWin(game.data()) {
		return true, game.winner()
	}
	return false, ""
}

func (game Game) winner() string {
	return game.rules.winner(game.data())
}

func (game Game) playMove(move int) Board {
	board := game.board
	markedBoard := board.placeMark(move, game.currentMark())
	return markedBoard
}

func (game Game) isAWin(board Board, markOne, markTwo string) bool {
	return game.rules.isAWin(board, markOne, markTwo)
}

func (game Game) data() (Board, string, string) {
	return game.board, game.markOne(), game.markTwo()
}

func (game Game) markOne() string {
	return game.playerOne.mark()
}

func (game Game) markTwo() string {
	return game.playerTwo.mark()
}
