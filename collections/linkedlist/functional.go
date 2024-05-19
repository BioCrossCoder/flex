package linkedlist

import "github.com/biocrosscoder/flex/common"

// Map applies the handler function to each element in the linked list and returns a new linked list containing the results.
func (l LinkedList) Map(handler func(any) any) LinkedList {
	result := NewLinkedList()
	for node := l.head.Next; node != l.tail; node = node.Next {
		_ = result.Append(handler(node.Value))
	}
	return *result
}

// Reduce applies the handler function to each element in the linked list and returns a single result.
// The initial value, if provided, is used as the initial result; otherwise, the value of the first element is used.
func (l LinkedList) Reduce(handler func(any, any) any, initial ...any) (result any, err error) {
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

// ReduceRight applies the handler function to each element in the linked list in reverse order and returns a single result.
// The initial value, if provided, is used as the initial result; otherwise, the value of the last element is used.
func (l LinkedList) ReduceRight(handler func(any, any) any, initial ...any) (result any, err error) {
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

// Filter returns a new linked list containing only the elements for which the condition function returns true.
func (l LinkedList) Filter(condition func(any) bool) LinkedList {
	result := NewLinkedList()
	for node := l.head.Next; node != l.tail; node = node.Next {
		if condition(node.Value) {
			_ = result.Append(node.Value)
		}
	}
	return *result
}

// Some checks if at least one element in the linked list satisfies the condition function.
func (l LinkedList) Some(condition func(any) bool) bool {
	node := l.head.Next
	for node != l.tail {
		if condition(node.Value) {
			return true
		}
		node = node.Next
	}
	return false
}

// Every checks if every element in the linked list satisfies the condition function.
func (l LinkedList) Every(condition func(any) bool) bool {
	node := l.head.Next
	for node != l.tail {
		if !condition(node.Value) {
			return false
		}
		node = node.Next
	}
	return true
}
