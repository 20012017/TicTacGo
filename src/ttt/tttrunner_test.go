package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatesPlayers(t *testing.T) {
	playerOne, playerTwo := new(TTT).createPlayers()

	assert.NotNil(t, playerOne)
	assert.NotNil(t, playerTwo)
}

func TestCreatesGame(t *testing.T) {
	game := new(TTT).createGame()

	assert.NotNil(t, game)
	assert.NotNil(t, game.board)
	assert.NotNil(t, game.playerOne)
	assert.NotNil(t, game.playerTwo)
	assert.NotNil(t, game.rules)
}

func TestCreatesDisplay(t *testing.T) {
	display := new(TTT).createDisplay()

	assert.NotNil(t, display)
}

func TestCreatesCliGame(t *testing.T) {
	cliGame := new(TTT).createCliGame()

	assert.NotNil(t, cliGame)
	assert.NotNil(t, cliGame.game)
	assert.NotNil(t, cliGame.display)
}
