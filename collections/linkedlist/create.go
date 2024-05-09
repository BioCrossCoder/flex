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

func (d LinkedList) Copy() LinkedList {
	backup := NewLinkedList()
	for node := d.head.Next; node != d.tail; node = node.Next {
		_ = backup.Append(node.Value)
	}
	return *backup
}

func (d LinkedList) Concat(another LinkedList) LinkedList {
	result := NewLinkedList()
	for _, l := range []LinkedList{d, another} {
		for node := l.head.Next; node != l.tail; node = node.Next {
			_ = result.Append(node.Value)
		}
	}
	return *result
}

func (d LinkedList) Slice(args ...int) LinkedList {
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
		return *NewLinkedList()
	}
	result := NewLinkedList()
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

func (d LinkedList) ToSpliced(start, deleteCount int, items ...any) LinkedList {
	newDeque := d.Copy()
	_ = newDeque.Splice(start, deleteCount, items...)
	return newDeque
}

func (d LinkedList) ToReversed() LinkedList {
	newDeque := d.Copy()
	_ = newDeque.Reverse()
	return newDeque
}

func (d LinkedList) With(index int, value any) LinkedList {
	newDeque := d.Copy()
	_ = newDeque.Set(index, value)
	return newDeque
}
