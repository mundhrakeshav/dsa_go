package arraysgfg

func Stock_buy_sell(arr []int) int {
	mp :=  0

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			mp += arr[i] - arr[i-1]
		}	
	}

	return mp
}
