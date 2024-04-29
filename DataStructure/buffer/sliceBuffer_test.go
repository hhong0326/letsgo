package buffer

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestWrite(t *testing.T) {
	b := NewSliceBuffer[byte]()
	b.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, b.Readable())
}

func TestRead(t *testing.T) {

	b := NewSliceBuffer[byte]()
	b.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, b.Readable())

	// 읽기 버퍼
	readedData := b.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readedData[i])
	}

	assert.Equal(t, 0, b.Readable())
}

func TestWriteAndRead(t *testing.T) {
	b := NewSliceBuffer[byte]()
	b.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, b.Readable())

	b.Write([]byte{5, 6})
	assert.Equal(t, 6, b.Readable())

	// 읽기 버퍼
	readedData := b.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readedData[i])
	}

	assert.Equal(t, 2, b.Readable())

	b.Write([]byte{7, 8, 9})
	assert.Equal(t, 5, b.Readable())

	readedData = b.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+5), readedData[i])
	}

	assert.Equal(t, 1, b.Readable())
}
