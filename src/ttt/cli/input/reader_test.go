package input

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadsInput(t *testing.T) {
	buffer := (new(bytes.Buffer))
	buffer.WriteString("hello\n")
	ioReader := bufio.NewReader(buffer)
	inputReader := Reader{ioReader}

	assert.Equal(t, "hello\n", inputReader.Read())
}
