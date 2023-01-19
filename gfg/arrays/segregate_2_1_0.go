package arraysgfg

func Segregate_2_1_0(arr []int) []int {
	ptr1, ptr2, i := 0, len(arr)-1, 0

	for i < ptr2 {
		if arr[i] == 0 {
			arr[ptr1], arr[i] = arr[i], arr[ptr1]
			ptr1++
			i++
		} else if arr[i] == 2 {
			arr[ptr2], arr[i] = arr[i], arr[ptr2]
			ptr2--;
		} else {
			i++;
		}
	}
	return arr;
}
