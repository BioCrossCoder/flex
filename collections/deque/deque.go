package deque

import (
	"flex/collections/list"
	"flex/common"
)

type Node struct {
	Value any
	Prev  *Node
	Next  *Node
}

type Deque struct {
	head *Node
	tail *Node
	size int
}

func (d Deque) parseIndex(index int) int {
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

func (d Deque) isIndexValid(index int) (err error) {
	if index < 0 || index >= d.Len() {
		err = common.ErrOutOfRange
	}
	return
}

func (d Deque) nearTail(index int) bool {
	return d.size-index <= index
}

func (d Deque) reverseIndex(index int) int {
	return index - d.size
}

func (d Deque) Len() int {
	return d.size
}

func (d Deque) Count(element any) (count int) {
	for node := d.head; node != nil; node = node.Next {
		if node.Value == element {
			count++
		}
	}
	return
}

func (d Deque) Includes(element any) bool {
	return d.IndexOf(element) != -1
}

func (d Deque) Empty() bool {
	return d.size == 0
}

func (d Deque) ToList() list.List {
	l := make(list.List, d.size)
	i := 0
	for node := d.head.Next; node != d.tail; node = node.Next {
		l[i] = node.Value
		i++
	}
	return l
}

func (d Deque) Equal(another Deque) bool {
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
