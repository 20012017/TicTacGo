package cli

import "ttt/core"

type Display interface {
	write(message string)

	welcome()

	prompt()

	showBoard(board core.TTTBoard)

	draw()

	goodbye()

	win(mark string)

	clear()
}
