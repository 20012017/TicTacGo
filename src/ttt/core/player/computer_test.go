package player

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
	"ttt/core/marks"
)

var computer *Computer = NewComputer(marks.X, NewNegamax(new(core.Rules)))

func TestHasAMark(t *testing.T) {
	assert.Equal(t, marks.Mark(marks.X), computer.Mark())
}

func TestMakesAMove(t *testing.T) {
	move, _ := computer.Move(core.NewBoard(9))

	assert.NotNil(t, move)
}

func TestDoesNotReturnAnError(t *testing.T) {
	_, err := computer.Move(core.NewBoard(9))

	assert.Nil(t, err)
}
