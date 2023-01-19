package searching_sorting_test

import (
	"fmt"
	searching_sorting "mundhrakeshav/dsa/gfg/searching_sorting"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	t.Run("TestBinarySearch", func(t *testing.T) {
		arr := []int{2, 8, 9, 11, 16, 32, 41, 45, 90}
		x := searching_sorting.BinarySearchRecursive(arr, 0, len(arr)-1, 2)
		fmt.Println(x)
		y := searching_sorting.BinarySearchIter(arr, 0, len(arr)-1, 41)
		fmt.Println(y)
	})

}
