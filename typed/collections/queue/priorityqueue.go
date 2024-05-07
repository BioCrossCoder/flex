package queue

import (
	"flex/common"
	"flex/typed/collections/dict"
	"flex/typed/collections/sortedcontainers"
	"flex/typed/collections/sortedcontainers/sortedlist"
)

type PriorityQueue[T any] interface {
	Enqueue(element T, priority int) (ok bool)
	Dequeue() (element T, ok bool)
	Peek() (element T, ok bool)
	Empty() bool
	Full() bool
	Size() int
}

type priorityQueue[T any] struct {
	elements sortedcontainers.SortedDict[int, Queue[T]]
	capacity int
	size     int
}

func NewPriorityQueue[T any](capacity int) (q PriorityQueue[T], err error) {
	if capacity <= 0 {
		err = common.ErrInvalidCapacity
		return
	}
	elements := sortedcontainers.NewSortedDict(sortedlist.DescendOrder, dict.Dict[int, Queue[T]]{})
	q = &priorityQueue[T]{*elements, capacity, 0}
	return
}

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

func (pq priorityQueue[T]) Peek() (element T, ok bool) {
	if !pq.Empty() {
		key, _ := pq.elements.KeyAt(0)
		group := pq.elements.Get(key)
		element, ok = group.Peek()
	}
	return
}

func (pq priorityQueue[T]) Empty() bool {
	return pq.Size() == 0
}

func (pq priorityQueue[T]) Full() bool {
	return pq.Size() == pq.capacity
}

func (pq priorityQueue[T]) Size() int {
	return pq.size
}
