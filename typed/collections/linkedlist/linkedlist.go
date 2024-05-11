package linkedlist

import "github.com/biocrosscoder/flex/common"

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
	length := l.Len()
	if index < 0 {
		index += length
	}
	if index < 0 {
		index = -1
		if accessible {
			index++
		}
	}
	if index >= length {
		index = length
		if accessible {
			index--
		}
	}
	return index
}

func (d LinkedList[T]) parseIndex(index int) int {
	length := d.Len()
	if index < 0 {
		index += length
		if index < 0 {
			return 0
		}
	} else if index > length {
		return length
	}
	return index
}

func (d LinkedList[T]) isIndexValid(index int) (err error) {
	if index < 0 || index >= d.Len() {
		err = common.ErrOutOfRange
	}
	return
}

func (d LinkedList[T]) nearTail(index int) bool {
	return d.size-index <= index
}

func (d LinkedList[T]) reverseIndex(index int) int {
	return index - d.size
}

func (d LinkedList[T]) Len() int {
	return d.size
}

func (d LinkedList[T]) Count(element T) (count int) {
	for node := d.head; node != nil; node = node.Next {
		if common.Equal(node.Value, element) {
			count++
		}
	}
	return
}

func (d LinkedList[T]) Includes(element T) bool {
	return d.IndexOf(element) != -1
}

func (d LinkedList[T]) Empty() bool {
	return d.size == 0
}

func (d LinkedList[T]) ToArray() []T {
	l := make([]T, d.size)
	i := 0
	for node := d.head.Next; node != d.tail; node = node.Next {
		l[i] = node.Value
		i++
	}
	return l
}

func (d LinkedList[T]) Equal(another LinkedList[T]) bool {
	length1 := d.Len()
	length2 := another.Len()
	if length1 != length2 {
		return false
	}
	node1 := d.head.Next
	node2 := another.head.Next
	for node1 != d.tail {
		if !common.Equal(node1.Value, node2.Value) {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}
