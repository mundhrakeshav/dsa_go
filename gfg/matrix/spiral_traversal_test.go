package matrix_test

import (
	"mundhrakeshav/dsa/gfg/matrix"
	"reflect"
	"testing"
)

func TestSpiralTraversal(t *testing.T) {
	t.Run("Rotate Matrix", func(t *testing.T) {
		a := [][]int{
			{1,		2, 		3, 		4},
			{5, 	6, 		7, 		8},
			{9, 	10, 	11, 	12},
			{13,	14,		15, 	16},
		}
		z := matrix.SpiralTraversal(a)
		expected := []int{1, 2, 3, 4, 8, 12, 16, 15, 14, 13, 9, 5, 6, 7, 11, 10};
		if !reflect.DeepEqual(z, expected) {
			t.Error("Invalid")
		}
	})
}
