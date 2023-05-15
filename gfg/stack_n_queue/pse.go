package stacknqueue

import (
	stack "mundhrakeshav/dsa/DS/stack"
)

// Prev small element
func PSE(arr []uint) []int {

	stk := stack.CreateNew[uint]()
	pse := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		for {
			if stk.IsEmpty() {
				stk = stk.Push(uint(i))
				pse[i] = -1
				break
			} else if peek := stk.StackPeek(); arr[peek] < arr[i] {
				stk = stk.Push(uint(i))
				pse[i] = int(peek)
				break
			} else {
				_, stk = stk.Pop()
			}
		}
	}
	return pse;
}
