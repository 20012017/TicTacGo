package players

import (
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/core"
)

type Human struct {
	mark          string
	inputReader   input.InputReader
	moveValidator *validators.Move
}

func NewHuman(mark string, inputReader input.InputReader, moveValidator *validators.Move) *Human {
	return &Human{mark, inputReader, moveValidator}
}

func (player Human) Mark() string {
	return player.mark
}

func (player Human) Move(board core.Board) (int, error) {
	move := player.inputReader.Read()
	_, err := player.moveValidator.Validate(move, board)
	if err != nil {
		return 0, err
	}
	return player.moveValidator.ValidMove(move), nil
}
