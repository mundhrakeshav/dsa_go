package arraysgfg_test

import (
	arraysgfg "mundhrakeshav/dsa/gfg/arrays"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestMaxSumSubarray(t *testing.T) {

	t.Run("TestBruteForce", func(t *testing.T) {
		arr := []int{-2, -3, 4, -1, -2, 1, 5, -3};
		res := arraysgfg.Max_Sum_Subarray_Brute_Force(arr)
		if res != 7 {
			t.Error("Invalid")
		}

	})
	t.Run("TestKadane", func(t *testing.T) {
		arr := []int{-2, -3, 4, -1, -2, 1, 5, -3};
		res := arraysgfg.Max_Sum_Subarray_Kadane(arr)
		if res != 7 {
			t.Error("Invalid")
		}

	})

}
