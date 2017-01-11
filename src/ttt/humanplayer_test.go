package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHumanPlayerHasAMark(t *testing.T) {
	player := HumanPlayer{"X", &InputReaderDouble{}}

	assert.Equal(t, "X", player.mark)
}

func TestHumanPlayerReturnsMark(t *testing.T) {
	player := HumanPlayer{"X", &InputReaderDouble{}}

	assert.Equal(t, "X", player.Mark())
}

func TestHumanPlayerReadsInput(t *testing.T) {
	inputReader := &InputReaderDouble{}
	player := HumanPlayer{"X", inputReader}

	player.move()

	assert.True(t, inputReader.called())
}

func TestHumanPlayerConvertsStringToIndex(t *testing.T) {
	player := HumanPlayer{"X", &InputReaderDouble{}}

	move := player.move()

	assert.Equal(t, 0, move)
}

type InputReaderDouble struct {
	wasCalled bool
}

func (reader *InputReaderDouble) setCalled(called bool) {
	reader.wasCalled = called
}

func (reader *InputReaderDouble) called() bool {
	return reader.wasCalled
}

func (reader *InputReaderDouble) read() string {
	reader.setCalled(true)
	return "1\n"
}
