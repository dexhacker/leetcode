package find_valid_matrix_given_row_and_column_sums

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	n, m := len(rowSum), len(colSum)
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			matrix[i][j] = minimum(rowSum[i], colSum[j])
			rowSum[i] -= matrix[i][j]
			colSum[j] -= matrix[i][j]
		}
	}

	return matrix
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
