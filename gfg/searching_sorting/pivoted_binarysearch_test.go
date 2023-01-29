package searching_sorting_test

import (
	searching_sorting "mundhrakeshav/dsa/gfg/searching_sorting"
	"testing"
)

func TestPivotedBinarySearch(t *testing.T) {

	t.Run("TestPivotedBinarySearch", func(t *testing.T) {
		arr := []int{5, 6, 7, 1, 2, 3}
		count := searching_sorting.PivotedBinarySearch(arr, 2)
		if count != 4 {
			t.Error("Inavlid")
		}
		arr = []int{5, 6, 7, 8, 0, 1, 2, 3, 4}
		count = searching_sorting.PivotedBinarySearch(arr, 2)
		if count != 6 {
			t.Error("Inavlid")
		}
		
		arr = []int{5, 6, 7, 8, 0, 1, 2, 3, 4}
		count = searching_sorting.PivotedBinarySearch(arr, 4)
		if count != 8 {
			t.Error("Inavlid")
		}
	})

}
