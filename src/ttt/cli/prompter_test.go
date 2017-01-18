package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/cli/display"
	"ttt/cli/players"
	"ttt/core"
)

type PrompterTest struct{}

var gameOptions map[int][]int = map[int][]int{
	1: []int{players.HUMAN, players.HUMAN},
	2: []int{players.HUMAN, players.COMPUTER},
	3: []int{players.COMPUTER, players.HUMAN},
	4: []int{players.COMPUTER, players.COMPUTER},
}

var prompterTest PrompterTest = PrompterTest{}

var prompter Prompter = Prompter{gameOptions}
var X, O string = core.MarkX(), core.MarkO()

func (prompterTest PrompterTest) newSpy() *display.Spy {
	return &display.Spy{}
}

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
