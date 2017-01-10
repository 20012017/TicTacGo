package ttt

import "fmt"

type DisplayBoard struct{}

const emptyBoard string = "-------------\n|[%1s]|[%1s]|[%1s]|\n-------------\n|[%1s]|[%1s]|[%1s]|\n-------------\n|[%1s]|[%1s]|[%1s]|\n-------------\n"

func (display DisplayBoard) show(board Board) string {
	c := display.formatBoard(board)
	return fmt.Sprintf(emptyBoard,
		c[0], c[1], c[2], c[3], c[4], c[5], c[6], c[7], c[8])
}

func (display DisplayBoard) formatBoard(board Board) []string {
	formattedBoard := []string{}
	for _, cell := range board.grid.cells {
		formattedBoard = append(formattedBoard, display.formatCell(cell))
	}
	return formattedBoard
}

func (display DisplayBoard) formatCell(cell string) string {
	if cell == "" {
		return "-"
	}
	return cell
}
