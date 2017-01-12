package cli

type InputReader interface {
	read() string
}
