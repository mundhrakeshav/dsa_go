package arraysgfg

import "fmt"

// Given a sorted array of positive integers, rearrange the array alternately min_idx.e first element should be the maximum value, second minimum value, third-second max, fourth-second min and so on.

func Min_Max_Rearrange(arr []int) []int {
	min_idx, max_idx, max := 0, len(arr)-1, arr[len(arr)-1]+1
	for l := 0; l < len(arr); l++ {
		if l&1 == 0 {
			arr[l] += arr[max_idx] % max * max
			max_idx--
		} else {
			arr[l] += arr[min_idx] % max * max
			min_idx++
		}
	}

	fmt.Println(arr)

	for i := 0; i < len(arr); i++ {
		arr[i] /= max;
	}

	return arr
}
