package searching_sorting_test

import (
	searching_sorting "mundhrakeshav/dsa/gfg/searching_sorting"
	"testing"
)

func TestInversion(t *testing.T) {

	t.Run("TestInversion", func(t *testing.T) {
		arr := []int{8,4,2,1}
		_, count := searching_sorting.Merge_Sort_Inversion(arr)
		if count != 6 {
			t.Error("Inavlid")
		}
		arr = []int{1, 20, 6, 4, 5}
		_, count = searching_sorting.Merge_Sort_Inversion(arr)
		if count != 5 {
			t.Error("Inavlid")
		}
		
		arr = []int{1, 2, 3, 4, 5}
		_, count = searching_sorting.Merge_Sort_Inversion(arr)
		if count != 0 {
			t.Error("Inavlid")
		}
	})

}
