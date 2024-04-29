package buffer

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRingWrite(t *testing.T) {
	r := NewRingBuffer[byte](10)
	r.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, r.Readable())
}

func TestRingRead(t *testing.T) {

	r := NewRingBuffer[byte](10)
	r.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, r.Readable())

	// 읽기 버퍼
	readedData := r.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readedData[i])
	}

	assert.Equal(t, 0, r.Readable())
}

func TestRingOverWrite(t *testing.T) {
	r := NewRingBuffer[byte](5)
	r.Write([]byte{1, 2, 3, 4})
	assert.Equal(t, 4, r.Readable())
	assert.Equal(t, 1, r.Writable())

	r.Write([]byte{5})
	assert.Equal(t, 5, r.Readable())
	assert.Equal(t, 0, r.Writable())

	written := r.Write([]byte{6})
	assert.Equal(t, 0, written)
	assert.Equal(t, 5, r.Readable())
	assert.Equal(t, 0, r.Writable())

	readedData := r.Read(4)
	for i := 0; i < 4; i++ {
		assert.Equal(t, byte(i+1), readedData[i])
	}

	assert.Equal(t, 1, r.Readable())
	assert.Equal(t, 0, r.Writable())
}

// func TestRingWriteAndRead(t *testing.T) {
// 	r := NewSliceBuffer[byte]()
// 	r.Write([]byte{1, 2, 3, 4})
// 	assert.Equal(t, 4, r.Readable())

// 	r.Write([]byte{5, 6})
// 	assert.Equal(t, 6, r.Readable())

// 	// 읽기 버퍼
// 	readedData := r.Read(4)
// 	for i := 0; i < 4; i++ {
// 		assert.Equal(t, byte(i+1), readedData[i])
// 	}

// 	assert.Equal(t, 2, r.Readable())

// 	r.Write([]byte{7, 8, 9})
// 	assert.Equal(t, 5, r.Readable())

// 	readedData = r.Read(4)
// 	for i := 0; i < 4; i++ {
// 		assert.Equal(t, byte(i+5), readedData[i])
// 	}

// 	assert.Equal(t, 1, r.Readable())
// }
