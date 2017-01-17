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

	expectedBoard := displayBoardTest.emptyStringBoard()
	stringBoard := displayBoard.show(board)

	assert.Equal(t, expectedBoard, stringBoard)
}

func TestDisplaysAMarkedBoard(t *testing.T) {
	displayBoard, board := Board{}, markedBoard

	expectedBoard := displayBoardTest.markedStringBoard()
	stringBoard := displayBoard.show(board)

	assert.Equal(t, expectedBoard, stringBoard)
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
