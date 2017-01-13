package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHasAnXMark(t *testing.T) {
	assert.Equal(t, "X", MarkX())
}

func TestHasAnOMark(t *testing.T) {
	assert.Equal(t, "O", MarkO())
}

func TestHasAnEmptyMark(t *testing.T) {
	assert.Equal(t, "", EmptyMark())
}
