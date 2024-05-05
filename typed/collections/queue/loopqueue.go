package queue

import (
	"flex/common"
	"flex/typed/collections/arraylist"
)

type loopQueue[T any] struct {
	data     arraylist.ArrayList[T]
	capacity int
	head     int
	tail     int
}

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

func (q *loopQueue[T]) Enqueue(element T) (ok bool) {
	if !q.Full() {
		_ = q.data.Set(q.tail, element)
		q.tail++
		q.tail %= (q.capacity + 1)
		ok = true
	}
	return
}

func (q *loopQueue[T]) Dequeue() (element T, ok bool) {
	if !q.Empty() {
		element, _ = q.data.At(q.head)
		q.head++
		q.head %= (q.capacity + 1)
		ok = true
	}
	return
}

func (q loopQueue[T]) Peek() (element T, ok bool) {
	if !q.Empty() {
		element, _ = q.data.At(q.head)
		ok = true
	}
	return
}

func (q loopQueue[T]) Empty() bool {
	return q.head == q.tail
}

func (q loopQueue[T]) Full() bool {
	return (q.tail+1)%(q.capacity+1) == q.head
}
