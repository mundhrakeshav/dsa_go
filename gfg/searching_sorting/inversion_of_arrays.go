package searching_sorting

func mergeAndCount(left, right []int, swaps *int) ([]int) {
	swp := *swaps;
	arr := make([]int, len(left) + len(right));
 	i, j := 0, 0; 
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[i+j] = left[i];
			i++
		} else {
			arr[i+j] = right[j];
			swp += len(left) - i;
			j++
			// len(left) - j
		}
	}
	*swaps = swp;
	for i < len(left) {
		arr[i + j] = left[i]
		i++
	}
	
	for j < len(right) {
		arr[i+j] = right[j]
		j++
	}
	return arr;
}


func Merge_Sort_Inversion(arr []int) ([]int, int) {
	swaps:=0;
	if len(arr) < 2 {
		return arr, 0;
	}
	middle := (len(arr)) / 2
	a, swps :=Merge_Sort_Inversion(arr[:middle])
	swaps += swps
	b, swps :=Merge_Sort_Inversion(arr[middle:])
	swaps += swps
	return mergeAndCount(a, b, &swaps), swaps;
}