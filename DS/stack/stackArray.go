package stack

type Stack[T any] []T


func CreateNew[T any](initial_data ...T) Stack[T] {
	return Stack[T](initial_data);
}

func (stk Stack[T]) Push(data T) Stack[T] {
	return append([]T{data}, stk...);
}

func (stk Stack[T]) Pop() (T, Stack[T]) {
	var pop T;
	if (len(stk) == 0) {
		return pop,stk
	}
	pop = stk[0]
	return pop, stk[1:];
}

func (stk Stack[T]) StackLength() int {
	return len(stk)
}

func (stk Stack[T]) StackPeek() T {

	return stk[0]
}

func (stk Stack[T]) IsEmpty() bool {
	return len(stk) == 0
}
