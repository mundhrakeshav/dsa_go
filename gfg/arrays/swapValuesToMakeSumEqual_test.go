package arraysgfg_test

import (
	"fmt"
	arraysgfg "mundhrakeshav/dsa/gfg/arrays"
	"testing"
)

func TestSwapValsToMakeSumEqual(t *testing.T) {

	t.Run("TestSwapValsToMakeSumEqual", func(t *testing.T) {
		arr1 := []int{4, 1, 2, 1, 1, 2};
		arr2 := []int{3, 6, 3, 3};
		i,j := arraysgfg.SwapValsToMakeSumEqual(arr1, arr2, len(arr1), len(arr2))
		fmt.Println(i, j)
		arr1 = []int{5, 7, 4, 6};
		arr2 = []int{1, 2, 3, 8};
		i,j = arraysgfg.SwapValsToMakeSumEqual(arr1, arr2, len(arr1), len(arr2))
		fmt.Println(i, j)
	})

}
