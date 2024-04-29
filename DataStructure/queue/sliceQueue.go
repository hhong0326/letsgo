package queue

type SliceQueue[T any] struct {
	arr []T
}

func NewSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{
		arr: []T{},
	}
}
func (q *SliceQueue[T]) Push(value T) {
	q.arr = append(q.arr, value)
}

func (q *SliceQueue[T]) Pop() T {

	var first T

	if len(q.arr) == 0 {
		return first
	}

	first = q.arr[0]
	q.arr = q.arr[1:]

	return first
}
