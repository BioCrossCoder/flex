package queue

import "flex/collections/linkedlist"

type Queue interface {
	Enqueue(element any) (ok bool)
	Dequeue() (element any, ok bool)
	Peek() (element any, ok bool)
	Empty() bool
	Full() bool
}

type linearQueue struct {
	data     *linkedlist.LinkedList
	capacity int
}

func NewQueue(capacity int) Queue {
	return &linearQueue{
		linkedlist.NewLinkedList(),
		capacity,
	}
}

func (q *linearQueue) Enqueue(element any) (ok bool) {
	if !q.Full() {
		_ = q.data.Append(element)
		ok = true
	}
	return
}

func (q *linearQueue) Dequeue() (element any, ok bool) {
	if !q.Empty() {
		element, _ = q.data.PopLeft()
		ok = true
	}
	return
}

func (q linearQueue) Peek() (element any, ok bool) {
	if !q.Empty() {
		element, _ = q.data.Head()
		ok = true
	}
	return
}

func (q linearQueue) Empty() bool {
	return q.data.Empty()
}

func (q linearQueue) Full() bool {
	return q.capacity > 0 && q.data.Len() == q.capacity
}
