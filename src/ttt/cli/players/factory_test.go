package players

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var playerFactory Factory = NewPlayerFactory(new(InputReaderDummy))

func TestPlayerFactoryReturnsHumanPlayer(t *testing.T) {
	player := playerFactory.Create(1, "X")

	typeOfPlayer := reflect.TypeOf(player)

	assert.Equal(t, "*players.Human", fmt.Sprint(typeOfPlayer))
	assert.Equal(t, "X", player.Mark())
}

func TestPlayerFactoryReturnsDelayedPlayer(t *testing.T) {
	player := playerFactory.Create(2, "O")

	typeOfPlayer := reflect.TypeOf(player)

	assert.Equal(t, "players.DelayedPlayer", fmt.Sprint(typeOfPlayer))
	assert.Equal(t, "O", player.Mark())
}

type InputReaderDummy struct{}

func (inputReader InputReaderDummy) Read() string {
	return ""
}
