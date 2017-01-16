package players

import "ttt/core"

type DelayedPlayer struct {
	player core.Player
	delay  Delay
}

func NewDelayedPlayer(player core.Player, delay Delay) DelayedPlayer {
	return DelayedPlayer{player, delay}
}

func (delayedPlayer DelayedPlayer) Mark() string {
	return delayedPlayer.player.Mark()
}

func (delayedPlayer DelayedPlayer) Move(board core.Board) (int, error) {
	delayedPlayer.delay.delay(1)
	return delayedPlayer.player.Move(board)
}
