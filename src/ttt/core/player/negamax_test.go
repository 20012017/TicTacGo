package player

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
	"ttt/core/marks"
)

var rules *core.Rules = new(core.Rules)
var X, O, EMPTY marks.Mark = marks.X, marks.O, marks.EMPTY

func TestScoresAWin(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]marks.Mark{X, X, X, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY})

	score := negamax.score(board, X, O)

	assert.Equal(t, 10, score)
}

func TestScoresALoss(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]marks.Mark{O, O, O, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY})

	score := negamax.score(board, X, O)

	assert.Equal(t, -10, score)
}

func TestScoresADraw(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]marks.Mark{O, X, O, X, O, X, X, O, X})

	score := negamax.score(board, X, O)

	assert.Equal(t, 0, score)
}

func TestReturnsTheOnlyFreeSpaceOnBoard(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]marks.Mark{X, O, X, O, X, O, O, X, EMPTY})

	move := negamax.move(board, O)

	assert.Equal(t, 8, move)
}

func TestReturnsTheOnlyFreeSpaceIfWin(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]marks.Mark{X, O, X, X, X, O, EMPTY, X, O})

	move := negamax.move(board, X)

	assert.Equal(t, 6, move)
}

func TestGoesForAWinOverALoss(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]marks.Mark{X, O, O, X, X, EMPTY, O, O, EMPTY})

	move := negamax.move(board, X)

	assert.Equal(t, 5, move)
}

func TestGoesForAWinIfOneIsAvailable(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]marks.Mark{X, X, EMPTY, O, O, EMPTY, EMPTY, EMPTY, EMPTY})

	move := negamax.move(board, X)

	assert.Equal(t, 2, move)
}

func TestBlocksAWin(t *testing.T) {
	negamax := NewNegamax(rules)
	board := core.NewMarkedBoard([]marks.Mark{X, EMPTY, EMPTY, O, O, EMPTY, X, EMPTY, EMPTY})

	move := negamax.move(board, X)

	assert.Equal(t, 5, move)
}

func TestReturnsTheTopCornerForAnEmptyBoard(t *testing.T) {
	negamax := NewNegamax(rules)

	move := negamax.move(core.NewBoard(9), X)

	assert.Equal(t, 0, move)
}

func TestReturnsTheCentreIfTheCornerIsTaken(t *testing.T) {
	negamax := NewNegamax(rules)

	board := core.NewMarkedBoard([]marks.Mark{X, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY})
	move := negamax.move(board, O)

	assert.Equal(t, 4, move)
}
