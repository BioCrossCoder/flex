package collections

import "github.com/biocrosscoder/flex/typed/collections/linkedlist"

type Stack[T any] interface {
	Push(element T) (ok bool)
	Pop() (element T, ok bool)
	Peek() (element T, ok bool)
	Empty() bool
	Full() bool
}

type stack[T any] struct {
	data     *linkedlist.LinkedList[T]
	capacity int
}

func NewStack[T any](capacity int) Stack[T] {
	return &stack[T]{
		linkedlist.NewLinkedList[T](),
		capacity,
	}
}

func (s *stack[T]) Push(element T) (ok bool) {
	if !s.Full() {
		_ = s.data.Append(element)
		ok = true
	}
	return
}

func (s *stack[T]) Pop() (element T, ok bool) {
	if !s.Empty() {
		element, _ = s.data.Pop()
		ok = true
	}
	return
}

func (s stack[T]) Peek() (element T, ok bool) {
	if !s.Empty() {
		element, _ = s.data.Tail()
		ok = true
	}
	return
}

func (s stack[T]) Empty() bool {
	return s.data.Empty()
}

func (s stack[T]) Full() bool {
	return s.capacity > 0 && s.data.Len() == s.capacity
}
