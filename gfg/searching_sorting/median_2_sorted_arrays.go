package searching_sorting

import "math"

func MedianOf2SortedArrays(arr1, arr2 []int, n int) int {
	i, j := 0, 0
	v1, v2 := 0, 0
	for {
		if arr1[i] <= arr2[j] {
			v1 = v2
			v2 = arr1[i]
			i++
		} else {
			v1 = v2
			v2 = arr2[j]
			j++
		}
		if i+j == n+1 {
			return ((v1 + v2) / 2)
		}
	}
}

func MedianOf2SortedArraysOptimized(arr1, arr2 []int, n int) int {
	return getMedian(arr1, arr2, 0, 0, n-1, n-1)
}

func getMedian(arr1, arr2 []int, startA, startB, endA, endB int) int {
	if int(math.Abs(float64(endA-startA))) == 1  {
		return (int(math.Max(float64(arr1[startA]), float64(arr2[startB]))) + int(math.Min(float64(arr1[endA]), float64(arr2[endB])))) / 2
	}
	m1 := median(arr1, startA, endA)
	m2 := median(arr2, startB, endB)

	if m1 == m2 {
		return m1
	}

	if m1 < m2 {
		return getMedian(arr1, arr2, (endA+startA+1)/2,
			startB, endA,
			(endB+startB+1)/2)
	} else {
		return getMedian(arr1, arr2, startA,
			(endB+startB+1)/2, (endA+startA+1)/2,
			endB);
	}
}

func median(arr []int, start, end int) int {
	n := (end - start) + 1
	if n%2 == 0 {
		return (arr[start+(n/2)] + arr[start+((n/2)-1)]) / 2
	} else {
		return arr[start+(n/2)]
	}
}
