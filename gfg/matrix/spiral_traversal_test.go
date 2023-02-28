package matrix_test

import (
	"fmt"
	"mundhrakeshav/dsa/gfg/matrix"
	"testing"
)

func TestSpiralTraversal(t *testing.T) {
	t.Run("Rotate Matrix", func(t *testing.T) {
		a := [][]int{
			{1,		2, 		3, 		4},
			{5, 	6, 		7, 		8},
			{9, 	10, 	11, 	12},
			{13,	14,		15, 	16},
			{17,	18,		19, 	20},
		}
		z := matrix.SpiralTraversal(a)
		fmt.Println(z)
	})
}
