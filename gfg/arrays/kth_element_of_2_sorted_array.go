package arraysgfg

func KthElementInSortedArray(a1 []int, a2 []int, k int) int {
	i, j, p := 0, 0, 0

	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			p++
			if p == k {
				return a1[i]
			}
			i++
		} else {
			p++
			if p == k {
				return a2[j]
			}
			j++
		}
	}

	for i < len(a1) {
		p++
		if p == k {
			return a1[i]
		}
		i++
	}
	for i < len(a2) {
		p++
		if p == k {
			return a2[j]
		}
		j++
	}

	return -1
}
