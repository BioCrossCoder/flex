package linkedlist

import (
	"flex/common"
)

func (d LinkedList) parseCount(counts ...int) int {
	if len(counts) == 0 {
		return 1
	}
	if counts[0] <= 0 {
		return d.Len()
	}
	return counts[0]
}

func (d *LinkedList) removeNode(node *listNode) (value any, prev, next *listNode) {
	value = node.Value
	prev = node.Prev
	next = node.Next
	prev.Next = next
	next.Prev = prev
	d.size--
	return
}

func (d *LinkedList) Remove(element any, counts ...int) *LinkedList {
	count := d.parseCount(counts...)
	node := d.head.Next
	for count > 0 && node.Next != nil {
		if !common.Equal(node.Value, element) {
			node = node.Next
			continue
		}
		_, _, node = d.removeNode(node)
		count--
	}
	return d
}

func (d *LinkedList) RemoveRight(element any, counts ...int) *LinkedList {
	count := d.parseCount(counts...)
	node := d.tail.Prev
	for count > 0 && node.Prev != nil {
		if !common.Equal(node.Value, element) {
			node = node.Prev
			continue
		}
		_, node, _ = d.removeNode(node)
		count--
	}
	return d
}

func (d *LinkedList) Clear() *LinkedList {
	d.head.Next = d.tail
	d.tail.Prev = d.head
	d.size = 0
	return d
}

func (d *LinkedList) insertNode(prev *listNode, element any) *listNode {
	following := prev.Next
	node := &listNode{
		Value: element,
		Prev:  prev,
		Next:  following,
	}
	prev.Next = node
	following.Prev = node
	d.size++
	return node
}

func (d *LinkedList) Append(element any) *LinkedList {
	_ = d.insertNode(d.tail.Prev, element)
	return d
}

func (d *LinkedList) AppendLeft(element any) *LinkedList {
	_ = d.insertNode(d.head, element)
	return d
}

func (d *LinkedList) Pop() (element any, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	element, _, _ = d.removeNode(d.tail.Prev)
	return
}

func (d *LinkedList) PopLeft() (element any, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	element, _, _ = d.removeNode(d.head.Next)
	return
}

func (d *LinkedList) Extend(another *LinkedList) *LinkedList {
	for node := another.head.Next; node != another.tail; node = node.Next {
		_ = d.Append(node.Value)
	}
	return d
}

func (d *LinkedList) ExtendLeft(another *LinkedList) *LinkedList {
	for node := another.head.Next; node != another.tail; node = node.Next {
		_ = d.AppendLeft(node.Value)
	}
	return d
}

func (d *LinkedList) Insert(index int, element any) *LinkedList {
	_ = d.insertNode(d.getNodeByIndex(d.parseIndex(index)-1), element)
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
	element, _, _ = d.removeNode(d.getNodeByIndex(index))
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
	count := d.parseCount(counts...)
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
	count := d.parseCount(counts...)
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
	result := NewLinkedList()
	if deleteCount <= 0 {
		return *result
	}
	node := d.getNodeByIndex(d.parseIndex(start))
	var value any
	for node != d.tail {
		if deleteCount == 0 {
			break
		}
		value, _, node = d.removeNode(node)
		_ = result.Append(value)
		deleteCount--
	}
	node = node.Prev
	for _, item := range items {
		node = d.insertNode(node, item)
	}
	return *result
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
	d.getNodeByIndex(index).Value = element
	return
}

func (d *LinkedList) RemoveIf(condition func(any) bool, counts ...int) LinkedList {
	count := d.parseCount(counts...)
	node := d.head.Next
	result := NewLinkedList()
	var value any
	for count > 0 && node.Next != nil {
		if !condition(node.Value) {
			node = node.Next
			continue
		}
		value, _, node = d.removeNode(node)
		_ = result.Append(value)
		count--
	}
	return *result
}

func (d *LinkedList) RemoveRightIf(condition func(any) bool, counts ...int) LinkedList {
	count := d.parseCount(counts...)
	node := d.tail.Prev
	result := NewLinkedList()
	var value any
	for count > 0 && node.Prev != nil {
		if !condition(node.Value) {
			node = node.Prev
			continue
		}
		value, node, _ = d.removeNode(node)
		_ = result.Append(value)
		count--
	}
	return *result
}

func (d *LinkedList) ReplaceIf(condition func(any) bool, newElement any, counts ...int) LinkedList {
	count := d.parseCount(counts...)
	node := d.head.Next
	result := NewLinkedList()
	for count > 0 && node.Next != nil {
		if condition(node.Value) {
			_ = result.Append(node.Value)
			node.Value = newElement
			count--
		}
		node = node.Next
	}
	return *result
}

func (d *LinkedList) ReplaceRightIf(condition func(any) bool, newElement any, counts ...int) LinkedList {
	count := d.parseCount(counts...)
	node := d.tail.Prev
	result := NewLinkedList()
	for count > 0 && node.Prev != nil {
		if condition(node.Value) {
			result.Append(node.Value)
			node.Value = newElement
			count--
		}
		node = node.Prev
	}
	return *result
}
