package cli

import "ttt/core"

type Display interface {
	write(message string)

	welcome()

	prompt()

	showBoard(board core.Board)

	draw()

	goodbye()

	win(mark string)

	clear()
}
