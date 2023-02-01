package searching_sorting_test

import (
	searching_sorting "mundhrakeshav/dsa/gfg/searching_sorting"
	"testing"
)


func TestMedian2SortedArrays(t *testing.T) {

	t.Run("TestMedian2SortedArrays", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 5, 11}
		arr2 := []int{6, 8, 19, 21, 31}
		median := searching_sorting.MedianOf2SortedArrays(arr1, arr2, len(arr1))
		if median != 7 {
			t.Error("Invalid")
		}
	})
	
	t.Run("TestMedian2SortedArrays", func(t *testing.T) {
		arr1 := []int{1, 12, 15, 26, 38}
		arr2 := []int{2, 13, 17, 30, 45}
		median := searching_sorting.MedianOf2SortedArrays(arr1, arr2, len(arr1))
		if median != 16 {
			t.Error("Invalid")
		}
	})
	
	t.Run("TestMedian2SortedArrays", func(t *testing.T) {
		arr1 := []int{1, 12, 15, 26, 38}
		arr2 := []int{2, 13, 17, 30, 45}
		median := searching_sorting.MedianOf2SortedArraysOptimized(arr1, arr2, len(arr1))
		if median != 16 {
			t.Error("Invalid")
		}
	})

}
