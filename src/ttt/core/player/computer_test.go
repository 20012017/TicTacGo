package player

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

var negamax *Negamax = NewNegamax(new(core.Rules))

func TestHasAMark(t *testing.T) {
	computer := NewComputer("X", negamax)

	assert.Equal(t, "X", computer.Mark())
}

func TestMakesAMove(t *testing.T) {
	computer := NewComputer("X", negamax)

	move, _ := computer.Move(core.NewBoard(9))

	assert.NotNil(t, move)
}

func TestDoesNotReturnAnError(t *testing.T) {
	computer := NewComputer("X", negamax)

	_, err := computer.Move(core.NewBoard(9))

	assert.Nil(t, err)
}
