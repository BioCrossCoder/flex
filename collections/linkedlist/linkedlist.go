// Package linkedlist provides a doubly linked list implementation.
package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// listNode represents a node in the linked list with a value and references to the previous and next nodes.
type listNode struct {
	Value any
	Prev  *listNode
	Next  *listNode
}

// LinkedList represents a doubly linked list with a head, tail, and size.
type LinkedList struct {
	head *listNode
	tail *listNode
	size int
}

// sliceIndex returns the effective index considering negative indices and out-of-bounds values.
func (l LinkedList) sliceIndex(index int, accessible bool) int {
	return list.SliceIndex(index, l.Len(), accessible)
}

// parseIndex returns the effective index considering negative indices.
func (l LinkedList) parseIndex(index int) int {
	return list.ParseIndex(index, l.Len())
}

// isIndexValid checks if the index is valid within the linked list size.
func (l LinkedList) isIndexValid(index int) error {
	return list.IsIndexValid(index, l.Len())
}

// nearTail checks if the index is closer to the tail of the linked list.
func (l LinkedList) nearTail(index int) bool {
	return l.size-index <= index
}

// reverseIndex returns the reversed index based on the linked list size.
func (l LinkedList) reverseIndex(index int) int {
	return index - l.size
}

// Len returns the size of the linked list.
func (l LinkedList) Len() int {
	return l.size
}

// Count returns the number of occurrences of a specific element in the linked list.
func (l LinkedList) Count(element any) (count int) {
	for node := l.head; node != nil; node = node.Next {
		if common.Equal(node.Value, element) {
			count++
		}
	}
	return
}

// Includes checks if a specified element is included in the linked list.
func (l LinkedList) Includes(element any) bool {
	return l.IndexOf(element) != -1
}

// Empty checks if the linked list is empty.
func (l LinkedList) Empty() bool {
	return l.size == 0
}

// ToArray converts the linked list to an array of elements.
func (l LinkedList) ToArray() []any {
	arr := make([]any, l.size)
	i := 0
	for node := l.head.Next; node != l.tail; node = node.Next {
		arr[i] = node.Value
		i++
	}
	return arr
}

// Equal checks if two linked lists are equal by comparing their elements.
func (l LinkedList) Equal(another LinkedList) bool {
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
