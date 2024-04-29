package stack

import "doublelinkedlist"

type Stack[T any] struct {
	l *doublelinkedlist.LinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		l: &doublelinkedlist.LinkedList[T]{},
	}
}

func (s *Stack[T]) Push(value T) {

	s.l.PushBack(value)
}

func (s *Stack[T]) Pop() *doublelinkedlist.Node[T] {

	if s.l.Count() == 0 {
		return nil
	}

	node := s.l.PopBack()

	return node
}
