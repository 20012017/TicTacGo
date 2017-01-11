package ttt

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DisplayTest struct{}

var displayTest DisplayTest = DisplayTest{}

func TestDisplayWritesToWriter(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	display.write("hello")

	assert.Equal(t, "hello", buffer.String())
}

func TestDisplayWelcomesPlayer(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	display.welcome()

	assert.Equal(t, "Welcome to Tic Tac Toe", buffer.String())
}

func TestDisplayPrintsBoard(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	board := NewBoard(9)
	display.showBoard(board)

	assert.Equal(t, displayTest.newBoard(), buffer.String())
}

func TestDisplayPromptsForMove(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	display.prompt()

	assert.Equal(t, "Where would you like to make a move?\nPlease choose a space between 1 and 9", buffer.String())
}

func (displayTest DisplayTest) setUpDisplay() (*bytes.Buffer, CliDisplay) {
	buffer, script := new(bytes.Buffer), new(Script)
	display := NewDisplay(buffer, script)
	return buffer, display
}

func (displayTest DisplayTest) newBoard() string {
	return `-------------
|[-]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
`
}
