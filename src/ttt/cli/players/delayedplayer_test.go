package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

type DelayPlayerTest struct{}

func TestHasMark(t *testing.T) {
	playerSpy := NewPlayerSpy("X", 1)
	delayedPlayer := NewDelayedPlayer(playerSpy, &DelaySpy{})

	mark := delayedPlayer.Mark()

	assert.True(t, playerSpy.markHasBeenCalled)
	assert.Equal(t, "X", mark)
}

func TestHasAMove(t *testing.T) {
	playerSpy := NewPlayerSpy("X", 1)
	delayedPlayer := NewDelayedPlayer(playerSpy, &DelaySpy{})

	move, _ := delayedPlayer.Move(core.NewBoard(9))

	assert.True(t, playerSpy.moveHasBeenCalled)
	assert.Equal(t, 1, move)
}

func TestDelaysTheMove(t *testing.T) {
	delaySpy := DelaySpy{}
	delayedPlayer := NewDelayedPlayer(NewPlayerSpy("X", 1), &delaySpy)

	delayedPlayer.Move(core.NewBoard(9))

	assert.True(t, delaySpy.hasBeenCalled)
}

type DelaySpy struct {
	hasBeenCalled bool
}

func (delaySpy *DelaySpy) delay(duration int) {
	delaySpy.hasBeenCalled = true
}

type PlayerSpy struct {
	mark              string
	move              int
	markHasBeenCalled bool
	moveHasBeenCalled bool
}

func NewPlayerSpy(mark string, move int) *PlayerSpy {
	return &PlayerSpy{mark, move, false, false}
}

func (playerSpy *PlayerSpy) Move(board core.Board) (int, error) {
	playerSpy.moveHasBeenCalled = true
	return playerSpy.move, nil
}

func (playerSpy *PlayerSpy) Mark() string {
	playerSpy.markHasBeenCalled = true
	return playerSpy.mark
}
