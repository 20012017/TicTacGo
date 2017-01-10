package ttt

type rules struct{}

func (rule rules) isAWin(lines []Line) bool {
	for _, line := range lines {
		if line.all(matches("X")) || line.all(matches("O")) {
			return true
		}
	}
	return false
}

func matches(mark string) func(string) bool {
	return func(str string) bool {
		return str == mark
	}
}
