package hash_test

import (
	hash "mundhrakeshav/dsa/gfg/hash"
	"reflect"
	"testing"
)

func Test0SumSubArray(t *testing.T) {

	t.Run("Test0SumSubArray", func(t *testing.T) {
		arr := []int{6, 3, -1, -3, 4, -2, 2, 4, 6, -12, -7}
		res := hash.Zero_sum_subarray(arr, len(arr))
		// 	 6  9   8   5  9   7  9  13  19 7  0
		//
		expected := []hash.Pair{
			{
				Start: 2,
				End:   4,
			},
			{
				Start: 2,
				End:   6,
			},
			{
				Start: 5,
				End:   6,
			},
			{
				Start: 6,
				End:   9,
			},
			{
				Start: 0,
				End:   10,
			},
		}

		if !reflect.DeepEqual(res, expected) {
			t.Error("Invalid")
		}
	})

}
