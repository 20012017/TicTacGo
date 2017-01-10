package ttt

type line struct {
	cells [][]string
}

func newLines(cells [][]string) line {
	return line{cells}
}

func (rows line) Reverse() [][]string {
	reversedRows := [][]string{}
	for _, row := range rows.cells {
		reversedRows = append(reversedRows, reverseLine(row))
	}
	return reversedRows
}

func reverseLine(row []string) []string {
	index, reverseIndex := 0, len(row)-1
	return reverseEachElement(index, reverseIndex, row)
}

func reverseEachElement(index, reverseIndex int, row []string) []string {
	reversedRow := make([]string, len(row))
	for i := 0; i < len(row); i++ {
		reversedRow[index] = row[reverseIndex]
		index++
		reverseIndex--
	}
	return reversedRow
}
