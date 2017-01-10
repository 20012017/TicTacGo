package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDisplaysAnEmptyBoard(t *testing.T) {
	displayBoard, board := DisplayBoard{}, NewBoard(9)
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
	displayBoard, board := DisplayBoard{}, NewMarkedBoard([]string{"X", "", "", "", "", "", "", "", ""})
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
