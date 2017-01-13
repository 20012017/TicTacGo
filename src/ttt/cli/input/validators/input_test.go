package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var inputValidator Input = Input{1, 4}

func TestDoesNotValidateWord(t *testing.T) {
	choice, err := inputValidator.Validate("hello\n")

	assert.Equal(t, 0, choice)
	assert.Equal(t, "Not a number", err.Error())
}

func TestDoesNotValidateAMoveLargerThanTheInput(t *testing.T) {
	choice, err := inputValidator.Validate("10\n")

	assert.Equal(t, 0, choice)
	assert.Equal(t, "Not between 1 and 4", err.Error())
}

func TestDoesNotValidateAMoveSmallerThanTheInput(t *testing.T) {
	choice, err := inputValidator.Validate("-1\n")

	assert.Equal(t, 0, choice)
	assert.Equal(t, "Not between 1 and 4", err.Error())
}

func TestValidatesAValidChoice(t *testing.T) {
	choice, err := inputValidator.Validate("1\n")

	assert.Nil(t, err)
	assert.Equal(t, 1, choice)
}
