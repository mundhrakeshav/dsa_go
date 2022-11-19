package misc

import (
	"fmt"
	"testing"
)

func TestSmallestPrimeFactors(t *testing.T) {
	t.Run("Prime Factor", func(t *testing.T) {
		s := Smallest_Prime_Factors(26)
		fmt.Println(s)

	})
}
