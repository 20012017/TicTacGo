package validators

import (
	"errors"
	"ttt/core"
	"ttt/core/marks"
)

const takenError string = "Already taken"

const minimumMove, indexDifference, invalidMove int = 1, 1, -1

type Move struct{}

func (validator Move) Validate(move string, board core.Board) (int, error) {
	inputValidator := NewInputValidator(minimumMove, board.Size())
	numberMove, err := inputValidator.Validate(move)
	if err != nil {
		return invalidMove, err
	}
	return validator.validateIndex(numberMove, board)
}

func (validator Move) validateIndex(move int, board core.Board) (int, error) {
	index := validator.convertToIndex(move)
	if validator.isInvalidCell(index, board) {
		return invalidMove, errors.New(takenError)
	}
	return index, nil
}

func (validator Move) convertToIndex(move int) int {
	return move - indexDifference
}

func (validator Move) isInvalidCell(index int, board core.Board) bool {
	return board.MarkAt(index) != marks.EMPTY
}
