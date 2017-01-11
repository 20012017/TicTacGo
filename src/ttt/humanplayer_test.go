package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type HumanPlayerTest struct {
	board Board
}

var humanPlayerTest HumanPlayerTest = HumanPlayerTest{NewBoard(9)}

func TestHasAMark(t *testing.T) {
	player := humanPlayerTest.newPlayer("1\n")

	assert.Equal(t, "X", player.mark)
}

func TestReturnsMark(t *testing.T) {
	player := humanPlayerTest.newPlayer("1\n")

	assert.Equal(t, "X", player.Mark())
}

func TestReadsUserInput(t *testing.T) {
	inputReader := &InputReaderDouble{move: "1\n"}
	player := HumanPlayer{"X", inputReader, MoveValidator{}}

	player.move(humanPlayerTest.board)

	assert.True(t, inputReader.called())
}

func TestReturnsErrorIfMoveIsInvalid(t *testing.T) {
	player := humanPlayerTest.newPlayer("-1\n")

	_, err := player.move(humanPlayerTest.board)

	assert.NotNil(t, err)
}

func TestReturnsMoveValid(t *testing.T) {
	player := humanPlayerTest.newPlayer("1\n")

	move, err := player.move(humanPlayerTest.board)

	assert.Equal(t, 0, move)
	assert.Nil(t, err)
}

func (humanPlayerTest HumanPlayerTest) newPlayer(move string) HumanPlayer {
	return HumanPlayer{"X", &InputReaderDouble{move: move}, MoveValidator{}}
}

type InputReaderDouble struct {
	move      string
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
	return reader.move
}
