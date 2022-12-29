package arraysgfg_test

import (
	arraysgfg "mundhrakeshav/dsa/arrays_gfg"
	"testing"
)

func TestMissingNumberArray(t *testing.T) {

	t.Run("TestMissingNumberArray_XOR", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, /*5,*/ 6, 7, 8, 9, 10}
		res := arraysgfg.MissingNumberArray_XOR(arr, len(arr))
		if res != 5 {
			t.Error("Invalid")
		}

	})
	t.Run("TestMissingNumberArray_Sum", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, /*5,*/ 6, 7, 8, 9, 10}
		res := arraysgfg.MissingNumberArray_Summation(arr, len(arr))
		if res != 5 {
			t.Error("Invalid")
		}

	})

}
