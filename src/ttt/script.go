package ttt

type Script struct{}

func (script Script) welcome() string {
	return "Welcome to Tic Tac Toe"
}

func (script Script) prompt() string {
	return "Where would you like to make a move?\nPlease choose a space between 1 and 9"
}

func (script Script) draw() string {
	return "It's a draw!"
}

func (script Script) goodbye() string {
	return "goodbye!"
}
