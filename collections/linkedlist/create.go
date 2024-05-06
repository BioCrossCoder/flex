package linkedlist

func NewLinkedList(elements ...any) *LinkedList {
	head := &listNode{}
	prev := head
	count := 0
	for _, element := range elements {
		node := &listNode{
			Value: element,
			Prev:  prev,
		}
		prev.Next = node
		prev = node
		count++
	}
	tail := &listNode{
		Prev: prev,
	}
	prev.Next = tail
	return &LinkedList{
		head: head,
		tail: tail,
		size: count,
	}
}

func (d LinkedList) Copy() LinkedList {
	return *NewLinkedList(d.ToArray()...)
}

func (d LinkedList) Concat(another LinkedList) LinkedList {
	return *NewLinkedList(append(d.ToArray(), another.ToArray()...)...)
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
	i := 0
	var values []any
	var node *listNode
	if d.size-1-start < start {
		node = d.tail
		reverseIndex := start - d.size
		for reverseIndex < 0 {
			reverseIndex++
			node = node.Prev
		}
	} else {
		node = d.head.Next
		for i := 0; i < start; i++ {
			node = node.Next
		}
	}
	if step < 0 {
		values = make([]any, (start-end-step-1)/(-step))
		for j := start; j > end; j += step {
			values[i] = node.Value
			i++
			for k := 0; k > step; k-- {
				node = node.Prev
			}
		}
	} else {
		values = make([]any, (end-start+step-1)/step)
		for j := start; j < end; j += step {
			values[i] = node.Value
			i++
			for k := 0; k < step; k++ {
				node = node.Next
			}
		}
	}
	return *NewLinkedList(values...)
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
