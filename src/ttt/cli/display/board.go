package display

import (
	"fmt"
	"ttt/core"
	"ttt/core/marks"
)

const emptyCell marks.Mark = marks.Mark("-")

type Board struct{}

const emptyBoard string = "-------------\n" +
	"|[%1s]|[%1s]|[%1s]|" +
	"\n-------------\n" +
	"|[%1s]|[%1s]|[%1s]|" +
	"\n-------------\n" +
	"|[%1s]|[%1s]|[%1s]|" +
	"\n-------------\n"

func (display Board) show(board core.Board) string {
	c := display.formatBoard(board)
	return fmt.Sprintf(emptyBoard,
		c[0], c[1], c[2], c[3], c[4], c[5], c[6], c[7], c[8])
}

func (display Board) formatBoard(board core.Board) []marks.Mark {
	var formattedBoard []marks.Mark
	for _, cell := range board.Grid().Cells() {
		formattedBoard = append(formattedBoard, display.formatCell(cell))
	}
	return formattedBoard
}

func (display Board) formatCell(cell marks.Mark) marks.Mark {
	if cell == marks.EMPTY {
		return emptyCell
	}
	return cell
}
