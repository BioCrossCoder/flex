package linkedlist

func NewLinkedList(elements ...any) *LinkedList {
	head, tail := &listNode{}, &listNode{}
	head.Next, tail.Prev = tail, head
	l := &LinkedList{head, tail, 0}
	for _, element := range elements {
		_ = l.Append(element)
	}
	return l
}

func (l LinkedList) Copy() LinkedList {
	backup := NewLinkedList()
	for node := l.head.Next; node != l.tail; node = node.Next {
		_ = backup.Append(node.Value)
	}
	return *backup
}

func (l LinkedList) Concat(another LinkedList) LinkedList {
	result := NewLinkedList()
	for _, l := range []LinkedList{l, another} {
		for node := l.head.Next; node != l.tail; node = node.Next {
			_ = result.Append(node.Value)
		}
	}
	return *result
}

func (l LinkedList) Slice(args ...int) LinkedList {
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
		return *NewLinkedList()
	}
	result := NewLinkedList()
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

func (l LinkedList) ToSpliced(start, deleteCount int, items ...any) LinkedList {
	newDeque := l.Copy()
	_ = newDeque.Splice(start, deleteCount, items...)
	return newDeque
}

func (l LinkedList) ToReversed() LinkedList {
	newDeque := l.Copy()
	_ = newDeque.Reverse()
	return newDeque
}

func (l LinkedList) With(index int, value any) LinkedList {
	newDeque := l.Copy()
	_ = newDeque.Set(index, value)
	return newDeque
}
