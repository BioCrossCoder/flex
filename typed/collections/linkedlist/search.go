package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// IndexOf returns the index of the first occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (l LinkedList[T]) IndexOf(element T) (index int) {
	index = -1
	for node := l.head.Next; node != l.tail; node = node.Next {
		index++
		if common.Equal(node.Value, element) {
			return
		}
	}
	return -1
}

// LastIndexOf returns the index of the last occurrence of the specified element in this list, or -1 if this list does not contain the element.
func (l LinkedList[T]) LastIndexOf(element T) (index int) {
	index = l.size
	for node := l.tail.Prev; node != l.head; node = node.Prev {
		index--
		if common.Equal(node.Value, element) {
			return
		}
	}
	return -1
}

// At returns the element at the specified position in this list.
// If the index is out of range, the function returns an error.
func (l LinkedList[T]) At(index int) (value T, err error) {
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

// Find returns the first element in this list that satisfies the given predicate, or nil if no such element is found.
func (l LinkedList[T]) Find(by func(T) bool) (element T, found bool) {
	for node := l.head.Next; node != l.tail; node = node.Next {
		if by(node.Value) {
			element = node.Value
			found = true
			break
		}
	}
	return
}

// FindIndex returns the index of the first element in this list that satisfies the given predicate, or -1 if no such element is found.
func (l LinkedList[T]) FindIndex(by func(T) bool) (index int) {
	index = -1
	for node := l.head.Next; node != l.tail; node = node.Next {
		index++
		if by(node.Value) {
			return index
		}
	}
	return -1
}

// FindLast returns the last element in this list that satisfies the given predicate, or nil if no such element is found.
func (l LinkedList[T]) FindLast(by func(T) bool) (element T, found bool) {
	for node := l.tail.Prev; node != l.head; node = node.Prev {
		if by(node.Value) {
			element = node.Value
			found = true
			break
		}
	}
	return
}

// FindLastIndex returns the index of the last element in this list that satisfies the given predicate, or -1 if no such element is found.
func (l LinkedList[T]) FindLastIndex(by func(T) bool) (index int) {
	index = l.size
	for node := l.tail.Prev; node != l.head; node = node.Prev {
		index--
		if by(node.Value) {
			return index
		}
	}
	return -1
}

// Head returns the first element in this list, or an error if the list is empty.
func (l LinkedList[T]) Head() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l.head.Next.Value
	return
}

// Tail returns the last element in this list, or an error if the list is empty.
func (l LinkedList[T]) Tail() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l.tail.Prev.Value
	return
}

// getNodeByIndex returns the node at the specified index in this list.
func (l LinkedList[T]) getNodeByIndex(index int) *listNode[T] {
	var node *listNode[T]
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

// FindIndexes returns the indexes of elements in this list that satisfy the given predicate.
// If no count is specified, all elements that satisfy the predicate are returned.
func (l LinkedList[T]) FindIndexes(by func(T) bool, counts ...int) (indexes []int) {
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

// FindLastIndexes returns the indexes of elements in this list that satisfy the given predicate, starting from the end of the list.
// If no count is specified, all elements that satisfy the predicate are returned.
func (l LinkedList[T]) FindLastIndexes(by func(T) bool, counts ...int) (indexes []int) {
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

// Finds returns the elements in this list that satisfy the given predicate.
// If no count is specified, all elements that satisfy the predicate are returned.
func (l LinkedList[T]) Finds(by func(T) bool, counts ...int) (elements []T) {
	count := l.searchCount(counts...)
	elements = make([]T, 0)
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

// FindLasts returns the elements in this list that satisfy the given predicate, starting from the end of the list.
// If no count is specified, all elements that satisfy the predicate are returned.
func (l LinkedList[T]) FindLasts(by func(T) bool, counts ...int) (elements []T) {
	count := l.searchCount(counts...)
	elements = make([]T, 0)
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

// searchCount returns the count of elements to search for, based on the specified counts.
func (l LinkedList[T]) searchCount(counts ...int) int {
	return list.SearchCount(l.Len(), counts...)
}
