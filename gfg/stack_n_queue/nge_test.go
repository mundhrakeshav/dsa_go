package stacknqueue_test

import (
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"reflect"
	"testing"
)

func TestNGE(t *testing.T) {
	tests := []struct {
		input  []uint
		output []int
	}{
		{[]uint{4, 5, 2, 25}, []int{5, 25, 25, -1}},
		{[]uint{13, 7, 6, 12}, []int{-1, 12, 12, -1}}, {[]uint{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, -1}},
		{[]uint{5, 4, 3, 2, 1}, []int{-1, -1, -1, -1, -1}},
		{[]uint{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, -1}},
		{[]uint{}, []int{}},                         // empty array
		{[]uint{1}, []int{-1}},                      // single element array
		{[]uint{1, 1, 1, 1}, []int{-1, -1, -1, -1}}, // array with repeated elements
		{[]uint{1, 2, 3, 4, 5, 4, 3, 2, 1}, []int{2, 3, 4, 5, -1, -1, -1, -1, -1}}, // array with multiple peaks

	}
	for _, test := range tests {
		result := stack_n_queue.NGE(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Next Greatest Element(%q) = %v, expected %v", test.input, result, test.output)
		}
	}

}
func TestNGEFwd(t *testing.T) {
	tests := []struct {
		input  []uint
		output map[uint]int
	}{
		{[]uint{4, 5, 2, 25}, map[uint]int{4: 5, 5: 25, 2: 25, 25: -1}},
		{[]uint{13, 7, 6, 12}, map[uint]int{13: -1, 7: 12, 6: 12, 12: -1}},
		{[]uint{1, 2, 3, 4, 5}, map[uint]int{1: 2, 2: 3, 3: 4, 4: 5, 5: -1}},
		{[]uint{5, 4, 3, 2, 1}, map[uint]int{1: -1, 2: -1, 3: -1, 4: -1, 5: -1}},
		{[]uint{}, map[uint]int{}},                // empty array
		{[]uint{1}, map[uint]int{1: -1}},          // single element array
		{[]uint{1, 1, 1, 1}, map[uint]int{1: -1}}, // array with repeated elements
	}
	for _, test := range tests {
		result := stack_n_queue.NGEFwd(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Next Greatest Element(%q) = %v, expected %v", test.input, result, test.output)
		}
	}

}
