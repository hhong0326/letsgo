package singlelinkedlist

type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{Value: value}
	l.count++

	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{Value: value}
	l.count++

	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	node.next = l.root
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {

	node := l.root
	count := 0
	for ; node != nil; node = node.next {
		count++
	}

	return count
}

func (l *LinkedList[T]) Count2() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {

	if idx >= l.Count2() {
		return nil
	}

	i := 0
	for node := l.root; node != nil; node = node.next {
		if i == idx {
			return node
		}
		i++
	}

	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {

	if !l.isIncluded(node) {
		return
	}

	newNode := &Node[T]{Value: value}
	l.count++

	newNode.next, node.next = node.next, newNode
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {

	if node == l.root {
		l.PushFront(value)
		return
	}

	prev := l.findPrevNode(node)
	if prev == nil {
		return
	}

	newNode := &Node[T]{Value: value}
	l.count++

	prev.next, newNode.next = newNode, node
}

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {

	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner.next == node {
			return inner
		}
	}

	return nil
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) PopFront() *Node[T] {

	if l.root == nil {
		return nil
	}

	n := l.root
	l.root.next, l.root = nil, l.root.next
	if l.root == nil {
		l.tail = nil
	}

	l.count--

	return n
}

// TODO: implement this
func (l *LinkedList[T]) PopBack() {

}

func (l *LinkedList[T]) Remove(node *Node[T]) {

	if node == l.root {
		l.PopFront()
		return
	}

	prev := l.findPrevNode(node)
	if prev == nil {
		return
	}

	prev.next, node.next = node.next, nil
	if node == l.tail {
		l.tail = prev
	}
	l.count--
}

func (l *LinkedList[T]) Reverse() {

	newLinkedList := &LinkedList[T]{}

	for l.root != nil {
		newLinkedList.PushFront(l.PopFront().Value)
	}

	l.count = newLinkedList.count
	l.root = newLinkedList.root
	l.tail = newLinkedList.tail
}

func (l *LinkedList[T]) Reverse2() {

	if l.root == nil {
		return
	}

	node := l.root
	next := l.root.next
	l.root.next = nil // root be tail

	for next != nil {
		nextnext := next.next
		next.next = node

		node = next
		next = nextnext
	}

	l.root, l.tail = l.tail, l.root
}
