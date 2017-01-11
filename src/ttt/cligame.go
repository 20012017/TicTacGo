package ttt

type CliGame struct {
	game    Game
	display Display
}

func NewCliGame(game Game, display Display) CliGame {
	return CliGame{game, display}
}

func (cliGame CliGame) start() {
	cliGame.welcome()

	for !cliGame.game.isOver() {
		move, _ := cliGame.game.currentPlayer().move(cliGame.game.board)
		cliGame.game.play(move)
	}
	cliGame.display.draw()
	cliGame.display.goodbye()
}

func (cliGame CliGame) welcome() {
	cliGame.display.welcome()
	cliGame.display.showBoard(cliGame.game.board)
	cliGame.display.prompt()
}
