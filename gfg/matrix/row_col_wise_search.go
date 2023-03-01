package matrix

// {1,		2, 		3, 		4},
// {5, 		6, 		7, 		8},
// {9, 		10, 	11, 	12},
// {13,		14,		15, 	16},

func row_col_wise_search(arr [][]int, m, n, key int) (int, int) {
	i, j := 0, n-1 // Top Right
	for i < m || j > n {
		if arr[i][j] > key {
			j--
		} else if arr[i][j] < key {
			i++
		} else {
			return i, j
		}
	}
	return -1, -1
}

func Row_col_wise_search(arr [][]int, key int) (int, int) {
	return row_col_wise_search(arr, len(arr), len(arr[0]), key)
}
