package cli

import (
	"ttt/cli/input/validators"
	"ttt/core"
)

type HumanPlayer struct {
	mark          string
	inputReader   InputReader
	moveValidator *validators.Move
}

func (player HumanPlayer) Mark() string {
	return player.mark
}

func (player HumanPlayer) Move(board core.TTTBoard) (int, error) {
	move := player.inputReader.Read()
	_, err := player.moveValidator.Validate(move, board)
	if err != nil {
		return 0, err
	}
	return player.moveValidator.ValidMove(move), nil
}
