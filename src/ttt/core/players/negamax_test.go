package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

func TestReturnsTheTopCornerForAnEmptyBoard(t *testing.T) {
	negamax := new(Negamax)

	move := negamax.BestMove(core.NewBoard(9))

	assert.Equal(t, 0, move)
}
