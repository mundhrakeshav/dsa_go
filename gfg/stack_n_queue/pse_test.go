package stacknqueue_test

import (
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"reflect"
	"testing"
)

func TestPSE(t *testing.T) {
	tests := []struct {
		input  []uint
		output []int
	}{
		{[]uint{2, 1, 5, 6, 2, 3, 1}, []int{-1, -1, 1, 2, 1, 4, -1}},
		{[]uint{6, 2, 5, 4, 5, 1, 6}, []int{-1, -1, 1, 1, 3, -1, 5}},
		{[]uint{3, 4, 1, 5, 6, 2, 7}, []int{-1, 0, -1, 2, 3, 2, 5}},
		{[]uint{}, []int{}},                         // empty array
		{[]uint{1}, []int{-1}},                      // single element array
		{[]uint{1, 1, 1, 1}, []int{-1, -1, -1, -1}}, // array with repeated elements
		{[]uint{1, 2, 3, 4, 5, 4, 3, 2, 1}, []int{-1, 0, 1, 2, 3, 2, 1, 0, -1}}, // array with multiple peaks

	}
	for _, test := range tests {
		result := stack_n_queue.PSE(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Next Greatest Element(%q) = %v, expected %v", test.input, result, test.output)
		}
	}
}
