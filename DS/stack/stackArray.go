package stack

func CreateNew[T any](_initialData ...T) []T {
	return []T(_initialData)
}

// stackPush push to first index of array
func StackPush[T any](stack []T, v T) []T {
	return append([]T{v}, stack...)
}

// stackLength return length of array
func StackLength[T any](stack []T) int {
	return len(stack)
}

// stackPeak return last input of array
func  StackPeek[T any](stack []T) T {
	return stack[0]
}

// stackEmpty check array is empty or not
func StackEmpty[T any](stack []T) bool {
	return len(stack) == 0
}

// stackPop return last input and remove it in array
func StackPop[T any](stack []T) (T, []T) {
	pop := stack[0]
	stack = stack[1:]
	return pop, stack
}
