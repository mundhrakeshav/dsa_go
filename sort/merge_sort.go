package sort

func merge(left []int, right []int) []int {
	arr := make([]int, len(left) + len(right))
	i := 0
	j := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[i+j] = left[i]
			i++;
			} else {
				arr[i+j] = right[j]
				j++
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

func Merge_Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var middle = len(arr) / 2
	var a = Merge_Sort(arr[:middle])
	var b = Merge_Sort(arr[middle:])
	return merge(a, b)
}
