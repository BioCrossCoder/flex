package linkedlist

import (
	"flex/common"
)

func (d *Deque) Remove(element any, counts ...int) *Deque {
	argCount := len(counts)
	count := 1
	if argCount >= 1 {
		count = counts[0]
	}
	if count <= 0 {
		count = d.Count(element)
	}
	node := d.head.Next
	for count > 0 && node.Next != nil {
		if node.Value != element {
			node = node.Next
			continue
		}
		previous := node.Prev
		following := node.Next
		previous.Next = following
		following.Prev = previous
		node = following
		count--
		d.size--
	}
	return d
}

func (d *Deque) RemoveRight(element any, counts ...int) *Deque {
	argCount := len(counts)
	count := 1
	if argCount >= 1 {
		count = counts[0]
	}
	if count <= 0 {
		count = d.Count(element)
	}
	node := d.tail.Prev
	for count > 0 && node.Prev != nil {
		if node.Value != element {
			node = node.Prev
			continue
		}
		previous := node.Prev
		following := node.Next
		previous.Next = following
		following.Prev = previous
		node = previous
		count--
		d.size--
	}
	return d
}

func (d *Deque) Clear() *Deque {
	d.head.Next = d.tail
	d.tail.Prev = d.head
	d.size = 0
	return d
}

func (d *Deque) Append(element any) *Deque {
	following := d.tail
	previous := following.Prev
	node := &Node{
		Value: element,
		Prev:  previous,
		Next:  following,
	}
	previous.Next = node
	following.Prev = node
	d.size++
	return d
}

func (d *Deque) AppendLeft(element any) *Deque {
	previous := d.head
	following := previous.Next
	node := &Node{
		Value: element,
		Prev:  previous,
		Next:  following,
	}
	previous.Next = node
	following.Prev = node
	d.size++
	return d
}

func (d *Deque) Pop() (element any, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	node := d.tail.Prev
	previous := node.Prev
	element = node.Value
	previous.Next = d.tail
	d.tail.Prev = previous
	d.size--
	return
}

func (d *Deque) PopLeft() (element any, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	node := d.head.Next
	following := node.Next
	element = node.Value
	following.Prev = d.head
	d.head.Next = following
	d.size--
	return
}

func (d *Deque) Extend(another *Deque) *Deque {
	for _, value := range another.ToArray() {
		_ = d.Append(value)
	}
	return d
}

func (d *Deque) ExtendLeft(another *Deque) *Deque {
	for _, value := range another.ToArray() {
		_ = d.AppendLeft(value)
	}
	return d
}

func (d *Deque) Insert(index int, element any) *Deque {
	validIndex := d.parseIndex(index)
	following := d.getNodeByIndex(validIndex)
	previous := following.Prev
	node := &Node{
		Value: element,
		Prev:  previous,
		Next:  following,
	}
	previous.Next = node
	following.Prev = node
	d.size++
	return d
}

func (d *Deque) RemoveByIndex(index int) (element any, err error) {
	if index < 0 {
		index += d.size
	}
	err = d.isIndexValid(index)
	if err != nil {
		return
	}
	node := d.getNodeByIndex(index)
	element = node.Value
	previous := node.Prev
	following := node.Next
	previous.Next = following
	following.Prev = previous
	d.size--
	return
}

func (d *Deque) Rotate(steps ...int) *Deque {
	if d.size <= 1 {
		return d
	}
	step := 1
	if len(steps) >= 1 {
		step = steps[0]
	}
	for step < 0 {
		value, _ := d.PopLeft()
		_ = d.Append(value)
		step++
	}
	for step > 0 {
		value, _ := d.Pop()
		_ = d.AppendLeft(value)
		step--
	}
	return d
}

func (d *Deque) Reverse() *Deque {
	previous := d.head
	node := previous.Next
	for node != nil {
		newNode := node.Next
		node.Next = previous
		previous = node
		node = newNode
	}
	follow := d.tail
	node = follow.Prev
	for node != nil {
		newNode := node.Prev
		node.Prev = follow
		follow = node
		node = newNode
	}
	d.head, d.tail = d.tail, d.head
	return d
}

func (d *Deque) ForEach(action func(any) any) *Deque {
	node := d.head.Next
	for node != d.tail {
		node.Value = action(node.Value)
		node = node.Next
	}
	return d
}

func (d *Deque) Replace(oldElement, newElement any, counts ...int) *Deque {
	if oldElement == newElement {
		return d
	}
	argCount := len(counts)
	count := 1
	if argCount >= 1 {
		count = counts[0]
	}
	if count <= 0 {
		count = d.Count(oldElement)
	}
	node := d.head.Next
	for count > 0 && node.Next != nil {
		if node.Value == oldElement {
			node.Value = newElement
			count--
		}
		node = node.Next
	}
	return d
}

func (d *Deque) ReplaceRight(oldElement, newElement any, counts ...int) *Deque {
	if oldElement == newElement {
		return d
	}
	argCount := len(counts)
	count := 1
	if argCount >= 1 {
		count = counts[0]
	}
	if count <= 0 {
		count = d.Count(oldElement)
	}
	node := d.tail.Prev
	for count > 0 && node.Prev != nil {
		if node.Value == oldElement {
			node.Value = newElement
			count--
		}
		node = node.Prev
	}
	return d
}

func (d *Deque) Splice(start, deleteCount int, items ...any) Deque {
	if deleteCount <= 0 {
		return *NewDeque()
	}
	start = d.parseIndex(start)
	end := d.parseIndex(start + deleteCount)
	count := end - start
	removedValues := make([]any, count)
	tail := d.getNodeByIndex(start)
	head := tail.Prev
	tail.Prev = nil
	head.Next = nil
	i := 0
	d.size -= count
	for i < count && tail.Next != nil {
		removedValues[i] = tail.Value
		i++
		tail = tail.Next
	}
	tail.Prev.Next = nil
	tail.Prev = nil
	for _, item := range items {
		head.Next = &Node{
			Value: item,
			Prev:  head,
			Next:  nil,
		}
		head = head.Next
		d.size++
	}
	head.Next = tail
	tail.Prev = head
	return *NewDeque(removedValues...)
}

func (d *Deque) Fill(element any, area ...int) *Deque {
	argCount := len(area)
	start := 0
	end := d.size
	if argCount >= 1 {
		start = d.parseIndex(area[0])
	}
	if argCount >= 2 {
		end = d.parseIndex(area[1])
	}
	count := end - start
	if d.size-end < start {
		node := d.getNodeByIndex(end - 1)
		for count > 0 {
			node.Value = element
			count--
			node = node.Prev
		}
	} else {
		node := d.getNodeByIndex(start)
		for count > 0 {
			node.Value = element
			count--
			node = node.Next
		}
	}
	return d
}

func (d *Deque) Set(index int, element any) (err error) {
	if index < 0 {
		index += d.size
	}
	err = d.isIndexValid(index)
	if err != nil {
		return
	}
	node := d.getNodeByIndex(index)
	node.Value = element
	return
}
