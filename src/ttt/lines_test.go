package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinesReverse(t *testing.T) {
	numberedLines := [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	lines := newLines(numberedLines)
	reversedLines := [][]string{{"3", "2", "1"}, {"6", "5", "4"}, {"9", "8", "7"}}
	assert.Equal(t, reversedLines, lines.Reverse())
}
