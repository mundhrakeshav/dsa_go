package matrix

import "math"

func firstOneSearch(arr []int, currentMaxOnesIndex int) int {
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
	maxRow, maxRowFirstOneIndex:= 0, math.MaxInt;

	for i, v := range arr {
		x := firstOneSearch(v, maxRowFirstOneIndex)
		
		if x >= 0 && len(v) - x > len(arr[maxRow]) - maxRowFirstOneIndex {
			maxRow = i;
			maxRowFirstOneIndex = x
		}
	}
	return maxRow
}
