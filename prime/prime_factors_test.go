package misc

import (
	"fmt"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	t.Run("Prime Factor", func(t *testing.T) {
		s := Prime_factors(1371)
		fmt.Println(s)

	})
}
