package cli

import (
	"bufio"
)

type CliInputReader struct {
	reader *bufio.Reader
}

func (input CliInputReader) read() string {
	userInput, _ := input.reader.ReadString('\n')
	return userInput
}
