package misc

import (
	"fmt"
	"time"
)

func Prime_factors(n uint32) []uint32 {
	timeN := time.Now()
	s := make([]uint32, 0, 10)
	if n <= 1 {
		return s
	}
	m := n

	// Check for 2 as Prime Factor
	for m%2 == 0 {
		s = append(s, 2)
		m = m / 2
	}

	// Check for 3 as Prime Factor
	for m%3 == 0 {
		s = append(s, 3)
		m = m / 3
	}

	// Check for all numbers except multiples of 2 and 3
	for i := uint32(5); m > 1; i = i + 6 {
		for m%i == 0 {
			s = append(s, i)
			m = m / i
		}
		for m%(i+2) == 0 {
			s = append(s, (i + 2))
			m = m / (i + 2)
		}
	}
	fmt.Println(time.Since(timeN))
	return s
}

// func Prime_factors(n uint32) []uint32 {
// 	timeN := time.Now()
// 	s := make([]uint32, 0, 10)
// 	if n <= 1 {
// 		return s
// 	}
// 	m := n

// 	for i := uint32(2); m > 1; i++ {
// 		for m%i == 0 {
// 			s = append(s, i)
// 			m = m / i
// 		}
// 	}
// 	fmt.Println(time.Since(timeN))
// 	return s
// }
