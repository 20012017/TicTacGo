package cli

import (
	"fmt"
	"ttt/core"
	"ttt/core/marks"
)

type Game struct {
	game       *core.Game
	display    DisplayWriter
	prompter   Prompter
	gameChoice int
}

func NewCliGame(game *core.Game, display DisplayWriter, prompter Prompter, gameChoice int) Game {
	return Game{game, display, prompter, gameChoice}
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

func (cliGame Game) displayResult(isWon bool, winner marks.Mark) {
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
		cliGame.showBoard()
		cliGame.display.Write(fmt.Sprintf("%s\n", err.Error()))
		cliGame.prompt()
		move, err = cliGame.currentPlayer().Move(cliGame.board())
	}
	return move
}

func (cliGame Game) initializeTurn() {
	cliGame.showBoard()
	cliGame.prompt()
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

func (cliGame Game) currentMark() marks.Mark {
	return cliGame.game.CurrentPlayer().Mark()
}

func (cliGame Game) Game() *core.Game {
	return cliGame.game
}

func (cliGame Game) prompt() {
	cliGame.prompter.Prompt(cliGame.gameChoice, cliGame.currentMark(), cliGame.display)
}
