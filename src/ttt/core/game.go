package core

type Game struct {
	playerOne, playerTwo Player
	board                TTTBoard
	rules                *Rules
}

func NewGame(playerOne, playerTwo Player, board TTTBoard, rule *Rules) Game {
	return Game{playerOne, playerTwo, board, rule}
}

func (game Game) Board() TTTBoard {
	return game.board
}

func (game Game) IsOver() bool {
	over := game.IsAWin() ||
		game.IsADraw()
	return over
}

func (game Game) Result() (bool, string) {
	if game.IsOver() && game.IsAWin() {
		return true, game.winner()
	}
	return false, ""
}

func (game Game) Play(move int) TTTBoard {
	game.board = game.playMove(move)
	return game.board
}

func (game Game) CurrentPlayer() Player {
	if game.board.countMarks()%2 == 0 {
		return game.playerOne
	}
	return game.playerTwo
}

func (game Game) IsADraw() bool {
	return game.rules.isADraw(game.data())
}

func (game Game) IsAWin() bool {
	return game.rules.isAWin(game.data())
}

func (game Game) currentMark() string {
	return game.CurrentPlayer().Mark()
}

func (game Game) winner() string {
	return game.winningMark()
}

func (game Game) playMove(move int) TTTBoard {
	board := game.board
	markedBoard := board.placeMark(move, game.currentMark())
	return markedBoard
}

func (game Game) winningMark() string {
	return game.rules.winner(game.data())
}

func (game Game) data() (TTTBoard, string, string) {
	return game.board, game.markOne(), game.markTwo()
}

func (game Game) markOne() string {
	return game.playerOne.Mark()
}

func (game Game) markTwo() string {
	return game.playerTwo.Mark()
}
