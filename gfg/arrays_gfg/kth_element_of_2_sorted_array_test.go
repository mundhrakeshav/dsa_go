package arraysgfg_test

import (
	"fmt"
	arraysgfg "mundhrakeshav/dsa/gfg/arrays_gfg"
	"testing"
)

func TestKthElementInSortedArray(t *testing.T) {

	t.Run("TestReverseSubarray", func(t *testing.T) {
		a1 := []int{2,3,6,7,9}
		a2 := []int{1,4,8,10}
		x := arraysgfg.KthElementInSortedArray(a1, a2, 5)
		fmt.Println(x)
		// if res[0] != 2 && res[1] != 5 && res[3] != 17 {
		// 	t.Error("Invalid")
		// }

	})

}
