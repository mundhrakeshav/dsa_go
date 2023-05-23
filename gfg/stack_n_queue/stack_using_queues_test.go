package stacknqueue_test

import (
	stack_n_queue "mundhrakeshav/dsa/gfg/stack_n_queue"
	"reflect"
	"testing"
)

type CreateNewTest[T any] struct {
	input           []T
	expected_length int
	name            string
}

func TestStackUsingQueues(t *testing.T) {
	createNewTests := []CreateNewTest[int]{
		{
			input:           []int{1, 2, 3, 227002894, 68652442, 36672393},
			expected_length: 6,
			name:            "Basic Initialization",
		},
		{
			input:           []int{},
			expected_length: 0,
			name:            "Empty",
		},
		{
			input:           make([]int, 10000),
			expected_length: 10000,
			name:            "Large Data",
		},
		{
			input:           (nil),
			expected_length: 0,
			name:            "Large Data",
		},
	}

	t.Run("Create New", func(t *testing.T) {

		for _, _tests := range createNewTests {
			new_stack := stack_n_queue.CreateNew(_tests.input...)
			if len(new_stack.S1) != _tests.expected_length {
				t.Errorf("Expected stack length to be %d, but got %d", _tests.expected_length, len(new_stack.S1))
			}
			for i, val := range _tests.input {
				if !reflect.DeepEqual(new_stack.S1[i], val) {
					t.Errorf("Expected stack[%d] to be %d, but got %d", i, val, new_stack.S1[i])
				}
			}
		}
	})

	t.Run("Push and Pop", func(t *testing.T) {
		// -- 1, 3, 5, 6, 7, 8, 9, 12, 10, 2, 43 -->
		stk := stack_n_queue.CreateNew(1, 3, 5, 6, 7, 8, 9, 12, 10, 2, 43)
		stk_len := len(stk.S1)
		insrt := 100
		// --100, 1, 3, 5, 6, 7, 8, 9, 12, 10, 2, 43 -->
		stk = stk.Push(insrt)
		if len(stk.S1) != stk_len+1 {
			t.Errorf("Expected stack length to be %d, but got %d", stk_len+1, len(stk.S1))
		}
		if stk.S1[0] != insrt {
			t.Errorf("Expected 0th element to be %d but got %d", insrt, stk.S1[0])
		}
		var pop int;
		pop, stk = stk.Pop()
		
		if pop != 43 {
			t.Errorf("Expected pop element to be %d but got %d", 1, pop)
		}
		pop, stk = stk.Pop()
		if pop != 2 {
			t.Errorf("Expected pop element to be %d but got %d", 1, pop)
		}
		// /*------------------------------------------------------------------------------------*/

		// // Test case: Pushing to an empty stack
		empty_stack := stack_n_queue.CreateNew([]int{}...)
		empty_stack = empty_stack.Push(1)
		if len(empty_stack.S1) != 1 {
			t.Errorf("Expected stack length to be 1, but got %d", len(empty_stack.S1))
		}
		if empty_stack.S1[0] != 1 {
			t.Errorf("Expected stack[0] to be 1, but got %v", empty_stack.S1[0])
		}
		_, empty_stack = empty_stack.Pop()
		if len(empty_stack.S1) != 0 {
			t.Errorf("Expected stack length to be 1, but got %d", len(empty_stack.S1))
		}
		_, empty_stack = empty_stack.Pop()
		if len(empty_stack.S1) != 0 {
			t.Errorf("Expected stack length to be 1, but got %d", len(empty_stack.S1))
		}
	})

}
