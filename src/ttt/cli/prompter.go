package cli

import (
	"ttt/cli/players"
	"ttt/core/marks"
)

type Prompter struct {
	gameOptions map[int][]int
}

func NewPrompter(gameOptions map[int][]int) Prompter {
	return Prompter{gameOptions}
}

func (prompter Prompter) Prompt(gameType int, currentMark string, display DisplayWriter) {
	player := prompter.currentPlayerType(gameType, currentMark)
	if player == players.HUMAN {
		display.HumanPrompt()
	} else {
		display.ComputerPrompt()
	}
}

func (prompter Prompter) currentPlayerType(gameType int, currentMark string) int {
	players := prompter.gameOptions[gameType]
	currentPlayer := prompter.currentPlayer(currentMark)
	return players[currentPlayer]
}

func (prompter Prompter) currentPlayer(mark string) int {
	if mark == marks.X {
		return 0
	}
	return 1
}
