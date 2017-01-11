package ttt

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadsInput(t *testing.T) {
	buffer := (new(bytes.Buffer))
	buffer.WriteString("hello\n")
	reader := bufio.NewReader(buffer)
	inputReader := CliInputReader{reader}

	assert.Equal(t, "hello\n", inputReader.read())
}
