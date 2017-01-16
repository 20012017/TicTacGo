package display

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

type WriterTest struct{}

var writerTest WriterTest = WriterTest{}

func TestWritesToWriter(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.Write("hello")

	assert.Equal(t, "hello", buffer.String())
}

func TestWelcomesPlayer(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.Welcome()

	assert.Equal(t,
		"**********************Welcome to Tic Tac Toe**********************\n",
		buffer.String())
}

func TestPrintsBoard(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	board := core.NewBoard(9)
	writer.ShowBoard(board)

	assert.Equal(t, writerTest.newBoard(), buffer.String())
}

func TestPromptsForMove(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.Prompt()

	assert.Equal(t,
		"Where would you like to make a move?\nPlease choose a space between 1 and 9\n",
		buffer.String())
}

func TestDraw(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.Draw()

	assert.Equal(t, "It's a draw!\n", buffer.String())
}

func TestGoodbye(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.Goodbye()

	assert.Equal(t, "goodbye!\n", buffer.String())
}

func TestWin(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.Win("X")

	assert.Equal(t, "X wins!\n", buffer.String())
}

func TestClearScreen(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.Clear()

	assert.Equal(t, "\033[2J\033[;H", buffer.String())
}

func TestMenu(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.Menu()

	assert.Equal(t,
		"Please choose a game option:\n1: Human v Human\n2: Human v Computer\n3: Computer v Human\n4: Computer v Computer\n",
		buffer.String())
}

func TestInvalidChoice(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.InvalidChoice()

	assert.Equal(t, "Please enter a number between 1 and 4\n", buffer.String())
}

func TestReplay(t *testing.T) {
	buffer, writer := writerTest.setUpWriter()

	writer.Replay()

	assert.Equal(t, "Would you like to play again?\nPlease type yes to replay\n", buffer.String())
}

func (writerTest WriterTest) setUpWriter() (*bytes.Buffer, Writer) {
	buffer, script := new(bytes.Buffer), new(Script)
	writer := NewDisplayWriter(buffer, script)
	return buffer, writer
}

func (writerTest WriterTest) newBoard() string {
	return `-------------
|[-]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
`
}
