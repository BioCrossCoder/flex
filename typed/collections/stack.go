package collections

import "github.com/biocrosscoder/flex/typed/collections/linkedlist"

// Stack represents a generic stack data structure with methods for push, pop, peek, check if empty or full.
type Stack[T any] interface {
	// Push adds an element to the stack if it is not full, and returns true if successful.
	Push(element T) (ok bool)
	// Pop removes and returns the top element from the stack if it is not empty, and returns true if successful.
	Pop() (element T, ok bool)
	// Peek returns the top element from the stack without removing it if the stack is not empty, and returns true if successful.
	Peek() (element T, ok bool)
	// Empty checks if the stack is empty and returns true if it contains no elements.
	Empty() bool
	// Full checks if the stack is full based on its capacity and returns true if it has reached its capacity.
	Full() bool
}

// stack represents the internal implementation of the generic stack using a linked list and a capacity constraint.
type stack[T any] struct {
	data     *linkedlist.LinkedList[T]
	capacity int
}

// NewStack creates a new instance of a generic stack with the given capacity.
func NewStack[T any](capacity int) Stack[T] {
	return &stack[T]{
		linkedlist.NewLinkedList[T](),
		capacity,
	}
}

// Push adds an element to the stack if it is not full, and returns true if successful.
func (s *stack[T]) Push(element T) (ok bool) {
	if !s.Full() {
		_ = s.data.Append(element)
		ok = true
	}
	return
}

// Pop removes and returns the top element from the stack if it is not empty, and returns true if successful.
func (s *stack[T]) Pop() (element T, ok bool) {
	if !s.Empty() {
		element, _ = s.data.Pop()
		ok = true
	}
	return
}

// Peek returns the top element from the stack without removing it if the stack is not empty, and returns true if successful.
func (s stack[T]) Peek() (element T, ok bool) {
	if !s.Empty() {
		element, _ = s.data.Tail()
		ok = true
	}
	return
}

// Empty checks if the stack is empty and returns true if it contains no elements.
func (s stack[T]) Empty() bool {
	return s.data.Empty()
}

// Full checks if the stack is full based on its capacity and returns true if it has reached its capacity.
func (s stack[T]) Full() bool {
	return s.capacity > 0 && s.data.Len() == s.capacity
}
