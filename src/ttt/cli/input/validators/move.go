package validators

import (
	"errors"
	"ttt/core"
)

const takenError string = "Already taken"

const minimumMove, indexDifference, sizeDifference int = 1, 1, 1

type Move struct{}

func (validator Move) Validate(move string, board core.Board) (int, error) {
	inputValidator := NewInputValidator(minimumMove, board.Size())
	numberMove, err := inputValidator.Validate(move)
	if err != nil {
		return -1, err
	}
	return validator.validateIndex(numberMove, board)
}

func (validator Move) validateIndex(move int, board core.Board) (int, error) {
	index := validator.convertToIndex(move)
	if validator.isInvalidCell(index, board) {
		return -1, errors.New(takenError)
	}
	return index, nil
}

func (validator Move) convertToIndex(move int) int {
	return move - 1
}

func (validator Move) isInvalidCell(index int, board core.Board) bool {
	return board.MarkAt(index) != ""
}
