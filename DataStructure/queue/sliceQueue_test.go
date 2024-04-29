package queue

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestSliceStackPushAndPop(t *testing.T) {

	s := NewSliceQueue[int]()

	s.Push(1)
	s.Push(2)

	assert.Equal(t, 1, s.Pop())
	assert.Equal(t, 2, s.Pop())
	assert.Equal(t, 0, s.Pop())
}
