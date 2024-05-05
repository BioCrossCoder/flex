package linkedlist

import "flex/common"

func (d LinkedList[T]) IndexOf(element T) (index int) {
	index = -1
	for node := d.head.Next; node != d.tail; node = node.Next {
		index++
		if common.Equal(node.Value, element) {
			return
		}
	}
	return -1
}

func (d LinkedList[T]) LastIndexOf(element T) (index int) {
	index = d.size
	for node := d.tail.Prev; node != d.head; node = node.Prev {
		index--
		if common.Equal(node.Value, element) {
			return
		}
	}
	return -1
}

func (d LinkedList[T]) At(index int) (value T, err error) {
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

func (d LinkedList[T]) Find(by func(T) bool) (element T, found bool) {
	for node := d.head.Next; node != d.tail; node = node.Next {
		if by(node.Value) {
			element = node.Value
			found = true
			break
		}
	}
	return
}

func (d LinkedList[T]) FindIndex(by func(T) bool) (index int) {
	index = -1
	for node := d.head.Next; node != d.tail; node = node.Next {
		index++
		if by(node.Value) {
			return index
		}
	}
	return -1
}

func (d LinkedList[T]) FindLast(by func(T) bool) (element T, found bool) {
	for node := d.tail.Prev; node != d.head; node = node.Prev {
		if by(node.Value) {
			element = node.Value
			found = true
			break
		}
	}
	return
}

func (d LinkedList[T]) FindLastIndex(by func(T) bool) (index int) {
	index = d.size
	for node := d.tail.Prev; node != d.head; node = node.Prev {
		index--
		if by(node.Value) {
			return index
		}
	}
	return -1
}

func (d LinkedList[T]) Head() (element T, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = d.head.Next.Value
	return
}

func (d LinkedList[T]) Tail() (element T, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = d.tail.Prev.Value
	return
}

func (d LinkedList[T]) getNodeByIndex(index int) *listNode[T] {
	var node *listNode[T]
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
