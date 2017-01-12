package display

type Script struct{}

func (script Script) welcome() string {
	return "Welcome to Tic Tac Toe\n"
}

func (script Script) prompt() string {
	return "Where would you like to make a move?\nPlease choose a space between 1 and 9\n"
}

func (script Script) draw() string {
	return "It's a draw!\n"
}

func (script Script) goodbye() string {
	return "goodbye!\n"
}

func (script Script) win(mark string) string {
	return "X wins!\n"
}
