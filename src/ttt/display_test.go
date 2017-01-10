package ttt

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDisplayWritesToWriter(t *testing.T) {
	buffer, display := setUpDisplay()
	display.write("hello")
	assert.Equal(t, "hello", buffer.String())
}

func TestDisplayWelcomesPlayer(t *testing.T) {
	buffer, display := setUpDisplay()
	display.welcome()
	assert.Equal(t, "Welcome to Tic Tac Toe", buffer.String())
}

func TestDisplayPrintsBoard(t *testing.T) {
	buffer, display := setUpDisplay()
	board := NewBoard(9)
	display.showBoard(board)
	assert.Equal(t, newBoard(), buffer.String())
}

func TestDisplayPromptsForMove(t *testing.T) {
	buffer, display := setUpDisplay()
	display.prompt()
	assert.Equal(t, "Where would you like to make a move?\nPlease choose a space between 1 and 9", buffer.String())
}

func setUpDisplay() (*bytes.Buffer, Display) {
	buffer, script := new(bytes.Buffer), new(Script)
	display := NewDisplay(buffer, script)
	return buffer, display
}

func newBoard() string {
	return `-------------
|[-]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
`
}
