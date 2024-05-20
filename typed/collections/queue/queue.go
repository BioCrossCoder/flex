// Package queue provides several implementations of a queue data structure.
package queue

import (
	"github.com/biocrosscoder/flex/typed/collections/linkedlist"
)

// Queue is the interface for a generic queue data structure
type Queue[T any] interface {
	// Enqueue adds an element to the end of the queue
	Enqueue(element T) (ok bool)
	// Dequeue removes and returns the element at the front of the queue
	Dequeue() (element T, ok bool)
	// Peek returns the element at the front of the queue without removing it
	Peek() (element T, ok bool)
	// Empty checks if the queue is empty
	Empty() bool
	// Full checks if the queue is full
	Full() bool
}

// linearQueue represents a linear queue data structure using a linked list
type linearQueue[T any] struct {
	data     *linkedlist.LinkedList[T]
	capacity int
}

// NewQueue creates a new linear queue with the specified capacity
func NewQueue[T any](capacity int) Queue[T] {
	return &linearQueue[T]{
		linkedlist.NewLinkedList[T](),
		capacity,
	}
}

// Enqueue adds an element to the rear of the queue
func (q *linearQueue[T]) Enqueue(element T) (ok bool) {
	if !q.Full() {
		_ = q.data.Append(element)
		ok = true
	}
	return
}

// Dequeue removes and returns the element at the front of the queue
func (q *linearQueue[T]) Dequeue() (element T, ok bool) {
	if !q.Empty() {
		element, _ = q.data.PopLeft()
		ok = true
	}
	return
}

// Peek returns the element at the front of the queue without removing it
func (q linearQueue[T]) Peek() (element T, ok bool) {
	if !q.Empty() {
		element, _ = q.data.Head()
		ok = true
	}
	return
}

// Empty checks if the queue is empty
func (q linearQueue[T]) Empty() bool {
	return q.data.Empty()
}

// Full checks if the queue is full based on its capacity
func (q linearQueue[T]) Full() bool {
	return q.capacity > 0 && q.data.Len() == q.capacity
}
