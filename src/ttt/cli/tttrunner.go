package cli

import (
	"bufio"
	"os"
	"ttt/cli/display"
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/cli/players"
)

type TTT struct{}

func (ttt TTT) Start() {
	NewGameRunner(ttt.CreateMenu()).Run()
}

func (ttt TTT) CreateMenu() Menu {
	return NewMenu(
		ttt.createDisplay(),
		ttt.createReader(),
		ttt.createPlayerFactory(),
		ttt.createInputValidator())
}

func (ttt TTT) createInputValidator() validators.Input {
	return validators.NewInputValidator(1, 4)
}

func (ttt TTT) createPlayerFactory() players.Factory {
	return players.NewPlayerFactory(ttt.createReader())
}

func (ttt TTT) createDisplay() DisplayWriter {
	return display.NewDisplayWriter(os.Stdout, new(display.Script))
}

func (ttt TTT) createReader() input.InputReader {
	return input.NewReader(bufio.NewReader(os.Stdin))
}
