package doublelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {

	l := LinkedList[int]{}
	// var l LinkedList[int] // same as above

	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())

	l.PushBack(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	assert.Equal(t, 2, l.Count())

	assert.Equal(t, 2, l.GetAt(1).Value)
}

func TestPushFront(t *testing.T) {

	l := LinkedList[int]{}
	// var l LinkedList[int] // same as above

	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())

	l.PushFront(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(2)

	assert.NotNil(t, l.root)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	assert.Equal(t, 2, l.Count())

	assert.Equal(t, 2, l.GetAt(0).Value)
	assert.Nil(t, l.GetAt(4))
}

func TestInsertAfter(t *testing.T) {

	l := LinkedList[int]{} // var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertAfter(node, 4)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)

	l.InsertAfter(l.Back(), 10)

	tempNode := &Node[int]{
		Value: 200,
	}

	l.InsertAfter(tempNode, 100)
	assert.Equal(t, 5, l.Count())

	assert.Equal(t, 4, node.next.Value)
	assert.Equal(t, 3, node.next.next.Value)

}

func TestInsertBefore(t *testing.T) {

	l := LinkedList[int]{} // var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertBefore(node, 4)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 2, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)

	l.InsertBefore(l.Front(), 10)

	tempNode := &Node[int]{
		Value: 200,
	}

	l.InsertBefore(tempNode, 100)
	assert.Equal(t, 5, l.Count())
}

func TestPopFront(t *testing.T) {

	l := LinkedList[int]{} // var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	node := l.PopFront()
	assert.Equal(t, 1, node.Value)

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Front().Value)

	l.PopFront()
	l.PopFront()

	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
}

func TestPopBack(t *testing.T) {

	l := LinkedList[int]{} // var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PopBack()

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Back().Value)

	l.PopBack()
	l.PopBack()

	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
}

func TestRemove(t *testing.T) {

	l := LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Remove(l.GetAt(1))

	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 3, l.GetAt(1).Value)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
}

func TestReverse(t *testing.T) {
	l := LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Reverse()
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
}
