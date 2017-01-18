package validators

import (
	"errors"
	"ttt/core"
	"ttt/core/marks"
)

const takenError string = "Already taken"

const minimumMove, indexDifference, invalidMove int = 1, 1, -1

type Move struct{}

func (moveValidator Move) Validate(move string, board core.Board) (int, error) {
	inputValidator := NewInputValidator(minimumMove, board.Size())
	numberMove, err := inputValidator.Validate(move)
	if err != nil {
		return invalidMove, err
	}
	return moveValidator.validateIndex(numberMove, board)
}

func (moveValidator Move) validateIndex(move int, board core.Board) (int, error) {
	index := moveValidator.convertToIndex(move)
	if moveValidator.isInvalidCell(index, board) {
		return invalidMove, errors.New(takenError)
	}
	return index, nil
}

func (moveValidator Move) convertToIndex(move int) int {
	return move - indexDifference
}

func (moveValidator Move) isInvalidCell(index int, board core.Board) bool {
	return board.MarkAt(index) != marks.EMPTY
}
