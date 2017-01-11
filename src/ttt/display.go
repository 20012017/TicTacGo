package ttt

import "io"

type Display struct {
	writer io.Writer
	script *Script
}

func NewDisplay(writer io.Writer, script *Script) Display {
	return Display{writer, script}
}

func (display Display) write(message string) {
	io.WriteString(display.writer, message)
}

func (display Display) welcome() {
	display.write(display.script.welcome())
}

func (display Display) prompt() {
	display.write(display.script.prompt())
}

func (display Display) showBoard(board Board) {
	displayBoard := new(DisplayBoard).show(board)
	display.write(displayBoard)
}
