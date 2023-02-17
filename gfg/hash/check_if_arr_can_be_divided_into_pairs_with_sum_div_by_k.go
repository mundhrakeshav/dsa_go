package hash

func CheckIfArrayCanBeDividedIntoPairsDivisibleByK(arr []int, k int) bool {
	if len(arr)%2 == 1 {
		return false
	}

	hm := make(map[int]int, 0)

	for i := 0; i < len(arr); i++ {
		hm[(((arr[i] % k) + k) % k)] += 1
	}

	if hm[0] % 2 == 1 {
		return false
	}

	for i := 1; i < k; i++ {
		
			if hm[i] != hm[k - i] {
				return false
			}
	}

	return true
}
