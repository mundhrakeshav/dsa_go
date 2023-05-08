package bitmanipulation_test

import (
	bit_manipulation "mundhrakeshav/dsa/gfg/bit_manipulation"
	"testing"
)

func TestBitDiff(t *testing.T) {

	t.Run("BitDiffTest", func(t *testing.T) {
		arr := []uint8{1, 2}
		
		if bit_manipulation.BitDiff(arr) != 4 {
			t.Error("Invalid")
		}
	})
	
	t.Run("BitDiffTest", func(t *testing.T) {
		arr := []uint8{1, 2}
		
		if bit_manipulation.BitDiffEfficient(arr) != 4 {
			t.Error("Invalid")
		}
	})
	
	t.Run("BitDiffTest", func(t *testing.T) {
		arr := []uint8{1, 3, 5}
		
		if bit_manipulation.BitDiff(arr) != 8 {
			t.Error("Invalid")
		}
	})
	
	t.Run("BitDiffTest", func(t *testing.T) {
		arr := []uint8{1, 3, 5}
		
		if bit_manipulation.BitDiffEfficient(arr) != 8 {
			t.Error("Invalid")
		}
	})
}
