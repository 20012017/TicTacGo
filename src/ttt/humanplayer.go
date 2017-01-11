package ttt

type HumanPlayer struct {
	mark          string
	inputReader   InputReader
	moveValidator MoveValidator
}

func (player HumanPlayer) Mark() string {
	return player.mark
}

func (player HumanPlayer) move(board Board) (int, error) {
	move := player.inputReader.read()
	_, err := player.moveValidator.validate(move, board)
	if err != nil {
		return 0, err
	}
	return player.moveValidator.validMove(move), nil
}
