package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasAnXMark(t *testing.T) {
	marks := NewMarks("X", "O", "")

	assert.Equal(t, "X", marks.X)
}

func TestHasAnOMark(t *testing.T) {
	marks := NewMarks("X", "O", "")

	assert.Equal(t, "O", marks.O)
}

func TestHasAnEmptyMark(t *testing.T) {
	marks := NewMarks("X", "O", "")

	assert.Equal(t, "", marks.Empty)
}
