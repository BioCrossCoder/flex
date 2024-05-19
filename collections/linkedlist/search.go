package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// IndexOf returns the index of the first occurrence of the specified element in the linked list.
func (l LinkedList) IndexOf(element any) (index int) {
	index = -1
	for node := l.head.Next; node != l.tail; node = node.Next {
		index++
		if common.Equal(node.Value, element) {
			return
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of the specified element in the linked list.
func (l LinkedList) LastIndexOf(element any) (index int) {
	index = l.size
	for node := l.tail.Prev; node != l.head; node = node.Prev {
		index--
		if common.Equal(node.Value, element) {
			return
		}
	}
	return -1
}

// At returns the element at the specified index in the linked list.
func (l LinkedList) At(index int) (value any, err error) {
	if index < 0 {
		index += l.size
	}
	err = l.isIndexValid(index)
	if err != nil {
		return
	}
	value = l.getNodeByIndex(index).Value
	return
}

// Find returns the first element in the linked list for which the given function returns true.
func (l LinkedList) Find(by func(any) bool) (element any, found bool) {
	for node := l.head.Next; node != l.tail; node = node.Next {
		if by(node.Value) {
			element = node.Value
			found = true
			break
		}
	}
	return
}

// FindIndex returns the index of the first element in the linked list for which the given function returns true.
func (l LinkedList) FindIndex(by func(any) bool) (index int) {
	index = -1
	for node := l.head.Next; node != l.tail; node = node.Next {
		index++
		if by(node.Value) {
			return index
		}
	}
	return -1
}

// FindLast returns the last element in the linked list for which the given function returns true.
func (l LinkedList) FindLast(by func(any) bool) (element any, found bool) {
	for node := l.tail.Prev; node != l.head; node = node.Prev {
		if by(node.Value) {
			element = node.Value
			found = true
			break
		}
	}
	return
}

// FindLastIndex returns the index of the last element in the linked list for which the given function returns true.
func (l LinkedList) FindLastIndex(by func(any) bool) (index int) {
	index = l.size
	for node := l.tail.Prev; node != l.head; node = node.Prev {
		index--
		if by(node.Value) {
			return index
		}
	}
	return -1
}

// Head returns the first element of the linked list.
func (l LinkedList) Head() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l.head.Next.Value
	return
}

// Tail returns the last element of the linked list.
func (l LinkedList) Tail() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l.tail.Prev.Value
	return
}

// getNodeByIndex returns the node at the specified index in the linked list.
func (l LinkedList) getNodeByIndex(index int) *listNode {
	var node *listNode
	if l.nearTail(index) {
		node = l.tail
		reverseIndex := l.reverseIndex(index)
		for reverseIndex < 0 {
			reverseIndex++
			node = node.Prev
		}
	} else {
		node = l.head.Next
		for index > 0 {
			index--
			node = node.Next
		}
	}
	return node
}

// FindIndexes returns the indexes of the elements in the linked list for which the given function returns true.
func (l LinkedList) FindIndexes(by func(any) bool, counts ...int) (indexes []int) {
	count := l.searchCount(counts...)
	indexes = make([]int, 0)
	index := 0
	for node := l.head.Next; node != l.tail; node = node.Next {
		if count == 0 {
			break
		}
		if by(node.Value) {
			indexes = append(indexes, index)
			count--
		}
		index++
	}
	return
}

// FindLastIndexes returns the indexes of the last elements in the linked list for which the given function returns true.
func (l LinkedList) FindLastIndexes(by func(any) bool, counts ...int) (indexes []int) {
	count := l.searchCount(counts...)
	indexes = make([]int, 0)
	index := l.size - 1
	for node := l.tail.Prev; node != l.head; node = node.Prev {
		if count == 0 {
			break
		}
		if by(node.Value) {
			indexes = append(indexes, index)
			count--
		}
		index--
	}
	return
}

// Finds returns the elements in the linked list for which the given function returns true.
func (l LinkedList) Finds(by func(any) bool, counts ...int) (elements []any) {
	count := l.searchCount(counts...)
	elements = make([]any, 0)
	for node := l.head.Next; node != l.tail; node = node.Next {
		if count == 0 {
			break
		}
		if by(node.Value) {
			elements = append(elements, node.Value)
			count--
		}
	}
	return
}

// FindLasts returns the last elements in the linked list for which the given function returns true.
func (l LinkedList) FindLasts(by func(any) bool, counts ...int) (elements []any) {
	count := l.searchCount(counts...)
	elements = make([]any, 0)
	for node := l.tail.Prev; node != l.head; node = node.Prev {
		if count == 0 {
			break
		}
		if by(node.Value) {
			elements = append(elements, node.Value)
			count--
		}
	}
	return
}

// searchCount returns the count of elements to search for based on the specified search counts.
func (l LinkedList) searchCount(counts ...int) int {
	return list.SearchCount(l.Len(), counts...)
}
