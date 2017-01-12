package cli

import (
	"errors"
	"strconv"
	"strings"
	"ttt/core"
)

const numberError, boundsError, takenError string = "Not a number",
	"Out of bounds",
	"Already taken"

type MoveValidator struct {
}

func (validator MoveValidator) validate(move string, board core.TTTBoard) (bool, error) {
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

func (validator MoveValidator) validMove(move string) int {
	numberMove, _ := validator.convertToInt(validator.formatInput(move))
	return validator.convertToIndex(numberMove)
}

func (validator MoveValidator) validations() []func(int, core.TTTBoard) (bool, error) {
	return []func(int, core.TTTBoard) (bool, error){
		validator.validateBounds,
		validator.validateIndex}
}

func (validator MoveValidator) validateIndex(move int, board core.TTTBoard) (bool, error) {
	if validator.isInvalidCell(move-1, board) {
		return false, errors.New(takenError)
	}
	return true, nil
}

func (validator MoveValidator) validateBounds(move int, board core.TTTBoard) (bool, error) {
	if validator.isOutOfBounds(move, board) {
		return false, errors.New(boundsError)
	}
	return true, nil
}

func (validator MoveValidator) convertToIndex(move int) int {
	return move - 1
}

func (validator MoveValidator) isInvalidCell(index int, board core.TTTBoard) bool {
	return board.MarkAt(index) != ""
}

func (validator MoveValidator) isOutOfBounds(move int, board core.TTTBoard) bool {
	return move >= board.Size() || move < 1
}

func (validator MoveValidator) isNotANumber(err error) bool {
	return err != nil
}

func (validator MoveValidator) convertToInt(move string) (int, error) {
	return strconv.Atoi(move)
}

func (validator MoveValidator) formatInput(input string) string {
	return strings.TrimSpace(input)
}
