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

type MenuSpy struct {
	game          CliGame
	showWasCalled bool
}

func NewMenuSpy(cliGame CliGame) *MenuSpy {
	return &MenuSpy{cliGame, false}
}

func (menuSpy *MenuSpy) show() CliGame {
	menuSpy.showWasCalled = true
	return menuSpy.game
}

type CliGameSpy struct {
	startWasCalled bool
}

func (gameSpy *CliGameSpy) Start() {
	gameSpy.startWasCalled = true
}
