package singlelinkedlist

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
	assert.Equal(t, 2, l.Count2())

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

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)

	tempNode := &Node[int]{
		Value: 200,
	}

	l.InsertAfter(tempNode, 100)
	assert.Equal(t, 4, l.Count2())

}

func TestInsertBefore(t *testing.T) {

	l := LinkedList[int]{} // var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertBefore(node, 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 2, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)

	tempNode := &Node[int]{
		Value: 200,
	}

	l.InsertBefore(tempNode, 100)
	assert.Equal(t, 4, l.Count2())
}

func TestInsertBeforeRoot(t *testing.T) {

	l := LinkedList[int]{} // var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(0)
	l.InsertBefore(node, 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(0).Value)
	assert.Equal(t, 1, l.GetAt(1).Value)

	tempNode := &Node[int]{
		Value: 200,
	}

	l.InsertBefore(tempNode, 100)
	assert.Equal(t, 4, l.Count2())

}

func TestPopFront(t *testing.T) {

	l := LinkedList[int]{} // var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PopFront()

	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 2, l.Front().Value)

	l.PopFront()
	l.PopFront()

	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
}

func TestRemove(t *testing.T) {

	l := LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Remove(l.GetAt(1))
	assert.Equal(t, 2, l.Count2())
	assert.Equal(t, 1, l.Front().Value)
	l.Remove(l.GetAt(0))
	assert.Equal(t, 1, l.Count2())
	assert.Equal(t, 3, l.Front().Value)
	l.Remove(l.GetAt(0))
	assert.Equal(t, 0, l.Count2())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
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

func TestReverse2(t *testing.T) {
	l := LinkedList[int]{}
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.Reverse2()
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
}
