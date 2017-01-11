package ttt

import (
	"bufio"
)

type InputReader struct {
	reader *bufio.Reader
}

func (input InputReader) read() string {
	userInput, _ := input.reader.ReadString('\n')
	return userInput
}
