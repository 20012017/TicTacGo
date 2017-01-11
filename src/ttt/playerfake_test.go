package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMark(t *testing.T) {
	playerFake := newPlayerFake("X", 0)

	assert.Equal(t, "X", playerFake.Mark())
}

func TestHasACurrentMoveNumber(t *testing.T) {
	playerFake := newPlayerFake("X", 0)

	assert.Equal(t, 0, playerFake.currentMove)
}

func TestCurrentMoveIncreasesAfterAMoveIsMade(t *testing.T) {
	playerFake := newPlayerFake("X", 0, 0)

	playerFake.move(NewBoard(9))

	assert.Equal(t, 1, playerFake.currentMove)
}

func TestReturnsValidAMove(t *testing.T) {
	playerFake := newPlayerFake("X", 0, 0)

	move, err := playerFake.move(NewBoard(9))

	assert.Equal(t, 0, move)
	assert.Nil(t, err)
}

func TestReturnsTwoMoves(t *testing.T) {
	playerFake := newPlayerFake("X", 0, 0, 1)

	playerFake.move(NewBoard(9))
	move, err := playerFake.move(NewBoard(9))

	assert.Equal(t, 1, move)
	assert.Nil(t, err)
}
