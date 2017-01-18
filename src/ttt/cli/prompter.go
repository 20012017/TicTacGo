package cli

import (
	"ttt/cli/players"
	"ttt/core"
)

type Prompter struct {
	gameOptions map[int][]int
}

func NewPrompter(gameOptions map[int][]int) Prompter {
	return Prompter{gameOptions}
}

func (prompter Prompter) Prompt(gameType int, currentMark string, display DisplayWriter) {
	gamePlayers := prompter.gameOptions[gameType]
	currentPlayer := prompter.currentPlayer(currentMark)
	player := gamePlayers[currentPlayer]

	if player == players.HUMAN {
		display.HumanPrompt()
	} else {
		display.ComputerPrompt()
	}
}

func (prompter Prompter) currentPlayer(mark string) int {
	if mark == core.MarkX() {
		return 0
	}
	return 1
}
