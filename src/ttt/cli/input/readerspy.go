package input

type InputReaderSpy struct {
	move      string
	WasCalled bool
}

func NewInputReaderSpy(move string) *InputReaderSpy {
	readerSpy := InputReaderSpy{move, false}
	return &readerSpy
}

func (reader *InputReaderSpy) Read() string {
	reader.WasCalled = true
	return reader.move
}
