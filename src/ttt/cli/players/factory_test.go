package players

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestPlayerFactoryReturnsHumanPlayer(t *testing.T) {
	playerFactory := NewPlayerFactory(new(InputReaderDummy))

	player := playerFactory.Create(1, "X")
	typeOfPlayer := reflect.TypeOf(player)

	assert.Equal(t, "*players.Human", fmt.Sprint(typeOfPlayer))
	assert.Equal(t, "X", player.Mark())
}

func TestPlayerFactoryReturnsComputerPlayer(t *testing.T) {
	playerFactory := NewPlayerFactory(new(InputReaderDummy))

	player := playerFactory.Create(2, "X")
	typeOfPlayer := reflect.TypeOf(player)

	assert.Equal(t, "*player.Computer", fmt.Sprint(typeOfPlayer))
	assert.Equal(t, "X", player.Mark())
}

type InputReaderDummy struct{}

func (inputReader InputReaderDummy) Read() string {
	return ""
}
