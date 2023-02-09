package arraysgfg

func AllocateMinPages(arr []int, n, m int) int {
	// arr => sorted array of pages in books
	// n => len(arr)
	// m => number of students

	low, high := arr[0], 0;
	for _, val := range arr {
		high += val;
	}
	res := -1
	for low <= high {
		mid := (low + high) / 2
		if isPossible(arr, n, m, mid) {
			res = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return res
}

func isPossible(arr []int, n, m, barrier int) bool {
	allocatedStu := 1
	pages := 0

	for i := 0; i < n; i++ {
		if arr[i] > barrier {
			return false
		}
		if pages+arr[i] > barrier {
			allocatedStu += 1
			pages += arr[i]
		} else {
			pages += arr[i]
		}
	}

	return  allocatedStu <= m;
}
