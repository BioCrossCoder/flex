package linkedlist

import "flex/common"

func (d Deque) Map(handler func(any) any) Deque {
	list := d.ToArray()
	for i, item := range list {
		list[i] = handler(item)
	}
	return *NewDeque(list...)
}

func (d Deque) Reduce(handler func(any, any) any, initial ...any) (result any, err error) {
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

func (d Deque) ReduceRight(handler func(any, any) any, initial ...any) (result any, err error) {
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

func (d Deque) Filter(condition func(any) bool) Deque {
	values := make([]any, 0)
	for node := d.head.Next; node != d.tail; node = node.Next {
		if condition(node.Value) {
			values = append(values, node.Value)
		}
	}
	return *NewDeque(values...)
}

func (d Deque) Some(condition func(any) bool) bool {
	node := d.head.Next
	for node != d.tail {
		if condition(node.Value) {
			return true
		}
		node = node.Next
	}
	return false
}

func (d Deque) Every(condition func(any) bool) bool {
	node := d.head.Next
	for node != d.tail {
		if !condition(node.Value) {
			return false
		}
		node = node.Next
	}
	return true
}
