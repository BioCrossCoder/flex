package linkedlist

import "flex/common"

func (d LinkedList[T]) Map(handler func(T) T) LinkedList[T] {
	list := d.ToArray()
	for i, item := range list {
		list[i] = handler(item)
	}
	return *NewLinkedList(list...)
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
	list := d.ToArray()
	startIndex := 0
	if initialCount == 0 {
		result = list[startIndex]
		startIndex++
	} else {
		result = initial[0]
	}
	for i := startIndex; i < d.Len(); i++ {
		result = handler(result, list[i])
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
	list := d.ToArray()
	startIndex := d.Len() - 1
	if initialCount == 0 {
		result = list[startIndex]
		startIndex--
	} else {
		result = initial[0]
	}
	for i := startIndex; i >= 0; i-- {
		result = handler(result, list[i])
	}
	return
}

func (d LinkedList[T]) Filter(condition func(T) bool) LinkedList[T] {
	values := make([]T, 0)
	for node := d.head.Next; node != d.tail; node = node.Next {
		if condition(node.Value) {
			values = append(values, node.Value)
		}
	}
	return *NewLinkedList(values...)
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
