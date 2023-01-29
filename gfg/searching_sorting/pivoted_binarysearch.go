package searching_sorting

func PivotedBinarySearch(arr []int, key int) int {
	low, high := 0, len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == key {
			return mid
		}
		if arr[low] < arr[mid] {
			// Left half sorted
			if key >= arr[low] && key <= arr[mid] {
				// Element on left
				high = mid - 1
			} else {
				// Element on right
				low = mid + 1
			}
		} else {
			// Right Half Sorted
			if key >= arr[mid] && key <= arr[high] {
				// Element on right
				low = mid + 1
				} else {
					// Element on left
					high = mid - 1
			}

		}
	}

	return -1
}
