package cli

import "ttt/core"

type HumanPlayer struct {
	mark          string
	inputReader   InputReader
	moveValidator MoveValidator
}

func (player HumanPlayer) Mark() string {
	return player.mark
}

func (player HumanPlayer) Move(board core.TTTBoard) (int, error) {
	move := player.inputReader.Read()
	_, err := player.moveValidator.validate(move, board)
	if err != nil {
		return 0, err
	}
	return player.moveValidator.validMove(move), nil
}
