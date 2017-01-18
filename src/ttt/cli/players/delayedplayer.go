package players

import (
	"ttt/core"
	"ttt/core/marks"
)

type DelayedPlayer struct {
	player core.Player
	delay  Delay
}

func NewDelayedPlayer(player core.Player, delay Delay) DelayedPlayer {
	return DelayedPlayer{player, delay}
}

func (delayedPlayer DelayedPlayer) Mark() marks.Mark {
	return delayedPlayer.player.Mark()
}

func (delayedPlayer DelayedPlayer) Move(board core.Board) (int, error) {
	delayedPlayer.delay.delay(1)
	return delayedPlayer.player.Move(board)
}
