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

	if r.writePoint >= r.readPoint {
		return len(r.buf) - r.writePoint + r.readPoint
	}
	return r.readPoint - r.writePoint
}

func (r *RingBuffer[T]) Write(data []T) int {

	available := r.Writable()
	if available == 0 {
		return 0
	}

	if available < len(data) {
		data = data[:available]
	}

	if len(r.buf)-r.writePoint < len(data) {
		copy(r.buf[r.writePoint:], data[:len(r.buf)-r.writePoint])
		copy(r.buf, data[len(r.buf)-r.writePoint:])
		r.writePoint = len(data) - (len(r.buf) - r.writePoint)
		available := r.Writable()
		return available
	}

	copy(r.buf[r.writePoint:], data)
	r.writePoint += len(data)

	return available
}

func (r *RingBuffer[T]) Readable() int {
	if r.writePoint >= r.readPoint {
		return r.writePoint - r.readPoint
	}

	return len(r.buf) - r.readPoint + r.writePoint
}

func (r *RingBuffer[T]) Read(n int) []T {
	if n > r.Readable() {
		n = r.Readable()
	}

	if len(r.buf)-r.readPoint < n {
		readedData := append(r.buf[r.readPoint:], r.buf[:n-len(r.buf)+r.readPoint]...)
		r.readPoint = n - len(r.buf) + r.readPoint
		return readedData
	}

	readedData := r.buf[r.readPoint : r.readPoint+n]
	r.readPoint += n

	return readedData

}
