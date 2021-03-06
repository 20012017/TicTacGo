package players

import (
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/core"
	"ttt/core/marks"
	"ttt/core/player"
)

type Factory struct {
	inputReader input.InputReader
}

func NewPlayerFactory(inputReader input.InputReader) Factory {
	return Factory{inputReader}
}

func (factory Factory) Create(choice int, mark marks.Mark) core.Player {
	if choice == HUMAN {
		return factory.createHumanPlayer(mark)
	}
	return factory.createComputerPlayer(mark)
}

func (factory Factory) createHumanPlayer(mark marks.Mark) core.Player {
	moveValidator := new(validators.Move)
	return NewHuman(mark, factory.inputReader, moveValidator)
}

func (factory Factory) createComputerPlayer(mark marks.Mark) core.Player {
	negamax := player.NewNegamax(new(core.Rules))
	computerPlayer := player.NewComputer(mark, negamax)
	return NewDelayedPlayer(computerPlayer, TimeDelay{})
}
