package search

func BinarySearchRecursive(arr []int, key int, low int, high int) int {
	if len(arr) == 0 {
		return -1
	}
	mid := (high + low) / 2
	arr_mid := arr[mid]
	if arr_mid == key {
		return mid
	}

	if key > arr[mid] {
		return BinarySearchRecursive(arr, key, mid+1, high)
	} else {
		return BinarySearchRecursive(arr, key, low, mid-1)
	}

}

func BinarySearchIterative(arr []int, key int) int {
	if len(arr) == 0 {
		return -1
	}
	low := 0
	high := len(arr) - 1
	for high > low {
		mid := (low + high) / 2
		if arr[mid] == key {
			return mid
		}
		if arr[mid] > key {
			high = mid - 1
			} else {
			low = mid + 1
		}
	} 
	return -1;
}
