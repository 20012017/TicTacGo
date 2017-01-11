package ttt

import "io"

type CliDisplay struct {
	writer io.Writer
	script *Script
}

func NewDisplay(writer io.Writer, script *Script) CliDisplay {
	return CliDisplay{writer, script}
}

func (display CliDisplay) write(message string) {
	io.WriteString(display.writer, message)
}

func (display CliDisplay) welcome() {
	display.write(display.script.welcome())
}

func (display CliDisplay) prompt() {
	display.write(display.script.prompt())
}

func (display CliDisplay) showBoard(board Board) {
	displayBoard := new(DisplayBoard).show(board)
	display.write(displayBoard)
}

func (display CliDisplay) draw() {
	display.write(display.script.draw())
}

func (display CliDisplay) goodbye() {
	display.write(display.script.goodbye())
}
