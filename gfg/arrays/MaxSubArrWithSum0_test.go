package arraysgfg_test

import (
	arraysgfg "mundhrakeshav/dsa/gfg/arrays"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestMaxSum0Subarray(t *testing.T) {

	t.Run("TestMaxSum0Subarray", func(t *testing.T) {
		arr := []int{15, -2, 2, -8, 1, 7, 10, 23, };
		res := arraysgfg.MaxSubArrSum0(arr, len(arr))
		if res != 5 {
			t.Error("Invalid")
		}
		arr = []int{1, 2, 3};
		res = arraysgfg.MaxSubArrSum0(arr, len(arr))
		if res != 0 {
			t.Error("Invalid")
		}
		
		arr = []int{1, 0, 3};
		res = arraysgfg.MaxSubArrSum0(arr, len(arr))
		if res != 1 {
			t.Error("Invalid")
		}
	})

	t.Run("TestMaxSum0SubarrayBruteForce", func(t *testing.T) {
		arr := []int{15, -2, 2, -8, 1, 7, 10, 23, };
		res := arraysgfg.MaxSubArrSum0BruteForce(arr, len(arr))
		if res != 5 {
			t.Error("Invalid")
		}
		arr = []int{1, 2, 3};
		res = arraysgfg.MaxSubArrSum0BruteForce(arr, len(arr))
		if res != 0 {
			t.Error("Invalid")
		}
		arr = []int{1, 0, 3};
		res = arraysgfg.MaxSubArrSum0BruteForce(arr, len(arr))
		if res != 1 {
			t.Error("Invalid")
		}
	})

}
