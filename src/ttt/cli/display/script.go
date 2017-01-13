package display

import "fmt"

type Script struct{}

func (script Script) welcome() string {
	return "**********************Welcome to Tic Tac Toe**********************\n"
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
	return fmt.Sprintf("%s wins!\n", mark)
}

func (script Script) menu() string {
	return "Please choose a game option:\n1: Human v Human\n2: Human v Computer\n3: Computer v Human\n4: Computer v Computer\n"
}

func (script Script) invalidChoice() string {
	return "Please enter a number between 1 and 4\n"
}
