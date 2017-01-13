package player

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

var rules *core.Rules = new(core.Rules)

func TestScoresAWin(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]string{"X", "X", "X", "", "", "", "", "", ""})

	score := negamax.score(board, 1, "X", "O")

	assert.Equal(t, 10, score)
}

func TestScoresALoss(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]string{"O", "O", "O", "", "", "", "", "", ""})

	score := negamax.score(board, 1, "X", "O")

	assert.Equal(t, -10, score)
}

func TestScoresADraw(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]string{"O", "X", "O", "X", "O", "X", "X", "O", "X"})

	score := negamax.score(board, 1, "X", "O")

	assert.Equal(t, 0, score)
}

func TestReturnsTheOnlyFreeSpaceOnBoard(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]string{"X", "O", "X", "O", "X", "O", "O", "X", ""})

	move := negamax.move(board, "O")

	assert.Equal(t, 8, move)
}

func TestReturnsTheOnlyFreeSpaceIfWin(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]string{"X", "O", "X", "X", "X", "O", "", "X", "O"})

	move := negamax.move(board, "X")

	assert.Equal(t, 6, move)
}

func TestGoesForAWinOverALoss(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]string{"X", "O", "O", "X", "X", "", "O", "O", ""})

	move := negamax.move(board, "X")

	assert.Equal(t, 5, move)
}

func TestGoesForAWinIfOneIsAvailable(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]string{"X", "X", "", "O", "O", "", "", "", ""})

	move := negamax.move(board, "X")

	assert.Equal(t, 2, move)
}

func TestBlocksAWin(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]string{"X", "", "", "O", "O", "", "X", "", ""})

	move := negamax.move(board, "X")

	assert.Equal(t, 5, move)
}

func TestReturnsTheTopCornerForAnEmptyBoard(t *testing.T) {
	negamax := NewNegamax(rules)

	move := negamax.move(core.NewBoard(9), "X")

	assert.Equal(t, 0, move)
}

func TestReturnsTheCentreIfTheCornerIsTaken(t *testing.T) {
	negamax := NewNegamax(rules)

	board := core.NewMarkedBoard([]string{"X", "", "", "", "", "", "", "", ""})
	move := negamax.move(board, "O")

	assert.Equal(t, 4, move)
}
