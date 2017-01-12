package display

import "ttt/core"

type DisplayWriter interface {
	Write(message string)

	Welcome()

	Prompt()

	ShowBoard(board core.TTTBoard)

	Draw()

	Goodbye()

	Win(mark string)

	Clear()
}
