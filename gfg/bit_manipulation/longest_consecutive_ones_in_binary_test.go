package bitmanipulation_test

import (
	bit_manipulation "mundhrakeshav/dsa/gfg/bit_manipulation"
	"testing"
)

func TestLongestConsecutiveOnesInBinary(t *testing.T) {
	testCases := []struct {
		input    uint
		expected uint
	}{
		{0b0, 0},
		{0b1, 1},
		{0b101, 1},
		{0b111, 3},
		{0b110111011, 3},
		{0b10000000000000000000000000000000, 1},
		{0b11111111111111111111111111111111, 32},
		{0xFFFFFFFF, 32}, // maximum value for a 32-bit unsigned integer
		{0x80000000, 1},  // largest power of 2
		{0x00000001, 1},  // smallest power of 2
		{0x55555555, 1},  // alternating 1s and 0s
		{0xAAAAAAAA, 1},  // alternating 0s and 1s
		{0x00000000, 0},  // zero
	}

	for _, tc := range testCases {
		result := bit_manipulation.LongestConsecutiveOnesInBinary(tc.input)
		if result != tc.expected {
			t.Errorf("Expected LongestConsecutiveOnesInBinary(%b) to be %d, but got %d", tc.input, tc.expected, result)
		} else {
			t.Log(tc.input, tc.expected, result)
		}
	}
}
