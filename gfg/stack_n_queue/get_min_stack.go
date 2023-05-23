package stacknqueue

import (
	stack "mundhrakeshav/dsa/DS/stack"
	types "mundhrakeshav/dsa/types"
)

type GetMinStack[T types.Number] struct {
	Stk stack.Stack[T]
	Min T
}

func CreateNewGetMinStack[T types.Number]() GetMinStack[T] {
	s1 := stack.CreateNew[T]()

	return GetMinStack[T]{
		Stk: s1,
	}

}

func (minStk GetMinStack[T]) Push(t T) GetMinStack[T] {
	if minStk.Stk.IsEmpty() {
		minStk.Stk = minStk.Stk.Push(t)
		minStk.Min = t
		return minStk
	}
	if minStk.Min < t {
		minStk.Stk = minStk.Stk.Push(t)
	} else {
		minStk.Stk = minStk.Stk.Push((2 * t) - minStk.Min)
		minStk.Min = t
	}
	return minStk
}
func (minStk GetMinStack[T]) Pop() (T, GetMinStack[T]) {
	if minStk.Stk.IsEmpty() {
		var pop T
		return pop, minStk
	}
	var pop T

	pop, minStk.Stk = minStk.Stk.Pop()
	if pop < minStk.Min {
		newMin := (2 * minStk.Min) - (pop)
		pop = minStk.Min
		minStk.Min = newMin
	}
	return pop, minStk
}
