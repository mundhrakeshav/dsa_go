package arraysgfg

// Given an array arr[] of size N-1 with integers in the range of [1, N], the task is to find the missing number from the first N integers.

func MissingNumberArray_XOR(arr []int , n int) int {
	XOR_Elements := 0;
	XOR_Idx := 0

	for i := 0; i < n; i++ {
		XOR_Elements = XOR_Elements ^ arr[i]
	}
	for i := 0; i <= n + 1; i++ {
		XOR_Idx = XOR_Idx ^ i
	}

	return XOR_Idx ^ XOR_Elements;

}

func MissingNumberArray_Summation(arr []int , n int) int {

	sum_expected := ((n+1)*(n+2))/2
	sum_actual := 0
	for i := 0; i < len(arr); i++ {
		sum_actual += arr[i]
	}
	return sum_expected - sum_actual
}