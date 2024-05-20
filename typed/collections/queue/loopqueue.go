package queue

import (
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/typed/collections/arraylist"
)

// loopQueue is a circular queue implemented using an arraylist and supports generic types.
type loopQueue[T any] struct {
	data     arraylist.ArrayList[T]
	capacity int
	head     int
	tail     int
}

// NewLoopQueue creates a new loopQueue with the specified capacity.
func NewLoopQueue[T any](capacity int) (q Queue[T], err error) {
	if capacity <= 0 {
		err = common.ErrInvalidCapacity
		return
	}
	q = &loopQueue[T]{
		make(arraylist.ArrayList[T], capacity+1),
		capacity,
		0,
		0,
	}
	return
}

// Enqueue adds an element to the rear of the queue.
func (q *loopQueue[T]) Enqueue(element T) (ok bool) {
	if !q.Full() {
		_ = q.data.Set(q.tail, element)
		q.tail++
		q.tail %= (q.capacity + 1)
		ok = true
	}
	return
}

// Dequeue removes and returns the front element of the queue.
func (q *loopQueue[T]) Dequeue() (element T, ok bool) {
	if !q.Empty() {
		element, _ = q.data.At(q.head)
		_ = q.data.Set(q.head, *new(T))
		q.head++
		q.head %= (q.capacity + 1)
		ok = true
	}
	return
}

// Peek returns the front element of the queue without removing it.
func (q loopQueue[T]) Peek() (element T, ok bool) {
	if !q.Empty() {
		element, _ = q.data.At(q.head)
		ok = true
	}
	return
}

// Empty returns true if the queue is empty, otherwise returns false.
func (q loopQueue[T]) Empty() bool {
	return q.head == q.tail
}

// Full returns true if the queue is full, otherwise returns false.
func (q loopQueue[T]) Full() bool {
	return (q.tail+1)%(q.capacity+1) == q.head
}
