package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

func New(s string) (Matrix, error) {
	m := make([][]int, 0)
	rows := strings.Split(s, "\n")
	col := -1
	for j, v := range rows {
		v = strings.TrimSpace(v)
		rowSpl := strings.Split(v, " ")
		if col == -1 {
			col = len(rowSpl)
		} else {
			if col != len(rowSpl) {
				return nil, errors.New("")
			}
		}
		m = append(m, make([]int, 0))
		for i := 0; i < len(rowSpl); i++ {
			in, err := strconv.Atoi(rowSpl[i])
			if err != nil {
				return nil, errors.New("")
			}
			m[j] = append(m[j], in)
		}
	}
	return m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
	r := make([][]int, 0)
	for i := 0; i < len(m[0]); i++ {
		tmp := make([]int, 0)
		for j := 0; j < len(m); j++ {
			tmp = append(tmp, m[j][i])
		}
		r = append(r, tmp)
	}
	return r
}

func (m Matrix) Rows() [][]int {
	r := make([][]int, 0)
	for i := 0; i < len(m); i++ {
		tmp := make([]int, 0)
		for _, v := range m[i] {
			tmp = append(tmp, v)
		}
		r = append(r, tmp)
	}
	return r
}

func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) || col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
