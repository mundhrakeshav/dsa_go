package arraysgfg_test

import (
	arrays "mundhrakeshav/dsa/gfg/arrays"
	"reflect"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestSlidingWindowMax(t *testing.T) {

	t.Run("TestSlidingWindowMax", func(t *testing.T) {
		arr := []int{1, 2, 3, 1, 4, 5, 2, 3, 6}
		arr = arrays.Sliding_Window_Max(arr, 3)
		if !reflect.DeepEqual(arr, []int{3, 3, 4, 5, 5, 5, 6}) {
			t.Error("Error")
		}
		arr = []int{8, 5, 10, 7, 9, 4, 15, 12, 90, 13}
		arr = arrays.Sliding_Window_Max(arr, 4)
		if !reflect.DeepEqual(arr, []int{10, 10, 10, 15, 15, 90, 90}) {
			t.Error("Error")
		}
	})
}
