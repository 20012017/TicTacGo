package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

func TestDoesNotValidateAWord(t *testing.T) {
	moveValidator := Move{}

	valid, err := moveValidator.Validate("hello\n", core.NewBoard(9))

	assert.False(t, valid)
	assert.Equal(t, "Not a number", err.Error())
}

func TestDoesNotValidateAMoveLargerThanTheBoard(t *testing.T) {
	moveValidator := Move{}

	valid, err := moveValidator.Validate("10\n", core.NewBoard(9))

	assert.False(t, valid)
	assert.Equal(t, "Out of bounds", err.Error())
}

func TestDoesNotValidateAMoveSmallerThanTheBoard(t *testing.T) {
	moveValidator := Move{}

	valid, err := moveValidator.Validate("-1\n", core.NewBoard(9))

	assert.False(t, valid)
	assert.Equal(t, "Out of bounds", err.Error())
}

func TestDoesNotValidateAMoveOnATakenCell(t *testing.T) {
	moveValidator := Move{}

	markedBoard := core.NewMarkedBoard([]string{"X", "", "", "", "", "", "", "", ""})
	valid, err := moveValidator.Validate("1\n", markedBoard)

	assert.False(t, valid)
	assert.Equal(t, "Already taken", err.Error())
}

func TestValidatesAValidMove(t *testing.T) {
	moveValidator := Move{}

	valid, err := moveValidator.Validate("1\n", core.NewBoard(9))

	assert.True(t, valid)
	assert.Nil(t, err)
}

func TestReturnsValidMove(t *testing.T) {
	moveValidator := Move{}

	move := moveValidator.ValidMove("1\n")

	assert.Equal(t, 0, move)
}
