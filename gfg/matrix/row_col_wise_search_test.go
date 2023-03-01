package matrix_test

import (
	"mundhrakeshav/dsa/gfg/matrix"
	"testing"
)

func TestRowColWiseSearch(t *testing.T) {
	t.Run("Search Matrix", func(t *testing.T) {
		a := [][]int{
			{1, 	2, 		3, 		4},
			{5, 	6, 		7, 		8},
			{9, 	10,		11, 	12},
			{13, 	14,		15, 	16},
		}
		i, j := matrix.Row_col_wise_search(a, 14)
		if i != 3 || j != 1 {
			t.Error("Invalid")
		}
		
	})
	
	t.Run("Search Matrix", func(t *testing.T) {
		a := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
			{13, 14, 15, 16},
		}

		i, j := matrix.Row_col_wise_search(a, 141)
		if i != -1 || j != -1 {
			t.Error("Invalid")
		}
		
	})
}
