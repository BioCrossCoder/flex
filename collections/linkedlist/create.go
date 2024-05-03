package linkedlist

func NewLinkedList(elements ...any) *LinkedList {
	head := &Node{}
	prev := head
	count := 0
	for _, element := range elements {
		node := &Node{
			Value: element,
			Prev:  prev,
		}
		prev.Next = node
		prev = node
		count++
	}
	tail := &Node{
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
		start = d.parseIndex(args[0])
	}
	if argsCount >= 2 {
		end = d.parseIndex(args[1])
	}
	if argsCount >= 3 {
		step = args[2]
	}
	if (start < end && step < 0) || (start > end && step > 0) || (start == end) || (step == 0) {
		return *NewLinkedList()
	}
	condition := func(start, end, step int) bool {
		if step > 0 {
			return start < end
		} else {
			return start > end
		}
	}
	values := make([]any, 0)
	var node *Node
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
	for i := start; condition(i, end, step); i += step {
		values = append(values, node.Value)
		if step < 0 {
			for j := 0; j > step; j-- {
				node = node.Prev
			}
		} else {
			for j := 0; j < step; j++ {
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
