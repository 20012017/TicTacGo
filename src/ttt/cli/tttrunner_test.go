package cli

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
	assert.NotNil(t, game.Board())
}

func TestCreatesDisplay(t *testing.T) {
	display := new(TTT).createDisplay()

	assert.NotNil(t, display)
}

func TestCreatesCliGame(t *testing.T) {
	cliGame := new(TTT).CreateCliGame()

	assert.NotNil(t, cliGame)
	assert.NotNil(t, cliGame.game)
	assert.NotNil(t, cliGame.display)
}
