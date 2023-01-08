package arraysgfg

import "math"

// Given an array arr[] of size N. The task is to find the sum of the contiguous subarray within a arr[] with the largest sum.



func Max_Sum_Subarray_Brute_Force(arr []int) int {
	maxSum := int(math.Inf(-12))
	for i := 0; i < len(arr); i++ {

		sum := 0
		for j := i; j < len(arr); j++ {
			sum += arr[j]
			if maxSum < sum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func Max_Sum_Subarray_Kadane(arr []int) int {
	maxSoFar := 0
	maxCurrent := 0
	for i := 0; i < len(arr); i++ {
		maxCurrent += arr[i]
		if maxCurrent < 0 {
			maxCurrent = 0
		}
		if maxSoFar < maxCurrent {
			maxSoFar = maxCurrent;
		}
	}
	return maxSoFar
}
