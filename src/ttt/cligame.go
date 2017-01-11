package ttt

type CliGame struct {
	game                 Game
	playerOne, playerTwo Player
	display              Display
}

func NewCliGame(game Game, playerOne, playerTwo Player, display Display) CliGame {
	return CliGame{game, playerOne, playerTwo, display}
}

func (cliGame CliGame) start() {
	cliGame.welcome()

}

func (cliGame CliGame) welcome() {
	cliGame.display.welcome()
	cliGame.display.showBoard(cliGame.game.board)
	cliGame.display.prompt()
}
