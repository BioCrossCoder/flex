// Package linkedlist provides a doubly linked list implementation.
package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// listNode is a node in the linked list, containing a value and pointers to the previous and next nodes.
type listNode[T any] struct {
	Value T
	Prev  *listNode[T]
	Next  *listNode[T]
}

// LinkedList represents a doubly linked list with references to the head, tail, and the size of the list.
type LinkedList[T any] struct {
	head *listNode[T]
	tail *listNode[T]
	size int
}

// sliceIndex returns the actual index for the given index value and whether out-of-bounds index is accessible
func (l LinkedList[T]) sliceIndex(index int, accessible bool) int {
	return list.SliceIndex(index, l.Len(), accessible)
}

// parseIndex returns the actual index for the given index value, handling negative indices and out-of-bounds indices
func (l LinkedList[T]) parseIndex(index int) int {
	return list.ParseIndex(index, l.Len())
}

// isIndexValid checks if the given index is valid within the linked list and returns an error if it's not
func (l LinkedList[T]) isIndexValid(index int) error {
	return list.IsIndexValid(index, l.Len())
}

// nearTail calculates whether the given index is closer to the tail of the linked list
func (l LinkedList[T]) nearTail(index int) bool {
	return l.size-index <= index
}

// reverseIndex returns the reverse index for the given index within the linked list
func (l LinkedList[T]) reverseIndex(index int) int {
	return index - l.size
}

// Len returns the size of the linked list
func (l LinkedList[T]) Len() int {
	return l.size
}

// Count returns the number of occurrences of the given element in the linked list
func (l LinkedList[T]) Count(element T) (count int) {
	for node := l.head; node != nil; node = node.Next {
		if common.Equal(node.Value, element) {
			count++
		}
	}
	return
}

// Includes checks if the linked list includes the given element
func (l LinkedList[T]) Includes(element T) bool {
	return l.IndexOf(element) != -1
}

// Empty checks if the linked list is empty
func (l LinkedList[T]) Empty() bool {
	return l.size == 0
}

// ToArray converts the linked list to an array
func (l LinkedList[T]) ToArray() []T {
	arr := make([]T, l.size)
	i := 0
	for node := l.head.Next; node != l.tail; node = node.Next {
		arr[i] = node.Value
		i++
	}
	return arr
}

// Equal checks if this linked list is equal to another linked list
func (l LinkedList[T]) Equal(another LinkedList[T]) bool {
	length1 := l.Len()
	length2 := another.Len()
	if length1 != length2 {
		return false
	}
	node1 := l.head.Next
	node2 := another.head.Next
	for node1 != l.tail {
		if !common.Equal(node1.Value, node2.Value) {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}
