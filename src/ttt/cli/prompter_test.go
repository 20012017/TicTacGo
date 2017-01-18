package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/cli/display"
	"ttt/core/marks"
)

type prompterTests struct{}

var gameOptions map[int][]int = GameOptions

var prompterTest prompterTests = prompterTests{}

var prompter Prompter = Prompter{gameOptions}
var X, O marks.Mark = marks.X, marks.O

func TestHumanVHumanGamePlayerOne(t *testing.T) {
	displaySpy := prompterTest.newSpy()

	prompter.Prompt(1, X, displaySpy)

	assert.True(t, displaySpy.HumanPromptHasBeenCalled)
}

func TestHumanVHumanGamePlayerTwo(t *testing.T) {
	displaySpy := prompterTest.newSpy()

	prompter.Prompt(1, O, displaySpy)

	assert.True(t, displaySpy.HumanPromptHasBeenCalled)
}

func TestHumanVComputerGamePlayerOne(t *testing.T) {
	displaySpy := prompterTest.newSpy()

	prompter.Prompt(2, X, displaySpy)

	assert.True(t, displaySpy.HumanPromptHasBeenCalled)
}

func TestHumanVComputerGamePlayerTwo(t *testing.T) {
	displaySpy := prompterTest.newSpy()

	prompter.Prompt(2, O, displaySpy)

	assert.True(t, displaySpy.ComputerPromptHasBeenCalled)
}

func TestComputerVHumanGamePlayerOne(t *testing.T) {
	displaySpy := prompterTest.newSpy()

	prompter.Prompt(3, X, displaySpy)

	assert.True(t, displaySpy.ComputerPromptHasBeenCalled)
}

func TestComputerVHumanGamePlayerTwo(t *testing.T) {
	displaySpy := prompterTest.newSpy()

	prompter.Prompt(3, O, displaySpy)

	assert.True(t, displaySpy.HumanPromptHasBeenCalled)
}

func TestComputerVComputerGamePlayerOne(t *testing.T) {
	displaySpy := prompterTest.newSpy()

	prompter.Prompt(4, X, displaySpy)

	assert.True(t, displaySpy.ComputerPromptHasBeenCalled)
}

func TestComputerVComputerGamePlayerTwo(t *testing.T) {
	displaySpy := prompterTest.newSpy()

	prompter.Prompt(4, O, displaySpy)

	assert.True(t, displaySpy.ComputerPromptHasBeenCalled)
}

func (prompterTest prompterTests) newSpy() *display.Spy {
	return &display.Spy{}
}
