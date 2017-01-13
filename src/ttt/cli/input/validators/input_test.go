package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoesNotValidateWord(t *testing.T) {
	inputValidator := Input{1, 4}

	choice, err := inputValidator.Validate("hello\n")

	assert.Equal(t, 0, choice)
	assert.Equal(t, "Not a number", err.Error())
}

func TestDoesNotValidateAMoveLargerThanTheInput(t *testing.T) {
	inputValidator := Input{1, 4}

	choice, err := inputValidator.Validate("10\n")

	assert.Equal(t, 0, choice)
	assert.Equal(t, "Not in menu", err.Error())
}

func TestDoesNotValidateAMoveSmallerThanTheInput(t *testing.T) {
	inputValidator := Input{1, 4}

	choice, err := inputValidator.Validate("-1\n")

	assert.Equal(t, 0, choice)
	assert.Equal(t, "Not in menu", err.Error())
}

func TestValidatesAValidChoice(t *testing.T) {
	inputValidator := Input{1, 4}

	choice, err := inputValidator.Validate("1\n")

	assert.Nil(t, err)
	assert.Equal(t, 1, choice)
}
