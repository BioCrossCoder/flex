package linkedlist

import "flex/common"

func (d LinkedList) Map(handler func(any) any) LinkedList {
	result := NewLinkedList()
	for node := d.head.Next; node != d.tail; node = node.Next {
		_ = result.Append(handler(node.Value))
	}
	return *result
}

func (d LinkedList) Reduce(handler func(any, any) any, initial ...any) (result any, err error) {
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

func (d LinkedList) ReduceRight(handler func(any, any) any, initial ...any) (result any, err error) {
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

func (d LinkedList) Filter(condition func(any) bool) LinkedList {
	result := NewLinkedList()
	for node := d.head.Next; node != d.tail; node = node.Next {
		if condition(node.Value) {
			_ = result.Append(node.Value)
		}
	}
	return *result
}

func (d LinkedList) Some(condition func(any) bool) bool {
	node := d.head.Next
	for node != d.tail {
		if condition(node.Value) {
			return true
		}
		node = node.Next
	}
	return false
}

func (d LinkedList) Every(condition func(any) bool) bool {
	node := d.head.Next
	for node != d.tail {
		if !condition(node.Value) {
			return false
		}
		node = node.Next
	}
	return true
}
