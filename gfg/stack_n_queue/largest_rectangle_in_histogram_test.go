package stacknqueue_test

import (
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"reflect"
	"testing"
)

func TestLargestRectangleInHistogram(t *testing.T) {
	tests := []struct {
		input  []uint
		output uint
	}{
		{[]uint{2, 1, 5, 6, 2, 3, 1}, 10},
		{[]uint{6, 2, 5, 4, 5, 1, 6}, 12},
		{[]uint{4, 5, 2, 10, 8}, 16},
		{[]uint{12}, 12},        // single element array
		{[]uint{1, 1, 1, 1}, 4}, // array with repeated elements
	}

	for _, test := range tests {
		result := stack_n_queue.LargestRectangleInHistogram(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Next Greatest Element(%q) = %v, expected %v", test.input, result, test.output)
		}
	}

}
