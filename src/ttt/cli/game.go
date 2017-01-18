package cli

import (
	"fmt"
	"ttt/core"
)

type Game struct {
	game    *core.Game
	display DisplayWriter
}

func NewCliGame(game *core.Game, display DisplayWriter) Game {
	return Game{game, display}
}

func (cliGame Game) Start() {
	cliGame.welcome()
	cliGame.play()
	cliGame.end()
}

func (cliGame Game) end() {
	cliGame.showBoard()
	cliGame.displayResult(cliGame.game.Result())
}

func (cliGame Game) displayResult(isWon bool, winner string) {
	if isWon {
		cliGame.display.Win(winner)
	} else {
		cliGame.display.Draw()
	}
}

func (cliGame *Game) play() {
	for !cliGame.game.IsOver() {
		cliGame.initializeTurn()
		cliGame.game.Play(cliGame.getValidMove())
	}
}

func (cliGame Game) getValidMove() int {
	move, err := cliGame.currentPlayer().Move(cliGame.board())
	for err != nil {
		cliGame.display.Write(fmt.Sprintf("%s\n", err.Error()))
		cliGame.display.HumanPrompt()
		move, err = cliGame.currentPlayer().Move(cliGame.board())
	}
	return move
}

func (cliGame Game) initializeTurn() {
	cliGame.showBoard()
	cliGame.display.HumanPrompt()
}

func (cliGame Game) welcome() {
	cliGame.display.Clear()
	cliGame.display.Welcome()
}

func (cliGame Game) showBoard() {
	cliGame.display.Clear()
	cliGame.display.ShowBoard(cliGame.board())
}

func (cliGame Game) board() core.Board {
	return cliGame.game.Board()
}

func (cliGame Game) currentPlayer() core.Player {
	return cliGame.game.CurrentPlayer()
}

func (cliGame Game) Game() *core.Game {
	return cliGame.game
}
