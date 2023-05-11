package stacknqueue_test

import (
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"reflect"
	"testing"
)

func TestStockSpan(t *testing.T) {
	tests := []struct {
		input  []uint
		output []uint
	}{
		{[]uint{100, 80, 60, 70, 60, 75, 85}, []uint{1, 1, 1, 2, 1, 4, 6}},
		{[]uint{10, 4, 5, 90, 120, 80}, []uint{1, 1, 2, 4, 5, 1}},
		{[]uint{10, 20, 30, 40, 50}, []uint{1, 2, 3, 4, 5}},
		{[]uint{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []uint{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}},
	}
	for _, test := range tests {
		result := stack_n_queue.StockSpan(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("StockSpan(%q) = %v, expected %v", test.input, result, test.output)
		}
	}

}
