package arraysgfg

import "math"

func Sliding_Window_Max(arr []int, k int) []int {
	if k > len(arr) {
		return []int{}
	}
	max := math.MinInt
	retArr := make([]int, 0)
	for i := 0; i < k; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	retArr = append(retArr, max)
	for i := k; i < len(arr); i++ {
		if max < arr[i] {
			max = arr[i]
		}
		retArr = append(retArr, max)
	}
	return retArr
}
