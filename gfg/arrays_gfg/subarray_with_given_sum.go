package arraysgfg

// Given an array arr[] of non-negative integers and an integer sum, find a subarray that adds to a given sum.

func Subarray_With_Given_Sum(arr []int, sum int) (int, int) {
	start, currentSum := 0, 0
	
	for i := 0; i < len(arr); i++ {
		currentSum += arr[i]
		for currentSum > sum {
			currentSum -= arr[start]
			start++;
		}
		if currentSum == sum {
			return start, i
		}
	}

	return -1, -1
}
