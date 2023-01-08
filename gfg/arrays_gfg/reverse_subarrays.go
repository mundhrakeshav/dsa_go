package arraysgfg

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Reverse_Subarrays(arr []int, k int) []int {
	n := len(arr)
	for i := 0; i < n; i += k {
		left, right := i, min(i+k-1, n-1)

		for left < right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}
	return arr
}