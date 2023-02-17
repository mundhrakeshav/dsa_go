package hash

import "math"

func LongestConsecutiveSubsequence(arr []int) int {
	hm := make(map[int]bool, 0)
	longest := 0
	for i := 0; i < len(arr); i++ {
		hm[arr[i]] = true
	}

	for i := 0; i < len(arr); i++ {
		currentLongest := 0
		j := arr[i]
		for j >= 0 && hm[j] {
			currentLongest++
			j--
		}
		longest = int(math.Max(float64(longest), float64(currentLongest)))
	}
	return longest
}
