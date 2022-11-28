package leetcode

import (
	"fmt"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestRemoveDuplicates(t *testing.T) {

	t.Run("TwoSum", func(t *testing.T) {

		arr := []int{0,0, 1, 1, 2, 3, 3, 3, 7, 7, 7, 7}
		x := removeDuplicates(arr)
		// if arr[0] != 0 && arr[1] != 1 {
		// 	t.Error("Invalid Return")
		// }
		fmt.Println(x, arr)
	})

}
