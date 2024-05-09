package linkedlist

func NewLinkedList[T any](elements ...T) *LinkedList[T] {
	head, tail := &listNode[T]{}, &listNode[T]{}
	head.Next, tail.Prev = tail, head
	l := &LinkedList[T]{head, tail, 0}
	for _, element := range elements {
		_ = l.Append(element)
	}
	return l
}

func (d LinkedList[T]) Copy() LinkedList[T] {
	backup := NewLinkedList[T]()
	for node := d.head.Next; node != d.tail; node = node.Next {
		_ = backup.Append(node.Value)
	}
	return *backup
}

func (d LinkedList[T]) Concat(another LinkedList[T]) LinkedList[T] {
	result := NewLinkedList[T]()
	for _, l := range []LinkedList[T]{d, another} {
		for node := l.head.Next; node != l.tail; node = node.Next {
			_ = result.Append(node.Value)
		}
	}
	return *result
}

func (d LinkedList[T]) Slice(args ...int) LinkedList[T] {
	argsCount := len(args)
	if argsCount == 0 {
		return d.Copy()
	}
	start := 0
	end := d.size
	step := 1
	if argsCount >= 1 {
		start = d.sliceIndex(args[0], true)
	}
	if argsCount >= 2 {
		end = d.sliceIndex(args[1], false)
	}
	if argsCount >= 3 {
		step = args[2]
	}
	if (start < end && step < 0) || (start > end && step > 0) || (start == end) || (step == 0) {
		return *NewLinkedList[T]()
	}
	result := NewLinkedList[T]()
	node := d.getNodeByIndex(start)
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

func (d LinkedList[T]) ToSpliced(start, deleteCount int, items ...T) LinkedList[T] {
	newDeque := d.Copy()
	_ = newDeque.Splice(start, deleteCount, items...)
	return newDeque
}

func (d LinkedList[T]) ToReversed() LinkedList[T] {
	newDeque := d.Copy()
	_ = newDeque.Reverse()
	return newDeque
}

func (d LinkedList[T]) With(index int, value T) LinkedList[T] {
	newDeque := d.Copy()
	_ = newDeque.Set(index, value)
	return newDeque
}
