package hash

type Pair struct {
	Start, End int;
}

// arr := []int{6, 3, -1, -3, 4, -2, 2, 4, 6, -12, -7}
// // 	 6  9   8   5  9


func Zero_sum_subarray(arr []int, n int) []Pair {
	hm := make(map[int][]int, 0)
	sum := 0;
	res := make([]Pair, 0)
	hm[0] = append(hm[0], 0)
	for i := 0; i < n; i++ {
		sum += arr[i]
	
		valArr, ok := hm[sum];

		if ok {
			for _, value := range valArr {
				res = append(res, Pair{
					Start: value,
					End: i,
				})
			}
		}


		hm[sum] = append(hm[sum], i + 1)
	}

	return res;

}
