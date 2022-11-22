package search

import (
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("Binary Search", func(t *testing.T) {
		arr := []int{1, 2, 5, 6, 11, 25, 32, 56, 345, 4574}
		x := BinarySearchRecursive(arr, 6, 0, len(arr) - 1)
		if arr[x] != 6 {
			t.Error("Incorrect Key")
		}

		arr = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
		x = BinarySearchRecursive(arr, 31, 0, len(arr) - 1)
		if arr[x] != 31 {
			t.Error("Incorrect Key")
		}
		arr = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
		x = BinarySearchIterative(arr, 31)
		if arr[x] != 31 {
			t.Error("Incorrect Key")
		}

		arr = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
		x = BinarySearchRecursive(arr, 63, 0, len(arr) - 1)
		if arr[x] != 63 {
			t.Error("Incorrect Key")
		}
		arr = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
		x = BinarySearchIterative(arr, 63)
		if arr[x] != 63 {
			t.Error("Incorrect Key")
		}

		arr = []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
		x = BinarySearchIterative(arr, 61)
		if x != -1 {
			t.Error("Incorrect Key")
		}
	})

}
