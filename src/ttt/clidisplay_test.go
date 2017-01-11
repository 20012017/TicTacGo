package ttt

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DisplayTest struct{}

var displayTest DisplayTest = DisplayTest{}

func TestWritesToWriter(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	display.write("hello")

	assert.Equal(t, "hello", buffer.String())
}

func TestWelcomesPlayer(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	display.welcome()

	assert.Equal(t, "Welcome to Tic Tac Toe", buffer.String())
}

func TestPrintsBoard(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	board := NewBoard(9)
	display.showBoard(board)

	assert.Equal(t, displayTest.newBoard(), buffer.String())
}

func TestPromptsForMove(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	display.prompt()

	assert.Equal(t, "Where would you like to make a move?\nPlease choose a space between 1 and 9", buffer.String())
}

func TestDraw(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	display.draw()

	assert.Equal(t, "It's a draw!", buffer.String())
}

func TestGoodbye(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	display.goodbye()

	assert.Equal(t, "goodbye!", buffer.String())
}

func TestWin(t *testing.T) {
	buffer, display := displayTest.setUpDisplay()

	display.win("X")

	assert.Equal(t, "X wins!", buffer.String())
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
