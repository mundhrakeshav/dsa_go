package matrix_test

import (
	"fmt"
	"mundhrakeshav/dsa/gfg/matrix"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	t.Run("Rotate Matrix", func(t *testing.T) {
		a := [][]int{
			{1,		2, 		3, 		4},
			{5, 	6, 		7, 		8},
			{9, 	10, 	11, 	12},
			{13,	14,		15, 	16},
		}
		z := matrix.RotateMatrix(a, 4, 4)
		fmt.Println(z)
	})
}
