package linkedlist

import (
	"flex/common"
)

func (d LinkedList[T]) parseCount(counts ...int) int {
	if len(counts) == 0 {
		return 1
	}
	if counts[0] <= 0 {
		return d.Len()
	}
	return counts[0]
}

func (d *LinkedList[T]) removeNode(node *listNode[T]) (value T, prev, next *listNode[T]) {
	value = node.Value
	prev = node.Prev
	next = node.Next
	prev.Next = next
	next.Prev = prev
	d.size--
	return
}

func (d *LinkedList[T]) Remove(element T, counts ...int) *LinkedList[T] {
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

func (d *LinkedList[T]) RemoveRight(element T, counts ...int) *LinkedList[T] {
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

func (d *LinkedList[T]) Clear() *LinkedList[T] {
	d.head.Next = d.tail
	d.tail.Prev = d.head
	d.size = 0
	return d
}

func (d *LinkedList[T]) insertNode(prev *listNode[T], element T) *listNode[T] {
	following := prev.Next
	node := &listNode[T]{
		Value: element,
		Prev:  prev,
		Next:  following,
	}
	prev.Next = node
	following.Prev = node
	d.size++
	return node
}

func (d *LinkedList[T]) Append(element T) *LinkedList[T] {
	_ = d.insertNode(d.tail.Prev, element)
	return d
}

func (d *LinkedList[T]) AppendLeft(element T) *LinkedList[T] {
	_ = d.insertNode(d.head, element)
	return d
}

func (d *LinkedList[T]) Pop() (element T, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	element, _, _ = d.removeNode(d.tail.Prev)
	return
}

func (d *LinkedList[T]) PopLeft() (element T, err error) {
	if d.Empty() {
		err = common.ErrEmptyList
		return
	}
	element, _, _ = d.removeNode(d.head.Next)
	return
}

func (d *LinkedList[T]) Extend(another *LinkedList[T]) *LinkedList[T] {
	for node := another.head.Next; node != another.tail; node = node.Next {
		_ = d.Append(node.Value)
	}
	return d
}

func (d *LinkedList[T]) ExtendLeft(another *LinkedList[T]) *LinkedList[T] {
	for node := another.head.Next; node != another.tail; node = node.Next {
		_ = d.AppendLeft(node.Value)
	}
	return d
}

func (d *LinkedList[T]) Insert(index int, element T) *LinkedList[T] {
	_ = d.insertNode(d.getNodeByIndex(d.parseIndex(index)-1), element)
	return d
}

func (d *LinkedList[T]) RemoveByIndex(index int) (element T, err error) {
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

func (d *LinkedList[T]) Rotate(steps ...int) *LinkedList[T] {
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

func (d *LinkedList[T]) Reverse() *LinkedList[T] {
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

func (d *LinkedList[T]) ForEach(action func(T) T) *LinkedList[T] {
	node := d.head.Next
	for node != d.tail {
		node.Value = action(node.Value)
		node = node.Next
	}
	return d
}

func (d *LinkedList[T]) Replace(oldElement, newElement T, counts ...int) *LinkedList[T] {
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

func (d *LinkedList[T]) ReplaceRight(oldElement, newElement T, counts ...int) *LinkedList[T] {
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

func (d *LinkedList[T]) Splice(start, deleteCount int, items ...T) LinkedList[T] {
	result := NewLinkedList[T]()
	if deleteCount <= 0 {
		return *result
	}
	node := d.getNodeByIndex(d.parseIndex(start))
	var value T
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

func (d *LinkedList[T]) Fill(element T, area ...int) *LinkedList[T] {
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

func (d *LinkedList[T]) Set(index int, element T) (err error) {
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

func (d *LinkedList[T]) RemoveIf(condition func(T) bool, counts ...int) LinkedList[T] {
	count := d.parseCount(counts...)
	node := d.head.Next
	result := NewLinkedList[T]()
	var value T
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

func (d *LinkedList[T]) RemoveRightIf(condition func(T) bool, counts ...int) LinkedList[T] {
	count := d.parseCount(counts...)
	node := d.tail.Prev
	result := NewLinkedList[T]()
	var value T
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

func (d *LinkedList[T]) ReplaceIf(condition func(T) bool, newElement T, counts ...int) LinkedList[T] {
	count := d.parseCount(counts...)
	node := d.head.Next
	result := NewLinkedList[T]()
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

func (d *LinkedList[T]) ReplaceRightIf(condition func(T) bool, newElement T, counts ...int) LinkedList[T] {
	count := d.parseCount(counts...)
	node := d.tail.Prev
	result := NewLinkedList[T]()
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
