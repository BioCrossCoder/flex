// Package queue provides several implementations of a queue data structure.
package queue

import "github.com/biocrosscoder/flex/collections/linkedlist"

// Queue is an interface for a queue data structure.
type Queue interface {
	Enqueue(element any) (ok bool)
	Dequeue() (element any, ok bool)
	Peek() (element any, ok bool)
	Empty() bool
	Full() bool
}

// linearQueue is a simple implementation of a queue using a linked list.
type linearQueue struct {
	data     *linkedlist.LinkedList
	capacity int
}

// NewQueue creates a new linear queue with the specified capacity.
func NewQueue(capacity int) Queue {
	return &linearQueue{
		linkedlist.NewLinkedList(),
		capacity,
	}
}

// Enqueue adds an element to the end of the queue.
func (q *linearQueue) Enqueue(element any) (ok bool) {
	if !q.Full() {
		_ = q.data.Append(element)
		ok = true
	}
	return
}

// Dequeue removes and returns the element at the front of the queue.
func (q *linearQueue) Dequeue() (element any, ok bool) {
	if !q.Empty() {
		element, _ = q.data.PopLeft()
		ok = true
	}
	return
}

// Peek returns the element at the front of the queue without removing it.
func (q linearQueue) Peek() (element any, ok bool) {
	if !q.Empty() {
		element, _ = q.data.Head()
		ok = true
	}
	return
}

// Empty checks if the queue is empty.
func (q linearQueue) Empty() bool {
	return q.data.Empty()
}

// Full checks if the queue is full.
func (q linearQueue) Full() bool {
	return q.capacity > 0 && q.data.Len() == q.capacity
}
