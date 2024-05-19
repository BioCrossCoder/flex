package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

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

func (l LinkedList) Head() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l.head.Next.Value
	return
}

func (l LinkedList) Tail() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l.tail.Prev.Value
	return
}

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

func (l LinkedList) searchCount(counts ...int) int {
	return list.SearchCount(l.Len(), counts...)
}
