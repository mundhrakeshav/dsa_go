package stack

func createNew[T any](_initialData ...T) []T {
	return []T(_initialData)
}

// stackPush push to first index of array
func stackPush[T any](stack []T, v T) []T {
	return append([]T{v}, stack...)
}

// stackLength return length of array
func stackLength[T any](stack []T) int {
	return len(stack)
}

// stackPeak return last input of array
func  stackPeak[T any](stack []T) any {
	return stack[0]
}

// stackEmpty check array is empty or not
func stackEmpty[T any](stack []T) bool {
	return len(stack) == 0
}

// stackPop return last input and remove it in array
func stackPop[T any](stack []T) (T, []T) {
	pop := stack[0]
	stack = stack[1:]
	return pop, stack
}
