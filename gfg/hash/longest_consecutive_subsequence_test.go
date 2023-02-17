package hash_test

import (
	hash "mundhrakeshav/dsa/gfg/hash"
	"testing"
)

func TestLongestConsecutiveSubsequence(t *testing.T) {

	t.Run("TestCommonElementsIn3SortedArrays", func(t *testing.T) {
		arr := []int{9, 7, 5, 3}
		res := hash.LongestConsecutiveSubsequence(arr)
		if res!=1 { 
			t.Error("Invalid")
		}
		
	})
	
	t.Run("TestCommonElementsIn3SortedArrays", func(t *testing.T) {
		arr := []int{1, 9, 3, 10, 4, 20, 2}
		res := hash.LongestConsecutiveSubsequence(arr)
		if res != 4 { 
			t.Error("Invalid")
		}
	})
	
	t.Run("TestCommonElementsIn3SortedArrays", func(t *testing.T) {
		arr := []int{36, 41, 56, 35, 44, 33, 34, 92, 43, 32, 42}
		res := hash.LongestConsecutiveSubsequence(arr)
		if res != 5 { 
			t.Error("Invalid")
		}
	})
	
	
}
