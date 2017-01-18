package players

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"ttt/core/marks"
)

var playerFactory Factory = NewPlayerFactory(new(inputReaderDummy))

func TestPlayerFactoryReturnsHumanPlayer(t *testing.T) {
	player := playerFactory.Create(HUMAN, marks.X)

	typeOfPlayer := reflect.TypeOf(player)

	assert.Equal(t, "*players.Human", fmt.Sprint(typeOfPlayer))
	assert.Equal(t, marks.Mark(marks.X), player.Mark())
}

func TestPlayerFactoryReturnsDelayedPlayer(t *testing.T) {
	player := playerFactory.Create(COMPUTER, marks.O)

	typeOfPlayer := reflect.TypeOf(player)

	assert.Equal(t, "players.DelayedPlayer", fmt.Sprint(typeOfPlayer))
	assert.Equal(t, marks.O, player.Mark())
}

type inputReaderDummy struct{}

func (inputReader inputReaderDummy) Read() string {
	return ""
}
