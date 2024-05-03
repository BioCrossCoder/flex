package linkedlist

func (d Deque) Map(handler func(any) any) Deque {
	return *NewDeque(d.ToArrayList().Map(handler)...)
}

func (d Deque) Reduce(handler func(any, any) any, initial ...any) (any, error) {
	return d.ToArrayList().Reduce(handler, initial...)
}

func (d Deque) ReduceRight(handler func(any, any) any, initial ...any) (any, error) {
	return d.ToArrayList().ReduceRight(handler, initial...)
}

func (d Deque) Filter(condition func(any) bool) Deque {
	values := make([]any, 0)
	for node := d.head.Next; node != d.tail; node = node.Next {
		if condition(node.Value) {
			values = append(values, node.Value)
		}
	}
	return *NewDeque(values...)
}

func (d Deque) Some(condition func(any) bool) bool {
	node := d.head.Next
	for node != d.tail {
		if condition(node.Value) {
			return true
		}
		node = node.Next
	}
	return false
}

func (d Deque) Every(condition func(any) bool) bool {
	node := d.head.Next
	for node != d.tail {
		if !condition(node.Value) {
			return false
		}
		node = node.Next
	}
	return true
}
