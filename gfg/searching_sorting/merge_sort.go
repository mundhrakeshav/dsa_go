package searching_sorting


func merge(left []int, right []int) []int {
	arr := make([]int, len(left)+ len(right) )
	i, j := 0,0;
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[i+j] = left[i]
			i++
		} else {
			arr[i+j] = right[j]
			j++;
		}
	}
	for i < len(left) {
		arr[i+j] = left[i]
		i++
	}

	for j < len(right) {
		arr[i+j] = right[j]
		j++
	}

	return arr;
}

func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	middle := len(arr)/2
	return merge(MergeSort(arr[:middle]), MergeSort(arr[middle:]))
}