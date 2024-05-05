package linkedlist

import (
	"flex/common"
)

func (d *LinkedList) Remove(element any, counts ...int) *LinkedList {
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
		if !common.Equal(node.Value, element) {
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

func (d *LinkedList) RemoveRight(element any, counts ...int) *LinkedList {
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
		if !common.Equal(node.Value, element) {
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

func (d *LinkedList) Clear() *LinkedList {
	d.head.Next = d.tail
	d.tail.Prev = d.head
	d.size = 0
	return d
}

func (d *LinkedList) Append(element any) *LinkedList {
	following := d.tail
	previous := following.Prev
	node := &listNode{
		Value: element,
		Prev:  previous,
		Next:  following,
	}
	previous.Next = node
	following.Prev = node
	d.size++
	return d
}

func (d *LinkedList) AppendLeft(element any) *LinkedList {
	previous := d.head
	following := previous.Next
	node := &listNode{
		Value: element,
		Prev:  previous,
		Next:  following,
	}
	previous.Next = node
	following.Prev = node
	d.size++
	return d
}

func (d *LinkedList) Pop() (element any, err error) {
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

func (d *LinkedList) PopLeft() (element any, err error) {
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

func (d *LinkedList) Extend(another *LinkedList) *LinkedList {
	for _, value := range another.ToArray() {
		_ = d.Append(value)
	}
	return d
}

func (d *LinkedList) ExtendLeft(another *LinkedList) *LinkedList {
	for _, value := range another.ToArray() {
		_ = d.AppendLeft(value)
	}
	return d
}

func (d *LinkedList) Insert(index int, element any) *LinkedList {
	validIndex := d.parseIndex(index)
	following := d.getNodeByIndex(validIndex)
	previous := following.Prev
	node := &listNode{
		Value: element,
		Prev:  previous,
		Next:  following,
	}
	previous.Next = node
	following.Prev = node
	d.size++
	return d
}

func (d *LinkedList) RemoveByIndex(index int) (element any, err error) {
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

func (d *LinkedList) Rotate(steps ...int) *LinkedList {
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

func (d *LinkedList) Reverse() *LinkedList {
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

func (d *LinkedList) ForEach(action func(any) any) *LinkedList {
	node := d.head.Next
	for node != d.tail {
		node.Value = action(node.Value)
		node = node.Next
	}
	return d
}

func (d *LinkedList) Replace(oldElement, newElement any, counts ...int) *LinkedList {
	if common.Equal(oldElement, newElement) {
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
		if common.Equal(node.Value, oldElement) {
			node.Value = newElement
			count--
		}
		node = node.Next
	}
	return d
}

func (d *LinkedList) ReplaceRight(oldElement, newElement any, counts ...int) *LinkedList {
	if common.Equal(oldElement, newElement) {
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
		if common.Equal(node.Value, oldElement) {
			node.Value = newElement
			count--
		}
		node = node.Prev
	}
	return d
}

func (d *LinkedList) Splice(start, deleteCount int, items ...any) LinkedList {
	if deleteCount <= 0 {
		return *NewLinkedList()
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
		head.Next = &listNode{
			Value: item,
			Prev:  head,
			Next:  nil,
		}
		head = head.Next
		d.size++
	}
	head.Next = tail
	tail.Prev = head
	return *NewLinkedList(removedValues...)
}

func (d *LinkedList) Fill(element any, area ...int) *LinkedList {
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

func (d *LinkedList) Set(index int, element any) (err error) {
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
