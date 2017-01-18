package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/core"
	"ttt/core/marks"
)

var X marks.Mark = marks.X

type HumanTest struct {
	board core.Board
}

var humanPlayerTest HumanTest = HumanTest{core.NewBoard(9)}

func TestHasAMark(t *testing.T) {
	player := humanPlayerTest.newPlayer("1\n")

	assert.Equal(t, X, player.mark)
}

func TestReturnsMark(t *testing.T) {
	player := humanPlayerTest.newPlayer("1\n")

	assert.Equal(t, X, player.Mark())
}

func TestReadsUserInput(t *testing.T) {
	inputReader := input.NewInputReaderSpy("1\n")
	player := Human{X, inputReader, new(validators.Move)}

	player.Move(humanPlayerTest.board)

	assert.True(t, inputReader.WasCalled)
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

func (humanTest HumanTest) newPlayer(move string) Human {
	return Human{X, input.NewInputReaderSpy(move), new(validators.Move)}
}
