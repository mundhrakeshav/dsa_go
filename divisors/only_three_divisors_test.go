package divisor

import (
	"fmt"
	"testing"
)

func TestOnly3Divisors(t *testing.T) {
	t.Run("All Divisors", func(t *testing.T) {
		s := Only3_Divisors(50)
		fmt.Println(s)
	})
}
