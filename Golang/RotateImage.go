package main

func rotate(matrix [][]int) {
	n := len(matrix) - 1
	_matrix := make([][]int, n + 1)
	rows := make([]int, n + 1)
	for j := 0; j <= n; j++ {
		for i := n; i >= 0; i-- {
			rows[n - i] =  matrix[i][j]
		}
		_rows := make([]int, n + 1)
		copy(_rows, rows)
		_matrix[j] = _rows
	}
	copy(matrix, _matrix)
}
