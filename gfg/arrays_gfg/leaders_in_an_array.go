package arraysgfg

/*
	Write a program to print all the LEADERS in the array.
	An element is a leader if it is greater than all the elements to its right side.
	And the rightmost element is always a leader.
*/

func Leaders_In_A_Array(arr []int) []int {
	var max int;
	var res []int;
	max = arr[len(arr) - 1]
	res = append(res, max)
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > max {
			max = arr[i];
			res = append(res, max)
		}
	}
	return res;
}