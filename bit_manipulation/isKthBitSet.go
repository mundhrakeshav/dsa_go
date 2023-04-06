package bitmanipulation

func IsKthBitSet(num, k uint8) bool {
	num = num >> k;
	// Approach 1
	// twos_comp := (^num) + 1
	// rmsbm := num & twos_comp
	// return math.Log2(float64(rmsbm)) == 0;

	// Approach 2
 	return float64(num & 1) == 1;
}