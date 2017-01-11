package ttt

type HumanPlayer struct {
	mark string
}

func (player HumanPlayer) Mark() string {
	return player.mark
}
