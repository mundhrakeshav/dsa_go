package stacknqueue_test

import (
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"testing"
)

func TestFirstCirularTour(t *testing.T) {
	tests := []struct {
		arr      []stack_n_queue.FirstCircularTour_FuelDist
		expected int
	}{
		{
			arr: []stack_n_queue.FirstCircularTour_FuelDist{
				{
					Fuel:     4,
					Distance: 6,
				},
				{
					Fuel:     6,
					Distance: 5,
				},
				{
					Fuel:     7,
					Distance: 3,
				},
				{
					Fuel:     4,
					Distance: 5,
				},
			},
			expected: 1,
		},
		{
			arr: []stack_n_queue.FirstCircularTour_FuelDist{
				{
					Fuel:     6,
					Distance: 4,
				},
				{
					Fuel:     3,
					Distance: 6,
				},
				{
					Fuel:     7,
					Distance: 3,
				},
				{
					Fuel:     4,
					Distance: 5,
				},
			},
			expected: 2,
		},
		{
			arr: []stack_n_queue.FirstCircularTour_FuelDist{
				{
					Fuel:     7,
					Distance: 6,
				},
				{
					Fuel:     8,
					Distance: 7,
				},
				{
					Fuel:     5,
					Distance: 8,
				},
				{
					Fuel:     11,
					Distance: 9,
				},
				{
					Fuel:     7,
					Distance: 7,
				},
				{
					Fuel:     6,
					Distance: 5,
				},
			},
			expected: 3,
		},
	}

	for _, test := range tests {
		result := stack_n_queue.FirstCircularTour(test.arr)
		if result != test.expected {
			t.Errorf("FirstCircularTour (%q) = %v, expected %v", test.arr, result, test.expected)
		}
	}
}