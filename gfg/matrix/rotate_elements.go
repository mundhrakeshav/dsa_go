package matrix

// m x n matrix
func RotateMatrix(matrix [][]int, m, n int) [][]int {
	row, col := 0, 0
	prev, curr := 0, 0

	// Loop in remaining rows and columns
	for row < m && col < n {
		//
		if row+1 == m || col+1 == n {
			break
		}
		//

		prev = matrix[row+1][col]

		// Shift first row
		for i := col; i < n; i++ {
			curr = matrix[row][i]
			matrix[row][i] = prev
			prev = curr
		}
		row++
		//

		// Shift last column
		for i := row; i < m; i++ {
			curr = matrix[i][n-1]
			matrix[i][n-1] = prev
			prev = curr
		}
		n--
		//

		// Shift last row
		if row < m {
			for i := n - 1; i >= col; i-- {
				curr = matrix[m-1][i]
				matrix[m-1][i] = prev
				prev = curr
			}
			m--
		}
		//

		// Shift last column
		if col < n {
			for i := m - 1; i >= row; i-- {
				curr = matrix[i][col]
				matrix[i][col] = prev
				prev = curr
			}
			col++
		}
		//

	}
	return matrix
}
