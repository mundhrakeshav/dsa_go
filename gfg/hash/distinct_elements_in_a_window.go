package hash

func DistinctElementsInAWindow(arr []int, k int) []int {
	dist_count := 0
	res := make([]int, 0)
	hm := make(map[int]int)

	for i := 0; i < k; i++ {
		if _, exists := hm[arr[i]]; !exists {
			dist_count++
		}
		hm[arr[i]] += 1
	}
	res = append(res, dist_count)

	for i := k; i < len(arr); i++ {
		if hm[arr[i-k]] == 1 {
			dist_count--
			delete(hm, arr[i-k])
		} else {
			hm[arr[i-k]] -= 1;
		}

		if _, exists := hm[arr[i]]; exists {
			hm[arr[i]] += 1
		} else {
			dist_count++
			hm[arr[i]] = 1
		}
		res = append(res, dist_count)
	}

	return res
}

// `1, 2, 1, 3`, 4, 2, 3}
