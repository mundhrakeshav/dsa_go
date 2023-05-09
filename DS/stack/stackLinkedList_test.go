package stack_test

import (
	stack "mundhrakeshav/dsa/DS/stack"

	// "reflect"
	"testing"
)

// TestStackArray for testing Stack with Array
func TestStackLinkedList(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	t.Run("Create New", func(t *testing.T) {
		val := 1
		ll := stack.CreateNewLinkedList(val)
		if ll.Head == nil {
			t.Errorf("Expected head to be not nil, but got nil")
		}
		if ll.Head.Val != val {
			t.Errorf("Expected head value to be %d, but got %d", val, ll.Head.Val)
		}
		if ll.Length != 1 {
			t.Errorf("Expected length to be 1, but got %d", ll.Length)
		}
	})

	t.Run("Create New Nil Value", func(t *testing.T) {
		var val *int
		ll := stack.CreateNewLinkedList(val)
		if ll.Head.Val != nil {
			t.Errorf("Expected ll to be nil, but got %v", ll)
		}
	})

	t.Run("Create New Struct Value", func(t *testing.T) {
		person := Person{"John", 30}
		ll := stack.CreateNewLinkedList(person)
		if ll.Head == nil {
			t.Errorf("Expected head to be not nil, but got nil")
		}
		if ll.Head.Val != person {
			t.Errorf("Expected head value to be %v, but got %v", person, ll.Head.Val)
		}
		if ll.Length != 1 {
			t.Errorf("Expected length to be 1, but got %d", ll.Length)
		}
	})

	t.Run("Create New Pointer Value", func(t *testing.T) {
		val := 1
		ptr := &val
		ll := stack.CreateNewLinkedList(ptr)
		if ll.Head == nil {
			t.Errorf("Expected head to be not nil, but got nil")
		}
		if ll.Head.Val != &val {
			t.Errorf("Expected head value to be %v, but got %v", &val, ll.Head.Val)
		}
		if ll.Length != 1 {
			t.Errorf("Expected length to be 1, but got %d", ll.Length)
		}
	})

	t.Run("Push: struct", func(t *testing.T) {
		person1 := Person{"John", 30}
		person2 := Person{"Jane", 25}
		ll := stack.CreateNewLinkedList(person1)
		ll.Push(person2)
		if ll.Head == nil {
			t.Errorf("Expected head to be not nil, but got nil")
		}
		if ll.Head.Val != person2 {
			t.Errorf("Expected head value to be %v, but got %v", person2, ll.Head.Val)
		}
		if ll.Length != 2 {
			t.Errorf("Expected length to be 2, but got %d", ll.Length)
		}
	})
	
	t.Run("Push", func(t *testing.T) {
		ll := &stack.LinkedListStack[int]{}
		ll.Push(1)
		ll.Push(2)
		ll.Push(3)
		if ll.Head == nil {
			t.Errorf("Expected head to be not nil, but got nil")
		}
		if ll.Head.Val != 3 {
			t.Errorf("Expected head value to be 3, but got %d", ll.Head.Val)
		}
		if ll.Length != 3 {
			t.Errorf("Expected length to be 3, but got %d", ll.Length)
		}
	})


	t.Run("Pop", func(t *testing.T) {
		ll := &stack.LinkedListStack[int]{}
		ll.Push(1)
		ll.Push(2)
		ll.Push(3)

		ll.Pop()
		if ll.Head == nil {
			t.Errorf("Expected head to be not nil, but got nil")
		}
		if ll.Head.Val != 2 {
			t.Errorf("Expected head value to be 2, but got %d", ll.Head.Val)
		}
		if ll.Length != 2 {
			t.Errorf("Expected length to be 2, but got %d", ll.Length)
		}
	
		ll.Pop()
		ll.Pop()
		if ll.Head != nil {
			t.Errorf("Expected head to be nil, but got %v", ll.Head)
		}
		if ll.Length != 0 {
			t.Errorf("Expected length to be 0, but got %d", ll.Length)
		}
	
		ll.Pop()
		if ll.Head != nil {
			t.Errorf("Expected head to be nil, but got %v", ll.Head)
		}
		if ll.Length != 0 {
			t.Errorf("Expected length to be 0, but got %d", ll.Length)
		}

	})
}
