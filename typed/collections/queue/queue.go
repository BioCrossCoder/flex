package queue

import (
	"flex/typed/collections/linkedlist"
)

type Queue[T any] interface {
	Enqueue(element T) (ok bool)
	Dequeue() (element T, ok bool)
	Peek() (element T, ok bool)
	Empty() bool
	Full() bool
}

type linearQueue[T any] struct {
	data     *linkedlist.LinkedList[T]
	capacity int
}

func NewQueue[T any](capacity int) Queue[T] {
	return &linearQueue[T]{
		linkedlist.NewLinkedList[T](),
		capacity,
	}
}

func (q *linearQueue[T]) Enqueue(element T) (ok bool) {
	if !q.Full() {
		_ = q.data.Append(element)
		ok = true
	}
	return
}

func (q *linearQueue[T]) Dequeue() (element T, ok bool) {
	if !q.Empty() {
		element, _ = q.data.PopLeft()
		ok = true
	}
	return
}

func (q linearQueue[T]) Peek() (element T, ok bool) {
	if !q.Empty() {
		element, _ = q.data.Head()
		ok = true
	}
	return
}

func (q linearQueue[T]) Empty() bool {
	return q.data.Empty()
}

func (q linearQueue[T]) Full() bool {
	return q.capacity > 0 && q.data.Len() == q.capacity
}
