package input

type InputReaderSpy struct {
	move        []string
	WasCalled   bool
	currentMove int
}

func NewInputReaderSpy(move ...string) *InputReaderSpy {
	readerSpy := InputReaderSpy{move, false, 0}
	return &readerSpy
}

func (reader *InputReaderSpy) Read() string {
	reader.WasCalled = true
	move := reader.move[reader.currentMove]
	reader.currentMove = reader.currentMove + 1
	return move
}
