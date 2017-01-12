package player

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

var rules *core.Rules = new(core.Rules)

func TestScoresAWin(t *testing.T) {
	computer := NewComputer("O", "X", rules)
	board := core.NewMarkedBoard([]string{"O", "O", "O", "", "", "", "", "", ""})

	score := computer.score(board, 1)

	assert.Equal(t, 10, score)
}

func TestScoresALoss(t *testing.T) {
	computer := NewComputer("X", "O", rules)
	board := core.NewMarkedBoard([]string{"O", "O", "O", "", "", "", "", "", ""})

	score := computer.score(board, 1)

	assert.Equal(t, -10, score)
}

func TestScoresADraw(t *testing.T) {
	computer := NewComputer("X", "O", rules)
	board := core.NewMarkedBoard([]string{"O", "X", "O", "X", "O", "X", "X", "O", "X"})

	score := computer.score(board, 1)

	assert.Equal(t, 0, score)
}

func TestReturnsTheOnlyFreeSpaceOnBoard(t *testing.T) {
	computer := NewComputer("O", "X", rules)
	board := core.NewMarkedBoard([]string{"X", "O", "X", "O", "X", "O", "O", "X", ""})

	move, _ := computer.Move(board)

	assert.Equal(t, 8, move)
}

func TestReturnsTheOnlyFreeSpaceIfWin(t *testing.T) {
	computer := NewComputer("X", "O", rules)
	board := core.NewMarkedBoard([]string{"X", "O", "X", "X", "X", "O", "", "X", "O"})

	move, _ := computer.Move(board)

	assert.Equal(t, 6, move)
}

func TestGoesForAWinOverALoss(t *testing.T) {
	computer := NewComputer("X", "O", rules)
	board := core.NewMarkedBoard([]string{"X", "O", "O", "X", "X", "", "O", "O", ""})

	move, _ := computer.Move(board)

	assert.Equal(t, 5, move)
}

func TestGoesForAWinIfOneIsAvailable(t *testing.T) {
	computer := NewComputer("X", "O", rules)
	board := core.NewMarkedBoard([]string{"X", "X", "", "O", "O", "", "", "", ""})

	move, _ := computer.Move(board)

	assert.Equal(t, 2, move)
}

func TestBlocksAWin(t *testing.T) {
	computer := NewComputer("X", "O", rules)
	board := core.NewMarkedBoard([]string{"X", "", "", "O", "O", "", "X", "", ""})

	move, _ := computer.Move(board)

	assert.Equal(t, 5, move)
}

func TestReturnsTheTopCornerForAnEmptyBoard(t *testing.T) {
	computer := NewComputer("O", "X", rules)

	move, _ := computer.Move(core.NewBoard(9))

	assert.Equal(t, 0, move)
}

func TestReturnsTheCentreIfTheCornerIsTaken(t *testing.T) {
	computer := NewComputer("O", "X", rules)

	board := core.NewMarkedBoard([]string{"X", "", "", "", "", "", "", "", ""})
	move, _ := computer.Move(board)

	assert.Equal(t, 4, move)
}
