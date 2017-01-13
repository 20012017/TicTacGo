package players

import (
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/core"
	"ttt/core/player"
)

type Factory struct {
	inputReader input.InputReader
}

func NewPlayerFactory(inputReader input.InputReader) Factory {
	return Factory{inputReader}
}

func (factory Factory) Create(choice int, mark string) core.Player {
	if choice == 1 {
		moveValidator := new(validators.Move)
		return NewHuman(mark, factory.inputReader, moveValidator)
	}
	negamax := player.NewNegamax(new(core.Rules))
	return player.NewComputer(mark, negamax)
}
