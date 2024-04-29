package doublelinkedlist

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {

	if idx >= l.Count() {
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

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{Value: value}

	if l.tail == nil {
		l.root = node
		l.tail = node
		l.count = 1

		return
	}

	l.tail.next = node
	node.prev = l.tail
	l.tail = node

	l.count++
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{Value: value}

	if l.root == nil {
		l.root = node
		l.tail = node
		l.count = 1

		return
	}

	l.root.prev = node
	node.next = l.root
	l.root = node

	l.count++
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

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {

	inner := l.root

	for ; inner != nil; inner = inner.next {
		if inner.next == node {
			return inner
		}
	}

	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {

	if !l.isIncluded(node) {
		return
	}

	newNode := &Node[T]{Value: value}

	nextNode := node.next
	node.next = newNode
	newNode.prev = node
	newNode.next = nextNode

	if nextNode != nil {
		nextNode.prev = newNode
	}

	if node == l.tail {
		l.tail = newNode
	}

	l.count++
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {

	prev := l.findPrevNode(node)
	if prev == nil {
		return
	}

	newNode := &Node[T]{Value: value}

	prevNode := node.prev
	node.prev = newNode
	newNode.next = node
	newNode.prev = prevNode

	if prevNode == nil {
		prevNode.next = newNode
	}

	if node == l.root {
		l.root = newNode
	}

	l.count++
}

func (l *LinkedList[T]) PopFront() *Node[T] {

	if l.root == nil {
		return nil
	}

	node := l.root

	l.root, l.root.prev = node.next, nil
	node.next = nil

	if l.root == nil {
		l.tail = nil
	}

	l.count--

	return node
}

func (l *LinkedList[T]) PopBack() *Node[T] {

	if l.tail == nil {
		return nil
	}

	node := l.tail

	l.tail, l.tail.next = l.tail.prev, nil

	if l.tail == nil {
		l.root = nil
	}

	l.count--
	return node
}

func (l *LinkedList[T]) Remove(node *Node[T]) {

	if !l.isIncluded(node) {
		return
	}

	if node == l.root {
		l.PopFront()
		l.count--
		return
	}

	if node == l.tail {

		l.PopBack()
		l.count--
		return
	}

	prev := l.findPrevNode(node)

	prev.next = node.next
	node.next.prev = prev

	node.next = nil
	node.prev = nil

	l.count--
}

func (l *LinkedList[T]) Reverse() {

	if l.root == nil {
		return
	}

	node := l.root
	var next *Node[T]

	for node != nil {
		next = node.next

		node.prev, node.next = node.next, node.prev

		node = next
	}

	l.root, l.tail = l.tail, l.root
}
