package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

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

func (l LinkedList[T]) Head() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l.head.Next.Value
	return
}

func (l LinkedList[T]) Tail() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l.tail.Prev.Value
	return
}

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

func (l LinkedList[T]) searchCount(counts ...int) int {
	return list.SearchCount(l.Len(), counts...)
}
