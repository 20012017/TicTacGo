package ttt

type Display interface {
	write(message string)

	welcome()

	prompt()

	showBoard(board Board)

	draw()

	goodbye()

	win(mark string)

	clear()
}
