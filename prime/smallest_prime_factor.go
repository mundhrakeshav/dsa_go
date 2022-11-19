package misc

import (
	"fmt"
	"math"
	"time"
)

func Smallest_Prime_Factors(n uint32) []uint32 {
	timeN := time.Now()
	s := make([]uint32, n+1)

	for i := uint32(2); i <= n ; i++ {
		if s[i] == 0 {
			for j := uint32(i); j <= n; j = j + i {
				if s[j] == 0 {
					s[j] = i
				}
			} 
		}
	}
	fmt.Println(time.Since(timeN))
	return s
}

// func Smallest_Prime_Factors(n uint32) []uint32 {
// 	timeN := time.Now()
// 	s := make([]uint32, n+1)
// 	// s[0] = 0
// 	// if n >= 1 {
// 	// 	s[1] = 1
// 	// }

// 	for i := uint32(n + 1); i > 1; i-- {
// 		if checkPrimeNumber(i) {
// 			m := i
// 			for j := i; m < n+1; {
// 				s[m] = j
// 				m = m + j
// 			}
// 		}
// 	}
// 	fmt.Println(time.Since(timeN))
// 	return s
// }

func checkPrimeNumber(num uint32) bool {
	if num < 2 {
		return true
	}
	sq_root := uint32(math.Sqrt(float64(num)))
	for i := uint32(2); i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
