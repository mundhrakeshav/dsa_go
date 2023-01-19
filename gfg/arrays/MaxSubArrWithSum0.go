package arraysgfg

import "math"

func MaxSubArrSum0BruteForce(arr []int, n int) int {
	res := 0
	for i := 0; i < n; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += arr[j]
			if sum == 0 {
				res = int(math.Max(float64(res), float64(j - i + 1)))
			}
		}
	}
	return res
}

func MaxSubArrSum0(arr []int, n int) int {
	max := 0
	sum := 0
	record := make(map[int]int)
	// 1, 0 ,3
	for i := 0; i < n; i++ {
		sum += arr[i]
		if val, isPresent := record[sum]; !isPresent {
			record[sum] = i;
		} else {
			max = i - val;
		}
	}
	return max
}
