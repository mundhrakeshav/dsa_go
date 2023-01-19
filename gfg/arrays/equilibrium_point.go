package arraysgfg

func Equilibrium_Point(arr []int, n uint) int {
	totSum := 0;
	for i := 0; i < len(arr); i++ {
		totSum += arr[i]
	}
	sum := 0;
	for i := 0; i < len(arr); i++ {
		if sum  == totSum - (sum + arr[i]) {
			return i
		}
		sum += arr[i]
	}

	return -1;

}