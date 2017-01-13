package display

import "ttt/core"

type Spy struct {
	WelcomeHasBeenCalled   bool
	ShowBoardHasBeenCalled bool
	PromptHasBeenCalled    bool
	DrawHasBeenCalled      bool
	GoodbyeHasBeenCalled   bool
	WinHasBeenCalled       bool
	ClearHasBeenCalled     bool
	WriteHasBeenCalled     bool
	WriteArgument          string
	MenuWasCalled          bool
	InvalidChoiceWasCalled bool
}

func (displaySpy *Spy) Write(message string) {
	displaySpy.WriteHasBeenCalled = true
	displaySpy.WriteArgument = message
}

func (displaySpy *Spy) Welcome() {
	displaySpy.WelcomeHasBeenCalled = true
}

func (displaySpy *Spy) Prompt() {
	displaySpy.PromptHasBeenCalled = true
}

func (displaySpy *Spy) ShowBoard(board core.Board) {
	displaySpy.ShowBoardHasBeenCalled = true
}

func (displaySpy *Spy) Draw() {
	displaySpy.DrawHasBeenCalled = true
}

func (displaySpy *Spy) Goodbye() {
	displaySpy.GoodbyeHasBeenCalled = true
}

func (displaySpy *Spy) Win(mark string) {
	displaySpy.WinHasBeenCalled = true
}

func (displaySpy *Spy) Clear() {
	displaySpy.ClearHasBeenCalled = true
}

func (displaySpy *Spy) Menu() {
	displaySpy.MenuWasCalled = true
}

func (displaySpy *Spy) InvalidChoice() {
	displaySpy.InvalidChoiceWasCalled = true
}
