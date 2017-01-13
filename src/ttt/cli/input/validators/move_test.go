package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

func TestDoesNotValidateAWord(t *testing.T) {
	moveValidator := Move{}

	move, err := moveValidator.Validate("hello\n", core.NewBoard(9))

	assert.Equal(t, -1, move)
	assert.Equal(t, "Not a number", err.Error())
}

func TestDoesNotValidateAMoveLargerThanTheBoard(t *testing.T) {
	moveValidator := Move{}

	move, err := moveValidator.Validate("10\n", core.NewBoard(9))

	assert.Equal(t, -1, move)
	assert.Equal(t, "Not between 1 and 9", err.Error())
}

func TestDoesNotValidateAMoveSmallerThanTheBoard(t *testing.T) {
	moveValidator := Move{}

	move, err := moveValidator.Validate("-1\n", core.NewBoard(9))

	assert.Equal(t, -1, move)
	assert.Equal(t, "Not between 1 and 9", err.Error())
}

func TestDoesNotValidateAMoveOnATakenCell(t *testing.T) {
	moveValidator := Move{}

	markedBoard := core.NewMarkedBoard([]string{"X", "", "", "", "", "", "", "", ""})
	move, err := moveValidator.Validate("1\n", markedBoard)

	assert.Equal(t, -1, move)
	assert.Equal(t, "Already taken", err.Error())
}

func TestValidatesAValidMove(t *testing.T) {
	moveValidator := Move{}

	move, err := moveValidator.Validate("4\n", core.NewBoard(9))

	assert.Equal(t, 3, move)
	assert.Nil(t, err)
}
