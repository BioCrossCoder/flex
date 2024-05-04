package queue

import (
	"container/heap"
	"flex/common"
)

type PriorityQueue interface {
	Enqueue(element any, priority int) (ok bool)
	Dequeue() (element any, ok bool)
	Peek() (element any, ok bool)
	Empty() bool
	Full() bool
	Size() int
}

type priorityQueue struct {
	data *priorityHeap
}

func NewPriorityQueue(capacity int) (q PriorityQueue, err error) {
	if capacity <= 0 {
		err = common.ErrInvalidCapacity
		return
	}
	ph := make(priorityHeap, 0, capacity)
	q = &priorityQueue{&ph}
	return
}

func (pq *priorityQueue) Enqueue(element any, priority int) (ok bool) {
	if !pq.Full() {
		heap.Push(pq.data,
			&priorityQueueElement{
				value:    element,
				priority: priority,
			},
		)
		ok = true
	}
	return
}

func (pq *priorityQueue) Dequeue() (element any, ok bool) {
	if !pq.Empty() {
		element = heap.Pop(pq.data).(*priorityQueueElement).value
		ok = true
	}
	return
}

func (pq *priorityQueue) Peek() (element any, ok bool) {
	if !pq.Empty() {
		element = (*pq.data)[0].value
		ok = true
	}
	return
}

func (pq priorityQueue) Empty() bool {
	return pq.Size() == 0
}

func (pq priorityQueue) Full() bool {
	return pq.Size() == cap(*pq.data)
}

func (pq priorityQueue) Size() int {
	return len(*pq.data)
}

type priorityQueueElement struct {
	value    any
	priority int
}

type priorityHeap []*priorityQueueElement

func (h priorityHeap) Len() int {
	return len(h)
}

func (h priorityHeap) Less(i, j int) bool {
	return h[i].priority > h[j].priority
}

func (h priorityHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *priorityHeap) Push(x any) {
	*h = append(*h, x.(*priorityQueueElement))
}

func (h *priorityHeap) Pop() any {
	oldHeap := *h
	n := oldHeap.Len()
	x := oldHeap[n-1]
	*h = oldHeap[0 : n-1]
	return x
}
