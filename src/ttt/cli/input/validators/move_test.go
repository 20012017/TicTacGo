package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
	"ttt/core/marks"
)

var moveValidator Move = Move{}
var markedBoard = core.NewMarkedBoard([]marks.Mark{"X", "", "", "", "", "", "", "", ""})

func TestDoesNotValidateAWord(t *testing.T) {
	move, err := moveValidator.Validate("hello\n", core.NewBoard(9))

	assert.Equal(t, -1, move)
	assert.Equal(t, "Not a number", err.Error())
}

func TestDoesNotValidateAMoveLargerThanTheBoard(t *testing.T) {
	move, err := moveValidator.Validate("10\n", core.NewBoard(9))

	assert.Equal(t, -1, move)
	assert.Equal(t, "Not between 1 and 9", err.Error())
}

func TestDoesNotValidateAMoveSmallerThanTheBoard(t *testing.T) {
	move, err := moveValidator.Validate("-1\n", core.NewBoard(9))

	assert.Equal(t, -1, move)
	assert.Equal(t, "Not between 1 and 9", err.Error())
}

func TestDoesNotValidateAMoveOnATakenCell(t *testing.T) {
	move, err := moveValidator.Validate("1\n", markedBoard)

	assert.Equal(t, -1, move)
	assert.Equal(t, "Already taken", err.Error())
}

func TestValidatesAValidMove(t *testing.T) {
	move, err := moveValidator.Validate("4\n", core.NewBoard(9))

	assert.Equal(t, 3, move)
	assert.Nil(t, err)
}
