package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

func TestMark(t *testing.T) {
	playerFake := NewFake("X", 0)

	assert.Equal(t, "X", playerFake.Mark())
}

func TestHasACurrentMoveNumber(t *testing.T) {
	playerFake := NewFake("X", 0)

	assert.Equal(t, 0, playerFake.currentMove)
}

func TestCurrentMoveIncreasesAfterAMoveIsMade(t *testing.T) {
	playerFake := NewFake("X", 0, 0)

	playerFake.Move(core.NewBoard(9))

	assert.Equal(t, 1, playerFake.currentMove)
}

func TestReturnsValidAMove(t *testing.T) {
	playerFake := NewFake("X", 0, 0)

	move, err := playerFake.Move(core.NewBoard(9))

	assert.Equal(t, 0, move)
	assert.Nil(t, err)
}

func TestReturnsTwoMoves(t *testing.T) {
	playerFake := NewFake("X", 0, 0, 1)

	playerFake.Move(core.NewBoard(9))
	move, err := playerFake.Move(core.NewBoard(9))

	assert.Equal(t, 1, move)
	assert.Nil(t, err)
}

func TestReturnsAnErrorIfMoveIsInvalid(t *testing.T) {
	playerFake := NewFake("X", 0, -1)

	_, err := playerFake.Move(core.NewBoard(9))

	assert.NotNil(t, err)
	assert.Equal(t, "Out of bounds", err.Error())
}
