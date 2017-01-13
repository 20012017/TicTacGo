package cli

import "ttt/cli/input"

type Menu struct {
	display     DisplayWriter
	inputReader input.InputReader
}

func NewMenu(display DisplayWriter, inputReader input.InputReader) Menu {
	return Menu{display, inputReader}
}

func (menu Menu) show() {
	menu.display.Menu()
	menu.inputReader.Read()
}
