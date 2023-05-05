package search

import (
	"fmt"
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

		arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		key := 5
		low := 0
		high := len(arr) - 1
		expected := 4
		result := BinarySearchRecursive(arr, key, low, high)
		if result != expected {
			t.Errorf("BinarySearchRecursive(%v, %d, %d, %d) = %d; want %d", arr, key, low, high, result, expected)
		}
	
		// key = 11
		// low = 0
		// high = len(arr) - 1
		// expected = -1
		// result = BinarySearchRecursive(arr, key, low, high)
		// if result != expected {
		// 	t.Errorf("BinarySearchRecursive(%v, %d, %d, %d) = %d; want %d", arr, key, low, high, result, expected)
		// }
	
		// arr = []int{}
		// key = 5
		// low = 0
		// high = len(arr) - 1
		// expected = -1
		// result = BinarySearchRecursive(arr, key, low, high)
		// if result != expected {
		// 	t.Errorf("BinarySearchRecursive(%v, %d, %d, %d) = %d; want %d", arr, key, low, high, result, expected)
		// }

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

    // // Test case where key is at the beginning of the array
    // arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    // key = 1
    // low = 0
    // high = len(arr) - 1
    // expected = 0
    // result = BinarySearchRecursive(arr, key, low, high)
    // if result != expected {
    //     t.Errorf("BinarySearchRecursive(%v, %d, %d, %d) = %d; want %d", arr, key, low, high, result, expected)
    // }

    // // Test case where key is at the end of the array
    // arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    // key = 10
    // low = 0
    // high = len(arr) - 1
    // expected = 9
    // result = BinarySearchRecursive(arr, key, low, high)
    // if result != expected {
    //     t.Errorf("BinarySearchRecursive(%v, %d, %d, %d) = %d; want %d", arr, key, low, high, result, expected)
    // }

    // // Test case where key is not in the array
    // arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    // key = 11
    // low = 0
    // high = len(arr) - 1
    // expected = -1
    // result = BinarySearchRecursive(arr, key, low, high)
    // if result != expected {
    //     t.Errorf("BinarySearchRecursive(%v, %d, %d, %d) = %d; want %d", arr, key, low, high, result, expected)
    // }

    // // Test case where array has only one element and it is the key
    // arr = []int{5}
    // key = 5
    // low = 0
    // high = len(arr) - 1
    // expected = 0
    // result = BinarySearchRecursive(arr, key, low, high)
    // if result != expected {
    //     t.Errorf("BinarySearchRecursive(%v, %d, %d, %d) = %d; want %d", arr, key, low, high, result, expected)
    // }

    // // Test case where array has only one element and it is not the key
    // arr = []int{5}
    // key = 6
    // low = 0
    // high = len(arr) - 1
    // expected = -1
    // result = BinarySearchRecursive(arr, key, low, high)
    // if result != expected {
    //     t.Errorf("BinarySearchRecursive(%v, %d, %d, %d) = %d; want %d", arr, key, low, high, result, expected)
    // }

    // // Test case where array is empty
    // arr = []int{}
    // key = 5
    // low = 0
    // high = len(arr) - 1
    // expected = -1
    // result = BinarySearchRecursive(arr, key, low, high)
    // if result != expected {
    //     t.Errorf("BinarySearchRecursive(%v, %d, %d, %d) = %d; want %d", arr, key, low, high, result, expected)
    // }
	////////////////////////////////!




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
		
		arr = []int{2, 8, 9, 11, 16, 32, 41, 45, 90}
		x = BinarySearchIterative(arr, 8)
		fmt.Println(x)
	})

}