package queue

import "doublelinkedlist"

type Queue[T any] struct {
	l *doublelinkedlist.LinkedList[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		l: &doublelinkedlist.LinkedList[T]{},
	}
}

func (q *Queue[T]) Push(value T) {

	q.l.PushBack(value)
}

func (q *Queue[T]) Pop() *doublelinkedlist.Node[T] {

	if q.l.Count() == 0 {
		return nil
	}

	node := q.l.PopFront()

	return node
}
