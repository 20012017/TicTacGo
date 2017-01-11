package ttt

import (
	"strconv"
	"strings"
)

type HumanPlayer struct {
	mark        string
	inputReader InputReader
}

func (player HumanPlayer) Mark() string {
	return player.mark
}

func (player HumanPlayer) move() int {
	move := player.inputReader.read()
	return player.convertToInt(player.trim(move)) - 1
}

func (player HumanPlayer) trim(move string) string {
	return strings.TrimSpace(move)
}

func (player HumanPlayer) convertToInt(move string) int {
	validmove, _ := strconv.Atoi(move)
	return validmove
}
