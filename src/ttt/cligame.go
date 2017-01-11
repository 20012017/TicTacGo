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
	cliGame.showBoard()
	cliGame.displayResult(cliGame.game.result())
	cliGame.display.goodbye()
}

func (cliGame CliGame) displayResult(isWon bool, winner string) {
	if isWon {
		cliGame.display.win(winner)
	} else {
		cliGame.display.draw()
	}
}

func (cliGame CliGame) play() {
	for !cliGame.game.isOver() {
		cliGame.initializeTurn()
		move, _ := cliGame.currentPlayer().move(cliGame.board())
		cliGame.game.play(move)
	}
}

func (cliGame CliGame) initializeTurn() {
	cliGame.showBoard()
	cliGame.display.prompt()
}

func (cliGame CliGame) welcome() {
	cliGame.display.clear()
	cliGame.display.welcome()
}

func (cliGame CliGame) showBoard() {
	cliGame.display.clear()
	cliGame.display.showBoard(cliGame.board())
}

func (cliGame CliGame) board() Board {
	return cliGame.game.board
}

func (cliGame CliGame) currentPlayer() Player {
	return cliGame.game.currentPlayer()
}
