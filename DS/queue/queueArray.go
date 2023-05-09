package queue

type Queue[T any] []T


func CreateNew[T any](_initialData ...T) Queue[T] {
	return Queue[T](_initialData)
}

func (_queue Queue[T]) Enqueue( _data T) []T {
	return append(_queue, _data)
}

func (_queue Queue[T]) Dequeue() []T {
	return _queue[1:]
}