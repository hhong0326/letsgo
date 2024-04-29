package stack

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestPushAndPop(t *testing.T) {
	s := New[int]()
	s.Push(1)
	s.Push(2)

	assert.Equal(t, 2, s.Pop().Value)
	assert.Equal(t, 1, s.Pop().Value)
}

func BenchmarkStack(b *testing.B) {

	s := New[int]()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func BenchmarkSliceStack(b *testing.B) {

	s := NewSliceStack[int]()

	for i := 0; i < b.N; i++ {
		s.Push(i)
	}

	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
