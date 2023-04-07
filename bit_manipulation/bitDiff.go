package bitmanipulation

import "math"

func BitDiff(arr []uint8) (bitDiff uint) {

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i == j || j < i {
				continue
			}
			bitDiff += getBitDiffInternal(arr[i], arr[j]) * 2
		}
	}
	return bitDiff
}

func BitDiffEfficient(arr []uint8) (bitDiff uint) {
	// 8 as we use 8 bit nums
	for i := 0; i < 8; i++ {
		count := uint8(0)
		for j := 0; j < len(arr); j++ {
			if (arr[j] & (1 << i)) != 0 {
				count++
			}
		}
		bitDiff += (uint(count) * (uint(len(arr)) - uint(count))) * 2

	}
	return bitDiff
}

func getBitDiffInternal(num1, num2 uint8) (diff uint) {
	for num1 != 0 || num2 != 0 {
		xorNums := num1 ^ num2
		twos_compXORNums := ^xorNums + 1
		rmsbm := xorNums & twos_compXORNums
		rmsb := math.Log2(float64(rmsbm)) + 1
		num1 = num1 >> uint8(rmsb)
		num2 = num2 >> uint8(rmsb)
		diff++
	}
	return diff
}
