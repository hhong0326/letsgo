package buffer

type RingBuffer[T any] struct {
	buf        []T
	readPoint  int
	writePoint int
}

func NewRingBuffer[T any](size int) *RingBuffer[T] {
	return &RingBuffer[T]{
		buf: make([]T, size),
	}
}

func (r *RingBuffer[T]) Writable() int {
	return len(r.buf) - r.writePoint
}

func (r *RingBuffer[T]) Write(data []T) int {

	available := r.Writable()
	if available == 0 {
		return 0
	}

	if available < len(data) {
		data = data[:available]
	}

	copy(r.buf[r.writePoint:], data)
	r.writePoint += len(data)

	return available
}

func (r *RingBuffer[T]) Readable() int {
	return r.writePoint - r.readPoint
}

func (r *RingBuffer[T]) Read(n int) []T {
	if n > r.Readable() {
		n = r.Readable()
	}

	readedData := r.buf[r.readPoint : r.readPoint+n]
	r.readPoint += n

	return readedData

}
