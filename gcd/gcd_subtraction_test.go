package gcd

import (
	"fmt"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestGCDSubtraction(t *testing.T) {

	t.Run("Should return GCD", func(t *testing.T) {
		var a, b uint32 = 9, 81
		gcd := GCD_Subtraction(a, b);
		fmt.Println(gcd)
		if gcd != 9 {
			t.Error("INVALID")	
		}

		
	})
}
