package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKnowsIfThereIsAWin(t *testing.T) {
	winningPositions := []Line{
		newLine("X", "X", "X"),
		newLine("", "", ""),
		newLine("", "", "")}
	rules := rules{}
	assert.True(t, rules.isAWin(winningPositions))
}
