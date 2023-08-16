package diamond

import (
	"bytes"
	"errors"
	"strings"
)

func Gen(inChar byte) (string, error) {
	if inChar < 'A' || inChar > 'Z' {
		return "", errors.New("")
	}
	rowLth := 2*(inChar-'A') + 1
	rows := make([]string, rowLth)
	for c := byte('A'); c <= inChar; c++ {
		row := bytes.Repeat([]byte{' '}, int(rowLth))
		colNum := inChar - c
		row[colNum] = c
		row[rowLth-colNum-1] = c
		rowNum := c - 'A'
		rows[rowNum] = string(row)
		rows[rowLth-rowNum-1] = string(row)
	}
	return strings.Join(rows, "\n"), nil
}
