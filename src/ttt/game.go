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
	board, markOne, markTwo := game.board, game.markOne(), game.markTwo()
	over := game.rules.isAWin(board, markOne, markTwo) ||
		game.rules.isADraw(board, markOne, markTwo)
	return over
}

func (game Game) result() (bool, string) {
	if game.isOver() && game.rules.isAWin(game.board, game.markOne(), game.markTwo()) {
		return true, game.winner()
	}
	return false, ""
}

func (game Game) winner() string {
	board, markOne, markTwo := game.board, game.markOne(), game.markTwo()
	return game.rules.winner(board, markOne, markTwo)
}

func (game Game) playMove(move int) Board {
	board := game.board
	markedBoard := board.placeMark(move, game.currentMark())
	return markedBoard
}

func (game Game) markOne() string {
	return game.playerOne.mark()
}

func (game Game) markTwo() string {
	return game.playerTwo.mark()
}
