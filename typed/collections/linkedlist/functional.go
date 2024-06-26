package linkedlist

import "github.com/biocrosscoder/flex/common"

// Map applies a function to each element of the linked list and returns a new linked list containing the results.
func (l LinkedList[T]) Map(handler func(T) T) LinkedList[T] {
	result := NewLinkedList[T]()
	for node := l.head.Next; node != l.tail; node = node.Next {
		_ = result.Append(handler(node.Value))
	}
	return *result
}

// Reduce reduces the linked list to a single value by applying a function cumulatively to the elements and an initial value (if provided).
func (l LinkedList[T]) Reduce(handler func(T, T) T, initial ...T) (result T, err error) {
	if l.Len() == 0 {
		err = common.ErrEmptyList
		return
	}
	initialCount := len(initial)
	if initialCount > 1 {
		err = common.ErrTooManyArguments
		return
	}
	startNode := l.head.Next
	if initialCount == 0 {
		result = startNode.Value
		startNode = startNode.Next
	} else {
		result = initial[0]
	}
	for node := startNode; node != l.tail; node = node.Next {
		result = handler(result, node.Value)
	}
	return
}

// ReduceRight reduces the linked list to a single value by applying a function cumulatively to the elements from right to left and an initial value (if provided).
func (l LinkedList[T]) ReduceRight(handler func(T, T) T, initial ...T) (result T, err error) {
	if l.Len() == 0 {
		err = common.ErrEmptyList
		return
	}
	initialCount := len(initial)
	if initialCount > 1 {
		err = common.ErrTooManyArguments
		return
	}
	startNode := l.tail.Prev
	if initialCount == 0 {
		result = startNode.Value
		startNode = startNode.Prev
	} else {
		result = initial[0]
	}
	for node := startNode; node != l.head; node = node.Prev {
		result = handler(result, node.Value)
	}
	return
}

// Filter returns a new linked list containing elements for which the specified condition function returns true.
func (l LinkedList[T]) Filter(condition func(T) bool) LinkedList[T] {
	result := NewLinkedList[T]()
	for node := l.head.Next; node != l.tail; node = node.Next {
		if condition(node.Value) {
			_ = result.Append(node.Value)
		}
	}
	return *result
}

// Some checks if there's at least one element in the linked list satisfying the specified condition function.
func (l LinkedList[T]) Some(condition func(T) bool) bool {
	node := l.head.Next
	for node != l.tail {
		if condition(node.Value) {
			return true
		}
		node = node.Next
	}
	return false
}

// Every checks if every element in the linked list satisfies the specified condition function.
func (l LinkedList[T]) Every(condition func(T) bool) bool {
	node := l.head.Next
	for node != l.tail {
		if !condition(node.Value) {
			return false
		}
		node = node.Next
	}
	return true
}
