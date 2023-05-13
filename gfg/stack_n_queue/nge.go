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
				nge[i] = int(v)
				break
			}
		}
	}
	return nge
}

func NGEFwd(arr []uint) map[uint]int {
	stk := stack.CreateNew[uint]()
	nge := make(map[uint]int)
	for i := 0; i < len(arr); i++ {
		if stk.IsEmpty() {
			stk = stk.Push(arr[i])
		} else  {
			for !stk.IsEmpty() && stk.StackPeek() < arr[i]  {
				var pop uint;
				pop, stk = stk.Pop()
				nge[pop] = int(arr[i])
			}
			stk = stk.Push(arr[i])
		}
	}
	for !stk.IsEmpty() {
		var pop uint
		pop, stk = stk.Pop()
		nge[pop] = -1
	}
	return nge
}
