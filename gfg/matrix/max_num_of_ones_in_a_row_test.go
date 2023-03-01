package matrix_test

import (
	"mundhrakeshav/dsa/gfg/matrix"
	"testing"
)

func TestMaxNumOfOnes(t *testing.T) {
	t.Run("Max no of Ones in a row", func(t *testing.T) {
		a := [][]int{
			{0, 0, 0, 1},
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
		}
		if matrix.Max_num_of_ones(a) != 2 {
			t.Error()
		}
		
	})
	
	t.Run("Max no of Ones in a row", func(t *testing.T) {
		a := [][]int{
			{0, 0, 0, 1, 1},
			{0, 1, 1, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 0, 0, 0, 1},
		}

		if matrix.Max_num_of_ones(a) != 1 {
			t.Error()
		}
	})
	
	t.Run("Max no of Ones in a row", func(t *testing.T) {
		a := [][]int{
			{0, 0, 0,},
			{0},
			{0, 0, 0, 0},
		}

		if matrix.Max_num_of_ones(a) != 0 {
			t.Error()
		}
	})
}
