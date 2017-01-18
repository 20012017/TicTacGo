package validators

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const numberError, rangeError string = "Not a number", "Not between %d and %d"
const invalidChoice int = 0

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
		return invalidChoice, errors.New(fmt.Sprintf(rangeError,
			inputValidator.start, inputValidator.end))
	}
	return choice, nil
}

func (inputValidator Input) validateNumber(choice int, err error) (int, error) {
	if err != nil {
		return invalidChoice, errors.New(numberError)
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
