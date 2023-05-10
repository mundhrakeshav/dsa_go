package stacknqueue_test

import (
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"reflect"
	"testing"
)

func TestMergeOverlapping(t *testing.T) {
	tests := []struct {
		input  []stack_n_queue.Interval
		output []stack_n_queue.Interval
	}{
		{[]stack_n_queue.Interval{{1, 3}, {2, 4}, {6, 8}, {9, 10}}, []stack_n_queue.Interval{{1, 4}, {6, 8}, {9, 10}}},
		{[]stack_n_queue.Interval{{6,8},{1,9},{2,4},{4,7}}, []stack_n_queue.Interval{{1, 9}}},
		{[]stack_n_queue.Interval{{1, 3}, {4, 6}, {7, 9}}, []stack_n_queue.Interval{{1, 3}, {4, 6}, {7, 9}}},
		{[]stack_n_queue.Interval{}, []stack_n_queue.Interval{}},
		{[]stack_n_queue.Interval{{1, 5}}, []stack_n_queue.Interval{{1, 5}}},
	}
	for _, test := range tests {
		result := stack_n_queue.MergeOverlapping(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Merge Overlapping(%q) = %v, expected %v", test.input, result, test.output)
		}
	}

}
