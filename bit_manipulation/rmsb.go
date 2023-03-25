package bitmanipulation

import (
	"math"
)

// Position of rightmost set bit
// right most set bit mask = x & x``


func GetRightMostSetBit(num uint8) float64 {
	ones_comp := ^num
	twos_comp := ones_comp + 1;
	rmsbm := num & twos_comp
	return math.Log2(float64(rmsbm)) + 1;
}