package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/cli/input/validators"
	"ttt/core"
)

type HumanTest struct {
	board core.TTTBoard
}

var humanPlayerTest HumanTest = HumanTest{core.NewBoard(9)}

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
	player := Human{"X", inputReader, new(validators.Move)}

	player.Move(humanPlayerTest.board)

	assert.True(t, inputReader.called())
}

func TestReturnsErrorIfMoveIsInvalid(t *testing.T) {
	player := humanPlayerTest.newPlayer("-1\n")

	_, err := player.Move(humanPlayerTest.board)

	assert.NotNil(t, err)
}

func TestReturnsMoveValid(t *testing.T) {
	player := humanPlayerTest.newPlayer("1\n")

	move, err := player.Move(humanPlayerTest.board)

	assert.Equal(t, 0, move)
	assert.Nil(t, err)
}

func (humanPlayerTest HumanTest) newPlayer(move string) Human {
	return Human{"X", &InputReaderDouble{move: move}, new(validators.Move)}
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

func (reader *InputReaderDouble) Read() string {
	reader.setCalled(true)
	return reader.move
}
