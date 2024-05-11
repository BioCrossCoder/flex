package queue

import (
	"github.com/biocrosscoder/flex/collections/arraylist"
	"github.com/biocrosscoder/flex/common"
)

type loopQueue struct {
	data     arraylist.ArrayList
	capacity int
	head     int
	tail     int
}

func NewLoopQueue(capacity int) (q Queue, err error) {
	if capacity <= 0 {
		err = common.ErrInvalidCapacity
		return
	}
	q = &loopQueue{
		arraylist.Repeat(nil, capacity+1),
		capacity,
		0,
		0,
	}
	return
}

func (q *loopQueue) Enqueue(element any) (ok bool) {
	if !q.Full() {
		_ = q.data.Set(q.tail, element)
		q.tail++
		q.tail %= (q.capacity + 1)
		ok = true
	}
	return
}

func (q *loopQueue) Dequeue() (element any, ok bool) {
	if !q.Empty() {
		element, _ = q.data.At(q.head)
		_ = q.data.Set(q.head, nil)
		q.head++
		q.head %= (q.capacity + 1)
		ok = true
	}
	return
}

func (q loopQueue) Peek() (element any, ok bool) {
	if !q.Empty() {
		element, _ = q.data.At(q.head)
		ok = true
	}
	return
}

func (q loopQueue) Empty() bool {
	return q.head == q.tail
}

func (q loopQueue) Full() bool {
	return (q.tail+1)%(q.capacity+1) == q.head
}
