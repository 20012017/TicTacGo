package cli

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShowsTheMenu(t *testing.T) {
	menuSpy := NewMenuSpy(new(CliGameSpy))
	gameRunner := NewGameRunner(menuSpy)

	gameRunner.Run()

	assert.True(t, menuSpy.showWasCalled)
}

func TestStartsTheGame(t *testing.T) {
	gameSpy := new(CliGameSpy)
	menuSpy := NewMenuSpy(gameSpy)
	gameRunner := NewGameRunner(menuSpy)

	gameRunner.Run()

	assert.True(t, gameSpy.startWasCalled)
}

func TestAsksForReply(t *testing.T) {
	menuSpy := NewMenuSpy(new(CliGameSpy))
	gameRunner := NewGameRunner(menuSpy)

	gameRunner.Run()

	assert.True(t, menuSpy.replayWasCalled)
}

type MenuSpy struct {
	game            CliGame
	showWasCalled   bool
	replayWasCalled bool
}

func NewMenuSpy(cliGame CliGame) *MenuSpy {
	return &MenuSpy{cliGame, false, false}
}

func (menuSpy *MenuSpy) show() CliGame {
	menuSpy.showWasCalled = true
	return menuSpy.game
}

func (menuSpy *MenuSpy) replay() bool {
	menuSpy.replayWasCalled = true
	return false
}

type CliGameSpy struct {
	startWasCalled bool
}

func (gameSpy *CliGameSpy) Start() {
	gameSpy.startWasCalled = true
}
