package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatesAMenu(t *testing.T) {
	menu := new(TTT).CreateMenu()

	assert.NotNil(t, menu)
}

func TestCreatesDisplay(t *testing.T) {
	display := new(TTT).createDisplay()

	assert.NotNil(t, display)
}
