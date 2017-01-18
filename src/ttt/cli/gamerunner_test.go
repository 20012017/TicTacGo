package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShowsTheMenu(t *testing.T) {
	menuSpy := NewMenuSpy(new(cliGameSpy))
	gameRunner := NewGameRunner(menuSpy)

	gameRunner.Run()

	assert.True(t, menuSpy.showWasCalled)
}

func TestStartsTheGame(t *testing.T) {
	gameSpy := new(cliGameSpy)
	menuSpy := NewMenuSpy(gameSpy)
	gameRunner := NewGameRunner(menuSpy)

	gameRunner.Run()

	assert.True(t, gameSpy.startWasCalled)
}

func TestAsksForReply(t *testing.T) {
	menuSpy := NewMenuSpy(new(cliGameSpy))
	gameRunner := NewGameRunner(menuSpy)

	gameRunner.Run()

	assert.True(t, menuSpy.replayWasCalled)
}

type menuTestSpy struct {
	game            CliGame
	showWasCalled   bool
	replayWasCalled bool
}

func NewMenuSpy(cliGame CliGame) *menuTestSpy {
	return &menuTestSpy{cliGame, false, false}
}

func (menuSpy *menuTestSpy) show() CliGame {
	menuSpy.showWasCalled = true
	return menuSpy.game
}

func (menuSpy *menuTestSpy) replay() bool {
	menuSpy.replayWasCalled = true
	return false
}

type cliGameSpy struct {
	startWasCalled bool
}

func (gameSpy *cliGameSpy) Start() {
	gameSpy.startWasCalled = true
}
