package arraysgfg_test

import (
	arraysgfg "mundhrakeshav/dsa/gfg/arrays"
	"testing"
)

func TestSubarrayWithGivenSum(t *testing.T) {

	t.Run("TestSubarrayWithGivenSum", func(t *testing.T) {
		arr := []int{1, 4, 20, 3, 10, 5};
		i, j := arraysgfg.Subarray_With_Given_Sum(arr, 33)
		if i!= 2 && j != 4{
			t.Error("Invalid")
		}
		arr = []int{1,2,3,7,5}
		i, j = arraysgfg.Subarray_With_Given_Sum(arr, 12)
		if i!= 1 && j != 3{
			t.Error("Invalid")
		}
		arr = []int{1,2,3,4,5,6,7,8,9,10}
		i, j = arraysgfg.Subarray_With_Given_Sum(arr, 10)
		if i!= 0 && j != 3{
			t.Error("Invalid")
		}
	})

}
