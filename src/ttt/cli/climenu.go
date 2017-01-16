package cli

type CliMenu interface {
	show() CliGame

	replay() bool
}
