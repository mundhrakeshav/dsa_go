package stack

type Stack[T any] []T


func CreateNew[T any](initial_data ...T) Stack[T] {
	return Stack[T](initial_data);
}

func (stk Stack[T]) Push(data T) Stack[T] {
	return append([]T{data}, stk...);
}

func (stk Stack[T]) Pop() Stack[T] {
	if (len(stk) == 0) {
		return stk
	}
	return stk[1:];
}


func StackLength[T any](stack []T) int {
	return len(stack)
}

func  StackPeek[T any](stack []T) T {
	return stack[0]
}

