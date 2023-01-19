package searching_sorting

func partition(arr []int, low, high int) int {
	_pivotElement := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < _pivotElement {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	_pivot := partition(arr, low, high)
	quickSort(arr, low, _pivot - 1)
	quickSort(arr, _pivot + 1, high)
}


func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}