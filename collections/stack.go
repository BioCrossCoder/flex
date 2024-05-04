package collections

import "flex/collections/linkedlist"

type Stack interface {
	Push(element any) bool
	Pop() (element any, ok bool)
	Peek() (element any, ok bool)
	Empty() bool
	Full() bool
}

type stack struct {
	data     *linkedlist.LinkedList
	capacity int
}

func NewStack(capacity int) Stack {
	return &stack{
		linkedlist.NewLinkedList(),
		capacity,
	}
}

func (s *stack) Push(element any) bool {
	if s.Full() {
		return false
	}
	_ = s.data.Append(element)
	return true
}

func (s *stack) Pop() (element any, ok bool) {
	element, err := s.data.Pop()
	if err == nil {
		ok = true
	}
	return
}

func (s stack) Peek() (element any, ok bool) {
	element, err := s.data.Tail()
	if err == nil {
		ok = true
	}
	return
}

func (s stack) Empty() bool {
	return s.data.Empty()
}

func (s stack) Full() bool {
	return s.capacity > 0 && s.data.Len() == s.capacity
}
