package queue

import (
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/typed/collections/dict"
	"github.com/biocrosscoder/flex/typed/collections/sortedcontainers"
	"github.com/biocrosscoder/flex/typed/collections/sortedcontainers/sortedlist"
)

// PriorityQueue defines the methods of a priority queue data structure.
type PriorityQueue[T any] interface {
	// Enqueue adds an element with a certain priority to the queue.
	Enqueue(element T, priority int) (ok bool)
	// Dequeue removes and returns the element with the highest priority from the queue.
	Dequeue() (element T, ok bool)
	// Peek returns the element with the highest priority from the queue without removing it.
	Peek() (element T, ok bool)
	// Empty checks if the queue is empty.
	Empty() bool
	// Full checks if the queue is full.
	Full() bool
	// Size returns the current size of the queue.
	Size() int
}

// priorityQueue represents the internal state of a priority queue data structure.
type priorityQueue[T any] struct {
	elements sortedcontainers.SortedDict[int, Queue[T]]
	capacity int
	size     int
}

// NewPriorityQueue creates a new priority queue with the specified capacity.
func NewPriorityQueue[T any](capacity int) (q PriorityQueue[T], err error) {
	if capacity <= 0 {
		err = common.ErrInvalidCapacity
		return
	}
	elements := sortedcontainers.NewSortedDict(sortedlist.DescendOrder, dict.Dict[int, Queue[T]]{})
	q = &priorityQueue[T]{*elements, capacity, 0}
	return
}

// Enqueue adds an element with a certain priority to the queue.
func (pq *priorityQueue[T]) Enqueue(element T, priority int) (ok bool) {
	if !pq.Full() {
		if !pq.elements.Has(priority) {
			pq.elements.Set(priority, NewQueue[T](-1))
		}
		group := pq.elements.Get(priority)
		_ = group.Enqueue(element)
		pq.size++
		ok = true
	}
	return
}

// Dequeue removes and returns the element with the highest priority from the queue.
func (pq *priorityQueue[T]) Dequeue() (element T, ok bool) {
	if !pq.Empty() {
		key, _ := pq.elements.KeyAt(0)
		group := pq.elements.Get(key)
		element, ok = group.Dequeue()
		if group.Empty() {
			_ = pq.elements.Delete(key)
		}
		pq.size--
	}
	return
}

// Peek returns the element with the highest priority from the queue without removing it.
func (pq priorityQueue[T]) Peek() (element T, ok bool) {
	if !pq.Empty() {
		key, _ := pq.elements.KeyAt(0)
		group := pq.elements.Get(key)
		element, ok = group.Peek()
	}
	return
}

// Empty checks if the queue is empty.
func (pq priorityQueue[T]) Empty() bool {
	return pq.Size() == 0
}

// Full checks if the queue is full.
func (pq priorityQueue[T]) Full() bool {
	return pq.Size() == pq.capacity
}

// Size returns the current size of the queue.
func (pq priorityQueue[T]) Size() int {
	return pq.size
}
