package stacknqueue

import (
	stack "mundhrakeshav/dsa/DS/stack"
)

type QueueUsingStack[T any] struct {
	S1 stack.Stack[T]
	S2 stack.Stack[T]
}

func CreateNew[T any](initial_data ...T) QueueUsingStack[T] {
	s1 := stack.CreateNew(initial_data...)
	s2 := stack.CreateNew[T]()

	return QueueUsingStack[T]{
		S1: s1, S2: s2,
	}

}

func (stk QueueUsingStack[T]) Top(_data T) T {
	if !stk.S2.IsEmpty() {
		return stk.S2.StackPeek()
	} else if !stk.S1.IsEmpty() {
		for !stk.S1.IsEmpty() {
			pop, _stk := stk.S1.Pop()
			stk.S2.Push(pop)
			stk.S1 = _stk
		}
		return stk.S1.StackPeek()
	} else {
		var t T
		return t
	}
}

func (stk QueueUsingStack[T]) Push(_data T) QueueUsingStack[T] {
	stk.S1 = stk.S1.Push(_data)
	return stk
}

func (stk QueueUsingStack[T]) Pop() (T, QueueUsingStack[T]) {
	if !stk.S2.IsEmpty() {
		pop, _stk := stk.S2.Pop()
		stk.S2 = _stk
		return pop, stk
	} else if !stk.S1.IsEmpty() {
		for !stk.S1.IsEmpty() {
			pop, _stk := stk.S1.Pop()
			stk.S2 = stk.S2.Push(pop)
			stk.S1 = _stk
		}
		pop, _stk := stk.S2.Pop()
		stk.S2 = _stk
		return pop, stk
	} else {
		var t T
		return t, stk
	}
}
