package cli

import "ttt/core"

type DisplayWriter interface {
	Write(message string)

	Welcome()

	HumanPrompt()

	ComputerPrompt()

	ShowBoard(board core.Board)

	Draw()

	Goodbye()

	Win(mark string)

	Clear()

	Menu()

	InvalidChoice()

	Replay()
}
