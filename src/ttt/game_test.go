package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCanPlayAMark(t *testing.T) {
	playerOne, playerTwo, board := newPlayerDouble("X"), newPlayerDouble("O"), NewBoard(9)
	game := NewGame(playerOne, playerTwo, board)
	game.play(0)
	assert.Equal(t, "X", game.boardAt(0))
}
