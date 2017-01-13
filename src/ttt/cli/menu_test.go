package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/cli/display"
	"ttt/cli/input"
)

func PrintsAMenuOfGameChoices(t *testing.T) {
	inputReader := input.NewInputReaderSpy("1\n")
	displaySpy := new(display.Spy)
	menu := NewMenu(new(display.Spy), inputReader)

	menu.show()

	assert.True(t, displaySpy.MenuWasCalled)
}

func ReadsInput(t *testing.T) {
	inputReader := input.NewInputReaderSpy("1\n")
	menu := NewMenu(new(display.Spy), inputReader)

	menu.show()

	assert.True(t, inputReader.WasCalled)
}
