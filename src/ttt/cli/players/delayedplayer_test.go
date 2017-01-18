package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
	"ttt/core/marks"
)

func TestHasMark(t *testing.T) {
	playerSpy := newPlayerSpy(marks.X, 1)
	delayedPlayer := NewDelayedPlayer(playerSpy, new(delaySpy))

	mark := delayedPlayer.Mark()

	assert.True(t, playerSpy.markHasBeenCalled)
	assert.Equal(t, marks.X, mark)
}

func TestHasAMove(t *testing.T) {
	playerSpy := newPlayerSpy(marks.X, 1)
	delayedPlayer := NewDelayedPlayer(playerSpy, new(delaySpy))

	move, _ := delayedPlayer.Move(core.NewBoard(9))

	assert.True(t, playerSpy.moveHasBeenCalled)
	assert.Equal(t, 1, move)
}

func TestDelaysTheMove(t *testing.T) {
	delay := new(delaySpy)
	delayedPlayer := NewDelayedPlayer(newPlayerSpy(marks.X, 1), delay)

	delayedPlayer.Move(core.NewBoard(9))

	assert.True(t, delay.hasBeenCalled)
}

type delaySpy struct {
	hasBeenCalled bool
}

func (delay *delaySpy) delay(duration int) {
	delay.hasBeenCalled = true
}

type delayPlayerSpy struct {
	mark              marks.Mark
	move              int
	markHasBeenCalled bool
	moveHasBeenCalled bool
}

func newPlayerSpy(mark marks.Mark, move int) *delayPlayerSpy {
	return &delayPlayerSpy{mark, move, false, false}
}

func (playerSpy *delayPlayerSpy) Move(board core.Board) (int, error) {
	playerSpy.moveHasBeenCalled = true
	return playerSpy.move, nil
}

func (playerSpy *delayPlayerSpy) Mark() marks.Mark {
	playerSpy.markHasBeenCalled = true
	return playerSpy.mark
}
