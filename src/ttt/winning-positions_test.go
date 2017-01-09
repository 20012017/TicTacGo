package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var winningPositions WinningPositions = NewWinningPositions([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"})

func TestReturnsRows(t *testing.T) {
	rows := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	assert.Equal(t, rows, winningPositions.rows)
}

func TestReturnsColumns(t *testing.T) {
	columns := [][]string{{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"}}
	assert.Equal(t, columns, winningPositions.columns)
}

func TestReturnsDiagonals(t *testing.T) {
	diagonals := [][]string{{"1", "5", "9"}, {"3", "5", "7"}}
	assert.Equal(t, diagonals, winningPositions.diagonals)
}
