package bitmanipulation_test

import (
	bit_manipulation "mundhrakeshav/dsa/gfg/bit_manipulation"
	"testing"
)

// Incomplete
func TestRMSB(t *testing.T) {
	t.Run("RMSB", func(t *testing.T) {
		s := bit_manipulation.GetRightMostSetBit(19)
		if s != 1 {
			t.Error("Invalid")
		}
	})
	t.Run("RMSB", func(t *testing.T) {
		s := bit_manipulation.GetRightMostSetBit(18)
		if s != 2 {
			t.Error("Invalid")
		}
	})
}
