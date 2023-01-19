package arraysgfg_test

import (
	"fmt"
	arraysgfg "mundhrakeshav/dsa/gfg/arrays"
	"testing"
)

func TestReverseSubarray(t *testing.T) {

	t.Run("TestReverseSubarray", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
		arr = arraysgfg.Reverse_Subarrays(arr, 3)
		fmt.Println(arr)
		// if res[0] != 2 && res[1] != 5 && res[3] != 17 {
		// 	t.Error("Invalid")
		// }

	})

}
