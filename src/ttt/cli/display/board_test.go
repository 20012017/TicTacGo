package display

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

type DisplayBoardTest struct{}

var displayBoardTest DisplayBoardTest = DisplayBoardTest{}
var markedBoard core.Board = core.NewMarkedBoard([]string{"X", "", "", "", "", "", "", "", ""})

func TestDisplaysAnEmptyBoard(t *testing.T) {
	displayBoard, board := Board{}, core.NewBoard(9)

	emptyStringBoard := displayBoardTest.emptyStringBoard()

	assert.Equal(t, emptyStringBoard, displayBoard.show(board))
}

func TestDisplaysAMarkedBoard(t *testing.T) {
	displayBoard, board := Board{}, markedBoard

	markedStringBoard := displayBoardTest.markedStringBoard()

	assert.Equal(t, markedStringBoard, displayBoard.show(board))
}

func (displayBoardTest DisplayBoardTest) emptyStringBoard() string {
	return "-------------\n" +
		"|[-]|[-]|[-]|\n" +
		"-------------\n" +
		"|[-]|[-]|[-]|\n" +
		"-------------\n" +
		"|[-]|[-]|[-]|\n" +
		"-------------\n"
}

func (displayBoardTest DisplayBoardTest) markedStringBoard() string {
	return "-------------\n" +
		"|[X]|[-]|[-]|\n" +
		"-------------\n" +
		"|[-]|[-]|[-]|\n" +
		"-------------\n" +
		"|[-]|[-]|[-]|\n" +
		"-------------\n"
}
