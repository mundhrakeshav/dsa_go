package searching_sorting

func BinarySearchRecursive(arr []int, low, high, key int) int {
	if low > high {
		return -1;
	}

	mid := low + ((high - low) / 2);

	if arr[mid] == key {
		return mid;
	}

	if arr[mid] > key {
		return BinarySearchRecursive(arr, low, mid - 1, key)
		} else {
		return BinarySearchRecursive(arr, mid + 1, high, key)

	}
}


func BinarySearchIter(arr []int, low, high, key int) int {
	for low <= high {
		mid := low + (high - low) / 2;
		if key == arr[mid] {
			return mid
		} else if key < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1;
}