package stack_test

import (
	stack "mundhrakeshav/dsa/DS/stack"
	"reflect"
	"testing"
)

type CreateNewTest[T any] struct {
	input           []T
	expected_length int
	name            string
}

type Person struct {
	Name string
	Age  int
}

// TestStackArray for testing Stack with Array
func TestStackArray(t *testing.T) {
	t.Run("Create New", func(t *testing.T) {
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

		for _, _tests := range createNewTests {
			new_stack := stack.CreateNew(_tests.input...)
			if len(new_stack) != _tests.expected_length {
				t.Errorf("Expected stack length to be %d, but got %d", _tests.expected_length, len(new_stack))
			}
			for i, val := range _tests.input {
				if !reflect.DeepEqual(new_stack[i], val) {
					t.Errorf("Expected stack[%d] to be %d, but got %d", i, val, new_stack[i])
				}
			}
		}

		struct_data := []Person{
			{Name: "Alice", Age: 25},
			{Name: "Bob", Age: 30},
		}
		stack_struct := stack.CreateNew(struct_data...)
		if len(stack_struct) != len(struct_data) {
			t.Errorf("Expected stack length to be %d, but got %d", len(struct_data), len(stack_struct))
		}
		for i, val := range struct_data {
			if !reflect.DeepEqual(stack_struct[i], val) {
				t.Errorf("Expected stack[%d] to be %v, but got %v", i, val, stack_struct[i])
			}
		}

	})

	t.Run("Push", func(t *testing.T) {
		stk := stack.CreateNew(1, 3, 5, 6, 7, 8, 9, 12, 10, 2, 43)
		stk_len := len(stk)
		insrt := 100
		stk = stk.Push(insrt)
		if len(stk) != stk_len+1 {
			t.Errorf("Expected stack length to be %d, but got %d", stk_len+1, len(stk))
		}
		if stk[0] != insrt {
			t.Errorf("Expected 0th element to be %d but got %d", insrt, stk[0])
		}

		/*------------------------------------------------------------------------------------*/

		// Test case: Pushing to an empty stack
		empty_stack := stack.CreateNew([]int{}...)
		empty_stack = empty_stack.Push(1)
		if len(empty_stack) != 1 {
			t.Errorf("Expected stack length to be 1, but got %d", len(empty_stack))
		}
		if empty_stack[0] != 1 {
			t.Errorf("Expected stack[0] to be 1, but got %v", empty_stack[0])
		}

		/*------------------------------------------------------------------------------------*/
		// Test case: Pushing to a full stack
		full_stack := stack.CreateNew([]int{1, 2, 3}...)
		full_stack = full_stack.Push(4)
		if len(full_stack) != 4 {
			t.Errorf("Expected stack length to be 4, but got %d", len(full_stack))
		}
		/*------------------------------------------------------------------------------------*/

		// Test case: Pushing a large amount of data
		large_data := make([]int, 10000)
		stk = stack.CreateNew([]int{1, 2, 3}...)
		for _, val := range large_data {
			stk = stk.Push(val)
		}
		if len(stk) != 10003 {
			t.Errorf("Expected stack length to be 1000003, but got %d", len(stk))
		}

		/*------------------------------------------------------------------------------------*/

		// Test case: Pushing a custom struct
		type Person struct {
			Name string
			Age  int
		}
		person := Person{Name: "Alice", Age: 25}
		struct_stk := stack.CreateNew([]Person{}...)
		struct_stk = struct_stk.Push(person)
		if len(struct_stk) != 1 {
			t.Errorf("Expected stack length to be 1, but got %d", len(struct_stk))
		}
		if struct_stk[0] != person {
			t.Errorf("Expected stack[0] to be %v, but got %v", person, struct_stk[0])
		}

	})
	t.Run("Pop", func(t *testing.T) {
		stk := stack.CreateNew(1, 2, 3, 4, 5)
		stk = stk.Pop()
		if len(stk) != 4 {
			t.Errorf("Expected stack length to be 4, but got %d", len(stk))
		}
		if stk[0] != 2 {
			t.Errorf("Expected top element to be 2, but got %d", stk[0])
		}
		/*------------------------------------------------------------------------------------*/

		// Pop another element off the stack
		stk = stk.Pop()

		// Check that the top element was removed
		if len(stk) != 3 {
			t.Errorf("Expected stack length to be 3, but got %d", len(stk))
		}
		if stk[0] != 3 {
			t.Errorf("Expected top element to be 3, but got %d", stk[0])
		}
		stk = stk.Pop().Pop().Pop()

		// Check that the stack is now empty
		if len(stk) != 0 {
			t.Errorf("Expected stack length to be 0, but got %d", len(stk))
		}

		// Should not throw if no elements are present
		stk.Pop().Pop().Pop()
		
		// Check that the stack is now empty
		if len(stk) != 0 {
			t.Errorf("Expected stack length to be 0, but got %d", len(stk))
		}
	})
}
