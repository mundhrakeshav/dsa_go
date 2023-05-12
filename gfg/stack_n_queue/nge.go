package stacknqueue

import (
	stack "mundhrakeshav/dsa/DS/stack"
)

func NGE(arr []uint) []int {
	stk := stack.CreateNew[uint]()
	nge := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		for {
			if stk.IsEmpty() {
				stk = stk.Push(arr[i])
				nge[i] = -1
				break
			} else if v := stk.StackPeek(); v <= arr[i] {
				_, stk = stk.Pop()
			} else {
				stk = stk.Push(arr[i])
				nge[i] = int(v);
				break
			}
		}
	}
	return nge;
}
