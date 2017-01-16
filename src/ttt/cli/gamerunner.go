package cli

type GameRunner struct {
	menu CliMenu
}

func NewGameRunner(menu CliMenu) GameRunner {
	return GameRunner{menu}
}

func (gameRunner GameRunner) Run() {
	replay := true
	for replay == true {
		game := gameRunner.menu.show()
		game.Start()
		replay = gameRunner.menu.replay()
	}
}
