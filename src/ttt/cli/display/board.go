package display

import (
	"fmt"
	"ttt/core"
)

const emptyCell string = "-"

type Board struct{}

const emptyBoard string = "-------------\n" +
	"|[%1s]|[%1s]|[%1s]|" +
	"\n-------------\n" +
	"|[%1s]|[%1s]|[%1s]|" +
	"\n-------------\n" +
	"|[%1s]|[%1s]|[%1s]|" +
	"\n-------------\n"

func (displayBoard Board) show(board core.Board) string {
	cells := displayBoard.formatBoard(board)
	return fmt.Sprintf(emptyBoard,
		cells[0],
		cells[1],
		cells[2],
		cells[3],
		cells[4],
		cells[5],
		cells[6],
		cells[7],
		cells[8])
}

func (displayBoard Board) formatBoard(board core.Board) []string {
	formattedBoard := []string{}
	for _, cell := range board.Grid().Cells() {
		formattedBoard = append(formattedBoard, displayBoard.formatCell(cell))
	}
	return formattedBoard
}

func (displayBoard Board) formatCell(cell string) string {
	if cell == core.EmptyMark() {
		return emptyCell
	}
	return cell
}
