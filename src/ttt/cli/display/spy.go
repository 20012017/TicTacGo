package display

import (
	"ttt/core"
	"ttt/core/marks"
)

type Spy struct {
	WelcomeHasBeenCalled        bool
	ShowBoardHasBeenCalled      bool
	HumanPromptHasBeenCalled    bool
	ComputerPromptHasBeenCalled bool
	DrawHasBeenCalled           bool
	GoodbyeHasBeenCalled        bool
	WinHasBeenCalled            bool
	ClearHasBeenCalled          bool
	WriteHasBeenCalled          bool
	WriteArgument               string
	MenuHasBeenCalled           bool
	InvalidChoiceHasBeenCalled  bool
	ReplayHasBeenCalled         bool
}

func (displaySpy *Spy) Write(message string) {
	displaySpy.WriteHasBeenCalled = true
	displaySpy.WriteArgument = message
}

func (displaySpy *Spy) Welcome() {
	displaySpy.WelcomeHasBeenCalled = true
}

func (displaySpy *Spy) HumanPrompt() {
	displaySpy.HumanPromptHasBeenCalled = true
}

func (displaySpy *Spy) ComputerPrompt() {
	displaySpy.ComputerPromptHasBeenCalled = true
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

func (displaySpy *Spy) Win(mark marks.Mark) {
	displaySpy.WinHasBeenCalled = true
}

func (displaySpy *Spy) Clear() {
	displaySpy.ClearHasBeenCalled = true
}

func (displaySpy *Spy) Menu() {
	displaySpy.MenuHasBeenCalled = true
}

func (displaySpy *Spy) InvalidChoice() {
	displaySpy.InvalidChoiceHasBeenCalled = true
}

func (displaySpy *Spy) Replay() {
	displaySpy.ReplayHasBeenCalled = true
}
