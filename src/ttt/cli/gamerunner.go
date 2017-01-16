package cli

type GameRunner struct {
	menu CliMenu
}

func NewGameRunner(menu CliMenu) GameRunner {
	return GameRunner{menu}
}

func (gameRunner GameRunner) Run() {
	game := gameRunner.menu.show()
	game.Start()
}
