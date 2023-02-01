package queue


func createNew[T any](_initialData ...T) []T {
	return []T(_initialData)
}

func enqueue[T any](_queue []T, _data T) []T {
	return append(_queue, _data)
}

func dequeue[T any](_queue []T) []T  {
	return _queue[1:]
}