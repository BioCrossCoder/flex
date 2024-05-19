package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

type listNode struct {
	Value any
	Prev  *listNode
	Next  *listNode
}

type LinkedList struct {
	head *listNode
	tail *listNode
	size int
}

func (l LinkedList) sliceIndex(index int, accessible bool) int {
	return list.SliceIndex(index, l.Len(), accessible)
}

func (l LinkedList) parseIndex(index int) int {
	return list.ParseIndex(index, l.Len())
}

func (l LinkedList) isIndexValid(index int) error {
	return list.IsIndexValid(index, l.Len())
}

func (l LinkedList) nearTail(index int) bool {
	return l.size-index <= index
}

func (l LinkedList) reverseIndex(index int) int {
	return index - l.size
}

func (l LinkedList) Len() int {
	return l.size
}

func (l LinkedList) Count(element any) (count int) {
	for node := l.head; node != nil; node = node.Next {
		if common.Equal(node.Value, element) {
			count++
		}
	}
	return
}

func (l LinkedList) Includes(element any) bool {
	return l.IndexOf(element) != -1
}

func (l LinkedList) Empty() bool {
	return l.size == 0
}

func (l LinkedList) ToArray() []any {
	arr := make([]any, l.size)
	i := 0
	for node := l.head.Next; node != l.tail; node = node.Next {
		arr[i] = node.Value
		i++
	}
	return arr
}

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
