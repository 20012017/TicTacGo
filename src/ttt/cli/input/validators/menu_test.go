package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoesNotValidateWord(t *testing.T) {
	menuValidator := Menu{1, 4}

	valid, err := menuValidator.Validate("hello\n")

	assert.False(t, valid)
	assert.Equal(t, "Not a number", err.Error())
}

func TestDoesNotValidateAMoveLargerThanTheMenu(t *testing.T) {
	menuValidator := Menu{1, 4}

	valid, err := menuValidator.Validate("10\n")

	assert.False(t, valid)
	assert.Equal(t, "Not in menu", err.Error())
}

func TestDoesNotValidateAMoveSmallerThanTheMenu(t *testing.T) {
	menuValidator := Menu{1, 4}

	valid, err := menuValidator.Validate("-1\n")

	assert.False(t, valid)
	assert.Equal(t, "Not in menu", err.Error())
}

func TestValidatesAValidChoice(t *testing.T) {
	menuValidator := Menu{1, 4}

	valid, err := menuValidator.Validate("1\n")

	assert.True(t, valid)
	assert.Nil(t, err)
}

func TestReturnsAValidChoice(t *testing.T) {
	menuValidator := Menu{1, 4}

	choice := menuValidator.ValidChoice("1\n")

	assert.Equal(t, 1, choice)
}
