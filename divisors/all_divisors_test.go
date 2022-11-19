package divisor

import (
	"fmt"
	"testing"
)

func TestAllDivisors(t *testing.T) {
	t.Run("All Divisors", func(t *testing.T) {
		s := All_Divisors(30)
		fmt.Println(s)

	})
}
