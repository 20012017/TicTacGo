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
	cliGame.play()
	cliGame.end()
}

func (cliGame CliGame) end() {
	isWon, winner := cliGame.game.result()
	if isWon {
		cliGame.display.win(winner)
	} else {
		cliGame.display.draw()
	}
	cliGame.display.goodbye()
}

func (cliGame CliGame) play() {
	for !cliGame.game.isOver() {
		move, _ := cliGame.game.currentPlayer().move(cliGame.game.board)
		cliGame.game.play(move)
	}
}

func (cliGame CliGame) welcome() {
	cliGame.display.welcome()
	cliGame.display.showBoard(cliGame.game.board)
	cliGame.display.prompt()
}
