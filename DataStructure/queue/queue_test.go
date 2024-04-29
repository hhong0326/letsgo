package queue

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPushAndPop(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)

	assert.Equal(t, 1, q.Pop().Value)
	assert.Equal(t, 2, q.Pop().Value)
}

func BenchmarkQueue(b *testing.B) {

	q := New[int]()

	for i := 0; i < b.N; i++ {
		q.Push(i)
	}

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}

func BenchmarkSliceQueue(b *testing.B) {

	q := NewSliceQueue[int]()

	for i := 0; i < b.N; i++ {
		q.Push(i)
	}

	for i := 0; i < b.N; i++ {
		q.Pop()
	}
}
