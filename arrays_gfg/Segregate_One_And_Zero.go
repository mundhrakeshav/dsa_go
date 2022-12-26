package arraysgfg

func Segregate_One_and_Zero(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		for arr[left] == 0 && left < right {
			left++
		}
		for arr[right] == 1 && left < right {
			right--
		}

		if left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return arr
}
