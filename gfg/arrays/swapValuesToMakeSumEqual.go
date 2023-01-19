package arraysgfg

import (
	"fmt"
	"sort"
)

func getSum(arr []int, l int) int {
	sum:=0;
	for i := 0; i < l; i++ {
		sum += arr[i]
	}
	return sum;
}

func SwapValsToMakeSumEqual(arr1, arr2 []int, n1, n2 int) (int, int) {
	
	sum1, sum2 := getSum(arr1, n1), getSum(arr2, n2)
	target := (sum1 - sum2 ) / 2;
	if target == 0 || (sum1 - sum2 ) % 2 != 0 {
		return -1, -1;
	}
	sort.Slice(arr1, func(i, j int) bool {
		return arr1[i] < arr1[j]
	})
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] < arr2[j]
	})

	fmt.Println(arr1, arr2)
	i, j := 0, 0;
	for i < n1 && j < n2 {
		if arr1[i] - arr2[j] < target {
			i++;
		} else if arr1[i] - arr2[j] > target {
			j++;
		} else {
			return i, j
		}
	}
	return -1, -1;
}