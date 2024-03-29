package matrix

//		→→→→→→→→→→→ n ←←←←←←←←←←←←←←
// 		↓ {1,	2, 		3, 		4},
// 	 m	↓ {5, 	6, 		7, 		8},
// 		↑ {9, 	10, 	11, 	12},
// 		↑ {13,	14,		15, 	16},

func SpiralTraversal(arr [][]int) (res []int) {
	row, col := 0, 0
	m, n := len(arr), len(arr[0])
	res = make([]int, 0, m*n)

	for row < m && col < n {
		//
		if row+1 == m || col+1 == n {
			break
		}
		//
		for i := col; i < n; i++ {
			res = append(res, arr[row][i])
		}
		row++

		for i := row; i < m; i++ {
			res = append(res, arr[i][n-1])
		}

		n--

		if row < m {
			for i := n -1; i >= col; i-- {
				res = append(res, arr[m-1][i])
			}
			m--
		}

		if col < n {
			for i := m -1; i >= row; i-- {
				res = append(res, arr[i][col])
			}
			col++
		}
	}

	return res
}
