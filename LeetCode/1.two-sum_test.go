package leetcode

import (
	"fmt"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestTwoSum(t *testing.T) {

	t.Run("TwoSum", func(t *testing.T) {

		arr := []int{2,7,11,15}
		arr = TwoSum(arr, 9);
		if arr[0] != 0 && arr[1] != 1 {
			t.Error("Invalid Return")
		}
		arr = []int{-5, 7, 4, -2, 6, 5, 8, 3, 2, -7, -1, 0, -3, 9, -6, -4, 10, 9, 1, -8, -9, -13}
		arr = TwoSum(arr, -10);
		fmt.Println(arr)
		// if arr[0] != 0 && arr[1] != 1 {
		// 	t.Error("Invalid Return")
		// }
	})

}
