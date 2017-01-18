package core

import "ttt/core/marks"

type Game struct {
	playerOne, playerTwo Player
	board                Board
	rules                *Rules
}

func NewGame(playerOne, playerTwo Player, board Board, rule *Rules) Game {
	return Game{playerOne, playerTwo, board, rule}
}

func (game Game) Board() Board {
	return game.board
}

func (game Game) IsOver() bool {
	return game.rules.IsOver(game.board, game.markOne(), game.markTwo())
}

func (game Game) Result() (bool, marks.Mark) {
	if game.IsOver() && game.IsAWin() {
		return true, game.winner()
	}
	return false, marks.EMPTY
}

func (game *Game) Play(move int) Board {
	game.board = game.playMove(move)
	return game.board
}

func (game Game) CurrentPlayer() Player {
	if game.board.CountMarks()%2 == 0 {
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

func (game Game) currentMark() marks.Mark {
	return game.CurrentPlayer().Mark()
}

func (game Game) winner() marks.Mark {
	return game.winningMark()
}

func (game Game) playMove(move int) Board {
	mark := game.CurrentPlayer().Mark()
	return game.board.PlaceMark(move, mark)
}

func (game Game) winningMark() marks.Mark {
	return game.rules.Winner(game.data())
}

func (game Game) data() (Board, marks.Mark, marks.Mark) {
	return game.board, game.markOne(), game.markTwo()
}

func (game Game) markOne() marks.Mark {
	return game.playerOne.Mark()
}

func (game Game) markTwo() marks.Mark {
	return game.playerTwo.Mark()
}
