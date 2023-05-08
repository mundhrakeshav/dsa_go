package bitmanipulation_test

import (
	bit_manipulation "mundhrakeshav/dsa/gfg/bit_manipulation"
	"testing"
)

// Incomplete
func TestKthBitSet(t *testing.T) {
	
	t.Run("IsKthBitSet", func(t *testing.T) {
		if !bit_manipulation.IsKthBitSet(5, 0) {
			t.Error("0th Bit is set")
		}
		if bit_manipulation.IsKthBitSet(5, 1) {
			t.Error("1st Bit is unset")
		}
		if !bit_manipulation.IsKthBitSet(5, 2) {
			t.Error("2nd Bit is set")
		}
	})
}
