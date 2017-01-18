package cli

import (
	"ttt/core"
	"ttt/core/marks"
)

type DisplayWriter interface {
	Write(message string)

	Welcome()

	HumanPrompt()

	ComputerPrompt()

	ShowBoard(board core.Board)

	Draw()

	Goodbye()

	Win(mark marks.Mark)

	Clear()

	Menu()

	InvalidChoice()

	Replay()
}
