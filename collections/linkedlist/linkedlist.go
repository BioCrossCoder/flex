package linkedlist

import (
	"flex/common"
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

func (d LinkedList) parseIndex(index int) int {
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

func (d LinkedList) isIndexValid(index int) (err error) {
	if index < 0 || index >= d.Len() {
		err = common.ErrOutOfRange
	}
	return
}

func (d LinkedList) nearTail(index int) bool {
	return d.size-index <= index
}

func (d LinkedList) reverseIndex(index int) int {
	return index - d.size
}

func (d LinkedList) Len() int {
	return d.size
}

func (d LinkedList) Count(element any) (count int) {
	for node := d.head; node != nil; node = node.Next {
		if node.Value == element {
			count++
		}
	}
	return
}

func (d LinkedList) Includes(element any) bool {
	return d.IndexOf(element) != -1
}

func (d LinkedList) Empty() bool {
	return d.size == 0
}

func (d LinkedList) ToArray() []any {
	l := make([]any, d.size)
	i := 0
	for node := d.head.Next; node != d.tail; node = node.Next {
		l[i] = node.Value
		i++
	}
	return l
}

func (d LinkedList) Equal(another LinkedList) bool {
	length1 := d.Len()
	length2 := another.Len()
	if length1 != length2 {
		return false
	}
	node1 := d.head.Next
	node2 := another.head.Next
	for node1 != d.tail {
		if node1.Value != node2.Value {
			return false
		}
		node1 = node1.Next
		node2 = node2.Next
	}
	return true
}
