package matrix

func firstOneSearch(arr []int) int {
	low, high := 0, len(arr)-1
	first := len(arr)
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == 1 {
			first = mid
			high--
		} else {
			low++
		}
	}
	return first
}

// Given a boolean 2D array, where each row is sorted. Find the row with the maximum number of 1s.
func Max_num_of_ones(arr [][]int) int {
	max, total:= -1, 0;

	for i, v := range arr {
		x := firstOneSearch(v)
		totalOnes := len(v) - x;
		if totalOnes > total {
			max = i;
			total = totalOnes
		}
	}
	return max
}
