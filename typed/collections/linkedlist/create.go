package linkedlist

// NewLinkedList creates a new LinkedList with the given elements.
func NewLinkedList[T any](elements ...T) *LinkedList[T] {
	head, tail := &listNode[T]{}, &listNode[T]{}
	head.Next, tail.Prev = tail, head
	l := &LinkedList[T]{head, tail, 0}
	for _, element := range elements {
		_ = l.Append(element)
	}
	return l
}

// Copy creates a copy of the LinkedList.
func (l LinkedList[T]) Copy() LinkedList[T] {
	backup := NewLinkedList[T]()
	for node := l.head.Next; node != l.tail; node = node.Next {
		_ = backup.Append(node.Value)
	}
	return *backup
}

// Concat concatenates the current LinkedList with another LinkedList.
func (l LinkedList[T]) Concat(another LinkedList[T]) LinkedList[T] {
	result := NewLinkedList[T]()
	for _, l := range []LinkedList[T]{l, another} {
		for node := l.head.Next; node != l.tail; node = node.Next {
			_ = result.Append(node.Value)
		}
	}
	return *result
}

// Slice returns a new LinkedList that is a subset of the original LinkedList based on the provided slice arguments.
func (l LinkedList[T]) Slice(args ...int) LinkedList[T] {
	argsCount := len(args)
	if argsCount == 0 {
		return l.Copy()
	}
	start := 0
	end := l.size
	step := 1
	if argsCount >= 1 {
		start = l.sliceIndex(args[0], true)
	}
	if argsCount >= 2 {
		end = l.sliceIndex(args[1], false)
	}
	if argsCount >= 3 {
		step = args[2]
	}
	if (start < end && step < 0) || (start > end && step > 0) || (start == end) || (step == 0) {
		return *NewLinkedList[T]()
	}
	result := NewLinkedList[T]()
	node := l.getNodeByIndex(start)
	if step < 0 {
		for i := start; i > end && node != nil; i += step {
			_ = result.Append(node.Value)
			for j := 0; j > step && node != nil; j-- {
				node = node.Prev
			}
		}
	} else {
		for i := start; i < end && node != nil; i += step {
			_ = result.Append(node.Value)
			for j := 0; j < step && node != nil; j++ {
				node = node.Next
			}
		}
	}
	return *result
}

// ToSpliced returns a new LinkedList with items spliced into the original LinkedList at the specified index.
func (l LinkedList[T]) ToSpliced(start, deleteCount int, items ...T) LinkedList[T] {
	newDeque := l.Copy()
	_ = newDeque.Splice(start, deleteCount, items...)
	return newDeque
}

// ToReversed returns a new LinkedList with the elements in reverse order compared to the original LinkedList.
func (l LinkedList[T]) ToReversed() LinkedList[T] {
	newDeque := l.Copy()
	_ = newDeque.Reverse()
	return newDeque
}

// With returns a new LinkedList with the value at the specified index replaced with the given value.
func (l LinkedList[T]) With(index int, value T) LinkedList[T] {
	newDeque := l.Copy()
	_ = newDeque.Set(index, value)
	return newDeque
}
