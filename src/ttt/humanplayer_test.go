package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHumanPlayerHasAMark(t *testing.T) {
	player := HumanPlayer{"X"}
	assert.Equal(t, "X", player.mark)
}

func TestHumanPlayerReturnsMark(t *testing.T) {
	player := HumanPlayer{"X"}
	assert.Equal(t, "X", player.Mark())
}
