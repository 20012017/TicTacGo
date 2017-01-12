package display

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

func TestDisplaysAnEmptyBoard(t *testing.T) {
	displayBoard, board := Board{}, core.NewBoard(9)
	stringBoard :=
		`-------------
|[-]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
`
	assert.Equal(t, stringBoard, displayBoard.show(board))
}

func TestDisplaysAMarkedBoard(t *testing.T) {
	displayBoard, board := Board{}, core.NewMarkedBoard([]string{"X", "", "", "", "", "", "", "", ""})
	stringBoard :=
		`-------------
|[X]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
|[-]|[-]|[-]|
-------------
`
	assert.Equal(t, stringBoard, displayBoard.show(board))
}
