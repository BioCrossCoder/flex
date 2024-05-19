package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

type listNode[T any] struct {
	Value T
	Prev  *listNode[T]
	Next  *listNode[T]
}

type LinkedList[T any] struct {
	head *listNode[T]
	tail *listNode[T]
	size int
}

func (l LinkedList[T]) sliceIndex(index int, accessible bool) int {
	return list.SliceIndex(index, l.Len(), accessible)
}

func (l LinkedList[T]) parseIndex(index int) int {
	return list.ParseIndex(index, l.Len())
}

func (l LinkedList[T]) isIndexValid(index int) error {
	return list.IsIndexValid(index, l.Len())
}

func (l LinkedList[T]) nearTail(index int) bool {
	return l.size-index <= index
}

func (l LinkedList[T]) reverseIndex(index int) int {
	return index - l.size
}

func (l LinkedList[T]) Len() int {
	return l.size
}

func (l LinkedList[T]) Count(element T) (count int) {
	for node := l.head; node != nil; node = node.Next {
		if common.Equal(node.Value, element) {
			count++
		}
	}
	return
}

func (l LinkedList[T]) Includes(element T) bool {
	return l.IndexOf(element) != -1
}

func (l LinkedList[T]) Empty() bool {
	return l.size == 0
}

func (l LinkedList[T]) ToArray() []T {
	arr := make([]T, l.size)
	i := 0
	for node := l.head.Next; node != l.tail; node = node.Next {
		arr[i] = node.Value
		i++
	}
	return arr
}

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
