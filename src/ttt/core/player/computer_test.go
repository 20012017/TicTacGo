package player

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
	"ttt/core/marks"
)

var negamax *Negamax = NewNegamax(new(core.Rules))

func TestHasAMark(t *testing.T) {
	computer := NewComputer(marks.X, negamax)

	assert.Equal(t, marks.Mark(marks.X), computer.Mark())
}

func TestMakesAMove(t *testing.T) {
	computer := NewComputer(marks.X, negamax)

	move, _ := computer.Move(core.NewBoard(9))

	assert.NotNil(t, move)
}

func TestDoesNotReturnAnError(t *testing.T) {
	computer := NewComputer(marks.X, negamax)

	_, err := computer.Move(core.NewBoard(9))

	assert.Nil(t, err)
}
