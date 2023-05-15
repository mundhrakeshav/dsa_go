package stacknqueue

import stack "mundhrakeshav/dsa/DS/stack"

// next small element
func NSE(arr []uint)[]int {

	stk := stack.CreateNew[uint]()
	nse := make([]int, len(arr))
	
	for i := len(arr) - 1; i >= 0; i-- {
		for {
			if stk.IsEmpty() {
				stk = stk.Push(uint(i))
				nse[i] = -1
				break
			} else if peek := stk.StackPeek(); arr[peek] < arr[i] {
				stk = stk.Push(uint(i))
				nse[i] = int(peek)
				break
			} else {
				_, stk = stk.Pop()
			}
		}
	}
	return nse;
}