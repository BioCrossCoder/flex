package linkedlist

import "github.com/biocrosscoder/flex/common"

func (d LinkedList) IndexOf(element any) (index int) {
	index = -1
	for node := d.head.Next; node != d.tail; node = node.Next {
		index++
		if common.Equal(node.Value, element) {
			return
		}
	}
	return -1
}

func (d LinkedList) LastIndexOf(element any) (index int) {
	index = d.size
	for node := d.tail.Prev; node != d.head; node = node.Prev {
		index--
		if common.Equal(node.Value, element) {
			return
		}
	}
	return -1
}

func (d LinkedList) At(index int) (value any, err error) {
	if index < 0 {
		index += d.size
	}
	err = d.isIndexValid(index)
	if err != nil {
		return
	}
	value = d.getNodeByIndex(index).Value
	return
}

func (d LinkedList) Find(by func(any) bool) (element any, found bool) {
	for node := d.head.Next; node != d.tail; node = node.Next {
		if by(node.Value) {
			element = node.Value
			found = true
			break
		}
	}
	return
}

func (d LinkedList) FindIndex(by func(any) bool) (index int) {
	index = -1
	for node := d.head.Next; node != d.tail; node = node.Next {
		index++
		if by(node.Value) {
			return index
		}
	}
	return -1
}

func (d LinkedList) FindLast(by func(any) bool) (element any, found bool) {
	for node := d.tail.Prev; node != d.head; node = node.Prev {
		if by(node.Value) {
			element = node.Value
			found = true
			break
		}
	}
	return
}

func (d LinkedList) FindLastIndex(by func(any) bool) (index int) {
	index = d.size
	for node := d.tail.Prev; node != d.head; node = node.Prev {
		index--
		if by(node.Value) {
			return index
		}
	}
	return -1
}

func (d LinkedList) Head() (element any, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = d.head.Next.Value
	return
}

func (d LinkedList) Tail() (element any, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = d.tail.Prev.Value
	return
}

func (d LinkedList) getNodeByIndex(index int) *listNode {
	var node *listNode
	if d.nearTail(index) {
		node = d.tail
		reverseIndex := d.reverseIndex(index)
		for reverseIndex < 0 {
			reverseIndex++
			node = node.Prev
		}
	} else {
		node = d.head.Next
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
	count:=l.searchCount(counts...)
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
	count:=l.searchCount(counts...)
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
	if len(counts) == 0 || counts[0] <= 0 {
		return l.Len()
	}
	return counts[0]
}
