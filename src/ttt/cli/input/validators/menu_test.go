package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoesNotValidateWord(t *testing.T) {
	menuValidator := Menu{1, 4}

	choice, err := menuValidator.Validate("hello\n")

	assert.Equal(t, 0, choice)
	assert.Equal(t, "Not a number", err.Error())
}

func TestDoesNotValidateAMoveLargerThanTheMenu(t *testing.T) {
	menuValidator := Menu{1, 4}

	choice, err := menuValidator.Validate("10\n")

	assert.Equal(t, 0, choice)
	assert.Equal(t, "Not in menu", err.Error())
}

func TestDoesNotValidateAMoveSmallerThanTheMenu(t *testing.T) {
	menuValidator := Menu{1, 4}

	choice, err := menuValidator.Validate("-1\n")

	assert.Equal(t, 0, choice)
	assert.Equal(t, "Not in menu", err.Error())
}

func TestValidatesAValidChoice(t *testing.T) {
	menuValidator := Menu{1, 4}

	choice, err := menuValidator.Validate("1\n")

	assert.Nil(t, err)
	assert.Equal(t, 1, choice)
}
