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

func (menu Menu) Validate(input string) (choice int, err error) {
	for _, validation := range menu.validations() {
		choice, err = validation(menu.convertToInt(menu.formatInput(input)))
		if err != nil {
			return choice, err
		}
	}
	return choice, nil
}

func (menu Menu) validations() []func(int, error) (int, error) {
	return []func(int, error) (int, error){
		menu.validateNumber,
		menu.validateRange,
	}
}

func (menu Menu) validateRange(choice int, err error) (int, error) {
	if choice < menu.start || choice > menu.end {
		return 0, errors.New("Not in menu")
	}
	return choice, nil
}

func (menu Menu) validateNumber(choice int, err error) (int, error) {
	if err != nil {
		return 0, errors.New("Not a number")
	}
	return choice, nil
}

func (menu Menu) formatInput(input string) string {
	return strings.TrimSpace(input)
}

func (menu Menu) convertToInt(input string) (int, error) {
	return strconv.Atoi(input)
}
