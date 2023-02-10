package hash_test

import (
	hash "mundhrakeshav/dsa/gfg/hash"
	"testing"
)

func TestCommonElementsIn3SortedArrays(t *testing.T) {

	t.Run("TestCommonElementsIn3SortedArrays", func(t *testing.T) {
		arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18}
		arr2 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18}
		arr3 := []int{3, 6, 9, 12, 15, 18}
		res := hash.CommonElementsIn3SortedArrays(arr1, arr2, arr3, len(arr1), len(arr2), len(arr2))
		if res[0] != 6 && res[1] != 12 && res[2] != 18 { 
			t.Error("Invalid")
		}
	})
	
	t.Run("TestCommonElementsIn3SortedArrays", func(t *testing.T) {
		arr1 := []int{1, 5, 10, 20, 40, 80} 
		arr2 := []int{6, 7, 20, 80, 100} 
		arr3 := []int{3, 4, 15, 20, 30, 70, 80, 120} 
		res := hash.CommonElementsIn3SortedArrays(arr1, arr2, arr3, len(arr1), len(arr2), len(arr2))
		if res[0] != 20 && res[1] != 80 { 
			t.Error("Invalid")
		}
	})
	
	t.Run("TestCommonElementsIn3SortedArrays", func(t *testing.T) {
		arr1 := []int{1, 5, 5}
		arr2 := []int{3, 4, 5, 5, 10}
		arr3 := []int{5, 5, 10, 20}
		res := hash.CommonElementsIn3SortedArrays(arr1, arr2, arr3, len(arr1), len(arr2), len(arr2))
		if res[0] != 5 && res[1] != 5 { 
			t.Error("Invalid")
		}
	})
	
}
