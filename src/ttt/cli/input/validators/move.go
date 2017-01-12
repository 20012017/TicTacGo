package validators

import (
	"errors"
	"strconv"
	"strings"
	"ttt/core"
)

const numberError, boundsError, takenError string = "Not a number",
	"Out of bounds",
	"Already taken"

type Move struct{}

func (validator Move) Validate(move string, board core.TTTBoard) (bool, error) {
	numberMove, err := validator.convertToInt(validator.formatInput(move))
	if validator.isNotANumber(err) {
		return false, errors.New(numberError)
	}
	for _, validation := range validator.validations() {
		valid, err := validation(numberMove, board)
		if err != nil {
			return valid, err
		}
	}
	return true, nil
}

func (validator Move) ValidMove(move string) int {
	numberMove, _ := validator.convertToInt(validator.formatInput(move))
	return validator.convertToIndex(numberMove)
}

func (validator Move) validations() []func(int, core.TTTBoard) (bool, error) {
	return []func(int, core.TTTBoard) (bool, error){
		validator.validateBounds,
		validator.validateIndex}
}

func (validator Move) validateIndex(move int, board core.TTTBoard) (bool, error) {
	if validator.isInvalidCell(move-1, board) {
		return false, errors.New(takenError)
	}
	return true, nil
}

func (validator Move) validateBounds(move int, board core.TTTBoard) (bool, error) {
	if validator.isOutOfBounds(move, board) {
		return false, errors.New(boundsError)
	}
	return true, nil
}

func (validator Move) convertToIndex(move int) int {
	return move - 1
}

func (validator Move) isInvalidCell(index int, board core.TTTBoard) bool {
	return board.MarkAt(index) != ""
}

func (validator Move) isOutOfBounds(move int, board core.TTTBoard) bool {
	return move >= board.Size() || move < 1
}

func (validator Move) isNotANumber(err error) bool {
	return err != nil
}

func (validator Move) convertToInt(move string) (int, error) {
	return strconv.Atoi(move)
}

func (validator Move) formatInput(input string) string {
	return strings.TrimSpace(input)
}
