package validators

import (
	"errors"
	"strconv"
	"strings"
)

type Menu struct {
	start, end int
}

func NewMenuValidator(start, end int) Menu {
	return Menu{start, end}
}

func (menu Menu) Validate(input string) (bool, error) {
	input = menu.formatInput(input)
	choice, err := menu.convertToInt(input)
	if err != nil {
		return false, errors.New("Not a number")
	}
	if choice < menu.start || choice > menu.end {
		return false, errors.New("Not in menu")
	}
	return true, nil
}

func (menu Menu) ValidChoice(input string) int {
	choice, _ := menu.convertToInt(menu.formatInput(input))
	return choice
}

func (menu Menu) formatInput(input string) string {
	return strings.TrimSpace(input)
}

func (menu Menu) convertToInt(input string) (int, error) {
	return strconv.Atoi(input)
}
