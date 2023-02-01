package sort

import types "mundhrakeshav/dsa/types"

func Insertion_Sort[T types.Ordered](arr []T) []T {
	for i := 1; i < len(arr); i++ {
		val := arr[i]
			hole := i
			for hole > 0 &&  arr[hole-1] > val {
				arr[hole]= arr[hole - 1]
				hole--;
			}
		arr[hole] = val
	}
	return arr
}
