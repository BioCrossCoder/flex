package linkedlist

import "flex/common"

func (d LinkedList[T]) Map(handler func(T) T) LinkedList[T] {
	result := NewLinkedList[T]()
	for node := d.head.Next; node != d.tail; node = node.Next {
		_ = result.Append(handler(node.Value))
	}
	return *result
}

func (d LinkedList[T]) Reduce(handler func(T, T) T, initial ...T) (result T, err error) {
	if d.Len() == 0 {
		err = common.ErrEmptyList
		return
	}
	initialCount := len(initial)
	if initialCount > 1 {
		err = common.ErrTooManyArguments
		return
	}
	startNode := d.head.Next
	if initialCount == 0 {
		result = startNode.Value
		startNode = startNode.Next
	} else {
		result = initial[0]
	}
	for node := startNode; node != d.tail; node = node.Next {
		result = handler(result, node.Value)
	}
	return
}

func (d LinkedList[T]) ReduceRight(handler func(T, T) T, initial ...T) (result T, err error) {
	if d.Len() == 0 {
		err = common.ErrEmptyList
		return
	}
	initialCount := len(initial)
	if initialCount > 1 {
		err = common.ErrTooManyArguments
		return
	}
	startNode := d.tail.Prev
	if initialCount == 0 {
		result = startNode.Value
		startNode = startNode.Prev
	} else {
		result = initial[0]
	}
	for node := startNode; node != d.head; node = node.Prev {
		result = handler(result, node.Value)
	}
	return
}

func (d LinkedList[T]) Filter(condition func(T) bool) LinkedList[T] {
	result := NewLinkedList[T]()
	for node := d.head.Next; node != d.tail; node = node.Next {
		if condition(node.Value) {
			_ = result.Append(node.Value)
		}
	}
	return *result
}

func (d LinkedList[T]) Some(condition func(T) bool) bool {
	node := d.head.Next
	for node != d.tail {
		if condition(node.Value) {
			return true
		}
		node = node.Next
	}
	return false
}

func (d LinkedList[T]) Every(condition func(T) bool) bool {
	node := d.head.Next
	for node != d.tail {
		if !condition(node.Value) {
			return false
		}
		node = node.Next
	}
	return true
}
