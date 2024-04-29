package buffer

type SliceBuffer[T any] struct {
	buf []T
}

func NewSliceBuffer[T any]() *SliceBuffer[T] {
	return &SliceBuffer[T]{}
}

func (b *SliceBuffer[T]) Write(data []T) {
	b.buf = append(b.buf, data...)
}

func (b *SliceBuffer[T]) Readable() int {
	return len(b.buf)
}

func (b *SliceBuffer[T]) Read(n int) []T {
	if n > len(b.buf) {
		n = len(b.buf)
	}

	readedData := b.buf[:n]
	// 읽은 만큼 전진
	b.buf = b.buf[n:]

	return readedData
}
