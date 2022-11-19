package gcd

import (
	"fmt"
	"testing"
)

func TestGCDDivision(t *testing.T) {

	t.Run("Should return GCD", func(t *testing.T) {
		var a, b uint32 = 24, 42
		gcd := GCD_Division_Iterative(a, b)
		fmt.Println(gcd)
	})
	
	t.Run("Should return GCD", func(t *testing.T) {
		var a, b uint32 = 81, 45
		gcd := GCD_Division_Recursive(a, b)
		fmt.Println(gcd)
	})
}
