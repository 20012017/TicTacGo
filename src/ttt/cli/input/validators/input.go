package validators

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Input struct {
	start, end int
}

func NewInputValidator(start, end int) Input {
	return Input{start, end}
}

func (inputValidator Input) Validate(input string) (choice int, err error) {
	for _, validation := range inputValidator.validations() {
		choice, err = validation(inputValidator.formatInput(input))
		if err != nil {
			return choice, err
		}
	}
	return choice, nil
}

func (inputValidator Input) validations() []func(int, error) (int, error) {
	return []func(int, error) (int, error){
		inputValidator.validateNumber,
		inputValidator.validateRange,
	}
}

func (inputValidator Input) validateRange(choice int, err error) (int, error) {
	if choice < inputValidator.start || choice > inputValidator.end {
		return 0, errors.New(fmt.Sprintf("Not between %d and %d",
			inputValidator.start, inputValidator.end))
	}
	return choice, nil
}

func (inputValidator Input) validateNumber(choice int, err error) (int, error) {
	if err != nil {
		return 0, errors.New("Not a number")
	}
	return choice, nil
}

func (inputValidator Input) formatInput(input string) (int, error) {
	formattedString := strings.TrimSpace(input)
	return inputValidator.convertToInt(formattedString)
}

func (inputValidator Input) convertToInt(input string) (int, error) {
	return strconv.Atoi(input)
}
