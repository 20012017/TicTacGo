package input

import "bufio"

type Reader struct {
	ioReader *bufio.Reader
}

func NewReader(ioReader *bufio.Reader) Reader {
	return Reader{ioReader}
}

func (reader Reader) Read() string {
	userInput, _ := reader.ioReader.ReadString('\n')
	return userInput
}
