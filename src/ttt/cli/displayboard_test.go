package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

func TestDisplaysAnEmptyBoard(t *testing.T) {
	displayBoard, board := DisplayBoard{}, core.NewBoard(9)
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
	displayBoard, board := DisplayBoard{}, core.NewMarkedBoard([]string{"X", "", "", "", "", "", "", "", ""})
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
