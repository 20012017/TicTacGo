package cli

import (
	"io"
	"ttt/core"
)

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

func (display CliDisplay) showBoard(board core.TTTBoard) {
	displayBoard := new(DisplayBoard).show(board)
	display.write(displayBoard)
}

func (display CliDisplay) draw() {
	display.write(display.script.draw())
}

func (display CliDisplay) goodbye() {
	display.write(display.script.goodbye())
}

func (display CliDisplay) win(mark string) {
	display.write(display.script.win(mark))
}

func (display CliDisplay) clear() {
	display.write("\033[2J\033[;H")
}
