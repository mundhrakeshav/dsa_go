package divisor

import (
	"fmt"
	prime "mundhrakeshav/dsa/prime"
	"time"
)

func Only3_Divisors(n uint32) []uint32 {
	timeN := time.Now()
	s := make([]uint32, 0, 10)
	
	for i := uint32(2); i * i <= n; i++ {
		if prime.Is_Prime(i) {
			s = append(s, i*i)
		}		
	}


	fmt.Println(time.Since(timeN))
	return s
}
