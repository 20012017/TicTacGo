package players

import (
	"ttt/cli/input"
	"ttt/cli/input/validators"
	"ttt/core"
	"ttt/core/marks"
)

type Human struct {
	mark          marks.Mark
	inputReader   input.InputReader
	moveValidator *validators.Move
}

func NewHuman(mark marks.Mark, inputReader input.InputReader, moveValidator *validators.Move) *Human {
	return &Human{mark, inputReader, moveValidator}
}

func (player Human) Mark() marks.Mark {
	return player.mark
}

func (player Human) Move(board core.Board) (int, error) {
	move := player.inputReader.Read()
	validatedMove, err := player.moveValidator.Validate(move, board)
	if err != nil {
		return validatedMove, err
	}
	return validatedMove, nil
}
