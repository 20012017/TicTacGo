package display

import (
	"io"
	"ttt/core"
	"ttt/core/marks"
)

type Writer struct {
	ioWriter io.Writer
	script   *Script
}

func NewDisplayWriter(writer io.Writer, script *Script) Writer {
	return Writer{writer, script}
}

func (writer Writer) Write(message string) {
	io.WriteString(writer.ioWriter, message)
}

func (writer Writer) Welcome() {
	writer.Write(writer.script.welcome())
}

func (writer Writer) HumanPrompt() {
	writer.Write(writer.script.prompt())
}

func (writer Writer) ComputerPrompt() {
	writer.Write(writer.script.computerPrompt())
}

func (writer Writer) ShowBoard(board core.Board) {
	displayBoard := new(Board).show(board)
	writer.Write(displayBoard)
}

func (writer Writer) Draw() {
	writer.Write(writer.script.draw())
}

func (writer Writer) Goodbye() {
	writer.Write(writer.script.goodbye())
}

func (writer Writer) Win(mark marks.Mark) {
	writer.Write(writer.script.win(string(mark)))
}

func (writer Writer) Clear() {
	writer.Write("\033[2J\033[;H")
}

func (writer Writer) Menu() {
	writer.Write(writer.script.menu())
}

func (writer Writer) InvalidChoice() {
	writer.Write(writer.script.invalidChoice())
}

func (writer Writer) Replay() {
	writer.Write(writer.script.replay())
}
