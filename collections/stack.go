package collections

import "github.com/biocrosscoder/flex/collections/linkedlist"

// Stack interface represents a stack data structure
type Stack interface {
	// Push adds an element to the top of the stack if it's not full
	Push(element any) (ok bool)
	// Pop removes and returns the top element from the stack if it's not empty
	Pop() (element any, ok bool)
	// Peek returns the top element of the stack without removing it if the stack is not empty
	Peek() (element any, ok bool)
	// Empty checks if the stack is empty
	Empty() bool
	// Full checks if the stack is full based on its capacity
	Full() bool
}

// stack struct represents the implementation of the Stack interface using a linked list
type stack struct {
	data     *linkedlist.LinkedList
	capacity int
}

// NewStack creates a new stack with the specified capacity
func NewStack(capacity int) Stack {
	return &stack{
		linkedlist.NewLinkedList(),
		capacity,
	}
}

// Push adds an element to the top of the stack if it's not full
func (s *stack) Push(element any) (ok bool) {
	if !s.Full() {
		_ = s.data.Append(element)
		ok = true
	}
	return
}

// Pop removes and returns the top element from the stack if it's not empty
func (s *stack) Pop() (element any, ok bool) {
	if !s.Empty() {
		element, _ = s.data.Pop()
		ok = true
	}
	return
}

// Peek returns the top element of the stack without removing it if the stack is not empty
func (s stack) Peek() (element any, ok bool) {
	if !s.Empty() {
		element, _ = s.data.Tail()
		ok = true
	}
	return
}

// Empty checks if the stack is empty
func (s stack) Empty() bool {
	return s.data.Empty()
}

// Full checks if the stack is full based on its capacity
func (s stack) Full() bool {
	return s.capacity > 0 && s.data.Len() == s.capacity
}
