package queue

import (
	"container/heap"
	"github.com/biocrosscoder/flex/common"
)

// PriorityQueue is an interface for a priority queue.
type PriorityQueue interface {
	Enqueue(element any, priority int) (ok bool)
	Dequeue() (element any, ok bool)
	Peek() (element any, ok bool)
	Empty() bool
	Full() bool
	Size() int
}

// priorityQueue is a priority queue implementation using a heap.
type priorityQueue struct {
	data *priorityHeap
}

// NewPriorityQueue creates a new priority queue with the given capacity.
func NewPriorityQueue(capacity int) (q PriorityQueue, err error) {
	if capacity <= 0 {
		err = common.ErrInvalidCapacity
		return
	}
	ph := make(priorityHeap, 0, capacity)
	q = &priorityQueue{&ph}
	return
}

// Enqueue adds an element to the priority queue with the given priority.
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

// Dequeue removes and returns the element with the highest priority from the priority queue.
func (pq *priorityQueue) Dequeue() (element any, ok bool) {
	if !pq.Empty() {
		element = heap.Pop(pq.data).(*priorityQueueElement).value
		ok = true
	}
	return
}

// Peek returns the element with the highest priority from the priority queue without removing it.
func (pq *priorityQueue) Peek() (element any, ok bool) {
	if !pq.Empty() {
		element = (*pq.data)[0].value
		ok = true
	}
	return
}

// Empty checks if the priority queue is empty.
func (pq priorityQueue) Empty() bool {
	return pq.Size() == 0
}

// Full checks if the priority queue is full.
func (pq priorityQueue) Full() bool {
	return pq.Size() == cap(*pq.data)
}

// Size returns the number of elements in the priority queue.
func (pq priorityQueue) Size() int {
	return len(*pq.data)
}

// priorityQueueElement is an element in the priority queue.
type priorityQueueElement struct {
	value    any
	priority int
}

// priorityHeap is a heap of priorityQueueElements.
type priorityHeap []*priorityQueueElement

// Len returns the length of the heap.
func (h priorityHeap) Len() int {
	return len(h)
}

// Less returns true if the element at index i has a higher priority than the element at index j.
func (h priorityHeap) Less(i, j int) bool {
	return h[i].priority > h[j].priority
}

// Swap swaps the elements at index i and j in the heap.
func (h priorityHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push adds an element to the heap.
func (h *priorityHeap) Push(x any) {
	*h = append(*h, x.(*priorityQueueElement))
}

// Pop removes and returns the element with the highest priority from the heap.
func (h *priorityHeap) Pop() any {
	oldHeap := *h
	n := oldHeap.Len()
	x := oldHeap[n-1]
	*h = oldHeap[0 : n-1]
	return x
}
