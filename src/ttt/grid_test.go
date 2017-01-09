package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var numberedSlice []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
var grid Grid = NewPopulatedGrid(numberedSlice)

func TestGridSplit(t *testing.T) {
	splitGrid := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	assert.Equal(t, splitGrid, grid.Split(3))
}

func TestGridTranspose(t *testing.T) {
	transposedGrid := [][]string{{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"}}
	assert.Equal(t, transposedGrid, grid.Transpose(3))
}

func TestGridReverse(t *testing.T) {
	reversedGrid := []string{"9", "8", "7", "6", "5", "4", "3", "2", "1"}
	assert.Equal(t, reversedGrid, grid.Reverse())
}
