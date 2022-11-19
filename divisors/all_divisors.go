package divisor

import (
	"fmt"
	"time"
)

func All_Divisors(n uint32) []uint32 {
	timeN := time.Now()
	s := make([]uint32, 0, 10)
	s = append(s, 1, n)
	
	for i := uint32(2); i * i <= n; i++ {
		if n % i == 0 {
			s = append(s, i)
			if i != n/i {
				s = append(s, n/i)
				
			}
		}
	}


	fmt.Println(time.Since(timeN))
	return s
}
