package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type CliGameTest struct {
}

var cliGameTest *CliGameTest = new(CliGameTest)

func TestStartsAGame(t *testing.T) {
	playerOne := cliGameTest.newPlayerStub("X", 0, 1, 6, 5, 8)
	playerTwo := cliGameTest.newPlayerStub("O", 4, 2, 3, 7)
	displaySpy := &DisplaySpy{}
	game := NewCliGame(Game{}, playerOne, playerTwo, displaySpy)

	game.start()

	assert.True(t, displaySpy.welcomeHasBeenCalled)
	assert.True(t, displaySpy.showBoardHasBeenCalled)
	assert.True(t, displaySpy.promptHasBeenCalled)
}

func (cliGameTest CliGameTest) newPlayerStub(mark string, moves ...int) PlayerStub {
	return PlayerStub{mark, moves, 0}
}

type PlayerStub struct {
	mark        string
	moves       []int
	currentMove int
}

func (player PlayerStub) Mark() string {
	return player.mark
}

func (player PlayerStub) move(board Board) (int, error) {
	move := player.moves[player.currentMove]
	player.currentMove = player.currentMove + 1
	return move, nil
}

type DisplaySpy struct {
	welcomeHasBeenCalled   bool
	showBoardHasBeenCalled bool
	promptHasBeenCalled    bool
}

func (displaySpy DisplaySpy) write(message string) {
}

func (displaySpy *DisplaySpy) welcome() {
	displaySpy.welcomeHasBeenCalled = true
}

func (displaySpy *DisplaySpy) prompt() {
	displaySpy.promptHasBeenCalled = true
}

func (displaySpy *DisplaySpy) showBoard(board Board) {
	displaySpy.showBoardHasBeenCalled = true
}
