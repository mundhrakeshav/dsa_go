package hash

func CommonElementsIn3SortedArrays(arr1, arr2, arr3 []int, l, n, m int) []int {
	i, j, k := 0, 0, 0
	res := make([]int, 0)
	for (i <= l - 1 && j <= n - 1 && k < m - 1) {
		for(arr1[i] > arr2[j]) {
			j++;
		}
		for(arr1[i] > arr3[k]) {
			k++;
		}
		if (arr1[i] == arr2[j]) && arr2[j] == arr3[k] {
			res = append(res, arr1[i])
		}
		i++;
	}
	return res;
}
