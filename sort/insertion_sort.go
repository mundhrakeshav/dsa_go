package sort

import types "mundhrakeshav/dsa/types"

func Insertion_Sort[T types.Ordered](arr []int) []int {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
		hole := i
		for hole > 0 && arr[hole -1] > val {
			arr[hole] = arr[hole-1]
			hole = hole - 1
		}
		arr[hole] = val;
	}
	return arr
}
