package leetcode

import (
	"testing"
)

// TestStackArray for testing Stack with Array
func TestRemoveDuplicates(t *testing.T) {

	t.Run("TwoSum", func(t *testing.T) {

		arr := []int{0,0, 1, 1, 2, 3, 3, 3, 7, 7, 7, 7}
		x := removeDuplicates(arr)
		if x != 5 {
			t.Error("Invalid Return")
		}
		
		if arr[0] != 0 {
			t.Error("Invalid Return")
		}
		if arr[1] != 1 {
			t.Error("Invalid Return")
		}
		if arr[2] != 2 {
			t.Error("Invalid Return")
		}
		if arr[3] != 3 {
			t.Error("Invalid Return")
		}
		if arr[4] != 7 {
			t.Error("Invalid Return")
		}
		
	})

}
