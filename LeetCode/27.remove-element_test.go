package leetcode

import (
	"fmt"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestRemoveElement(t *testing.T) {

	t.Run("Remove Element", func(t *testing.T) {

		arr := []int{3,2,2,3}
		x := removeElement(arr, 3)
		fmt.Println(x, arr)
		
	})

}
