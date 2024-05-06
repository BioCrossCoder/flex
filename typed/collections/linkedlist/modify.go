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

func (d *LinkedList[T]) removeNode(node *listNode[T]) (prev, next *listNode[T]) {
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
		_, node = d.removeNode(node)
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
		node, _ = d.removeNode(node)
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

func (d *LinkedList[T]) Append(element T) *LinkedList[T] {
	following := d.tail
	previous := following.Prev
	node := &listNode[T]{
		Value: element,
		Prev:  previous,
		Next:  following,
	}
	previous.Next = node
	following.Prev = node
	d.size++
	return d
}

func (d *LinkedList[T]) AppendLeft(element T) *LinkedList[T] {
	previous := d.head
	following := previous.Next
	node := &listNode[T]{
		Value: element,
		Prev:  previous,
		Next:  following,
	}
	previous.Next = node
	following.Prev = node
	d.size++
	return d
}

func (d *LinkedList[T]) Pop() (element T, err error) {
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

func (d *LinkedList[T]) PopLeft() (element T, err error) {
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

func (d *LinkedList[T]) Extend(another *LinkedList[T]) *LinkedList[T] {
	for _, value := range another.ToArray() {
		_ = d.Append(value)
	}
	return d
}

func (d *LinkedList[T]) ExtendLeft(another *LinkedList[T]) *LinkedList[T] {
	for _, value := range another.ToArray() {
		_ = d.AppendLeft(value)
	}
	return d
}

func (d *LinkedList[T]) Insert(index int, element T) *LinkedList[T] {
	validIndex := d.parseIndex(index)
	following := d.getNodeByIndex(validIndex)
	previous := following.Prev
	node := &listNode[T]{
		Value: element,
		Prev:  previous,
		Next:  following,
	}
	previous.Next = node
	following.Prev = node
	d.size++
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
	node := d.getNodeByIndex(index)
	element = node.Value
	previous := node.Prev
	following := node.Next
	previous.Next = following
	following.Prev = previous
	d.size--
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
	if deleteCount <= 0 {
		return *NewLinkedList[T]()
	}
	start = d.parseIndex(start)
	end := d.parseIndex(start + deleteCount)
	count := end - start
	removedValues := make([]T, count)
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
		head.Next = &listNode[T]{
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
	node := d.getNodeByIndex(index)
	node.Value = element
	return
}

func (d *LinkedList[T]) RemoveIf(condition func(T) bool, counts ...int) LinkedList[T] {
	count := d.parseCount(counts...)
	node := d.head.Next
	removedValues := make([]T, count)
	i := 0
	for count > 0 && node.Next != nil {
		if !condition(node.Value) {
			node = node.Next
			continue
		}
		removedValues[i] = node.Value
		i++
		_, node = d.removeNode(node)
		count--
	}
	return *NewLinkedList(removedValues[:i:i]...)
}

func (d *LinkedList[T]) RemoveRightIf(condition func(T) bool, counts ...int) LinkedList[T] {
	count := d.parseCount(counts...)
	node := d.tail.Prev
	removedValues := make([]T, count)
	i := 0
	for count > 0 && node.Prev != nil {
		if !condition(node.Value) {
			node = node.Prev
			continue
		}
		removedValues[i] = node.Value
		i++
		node, _ = d.removeNode(node)
		count--
	}
	return *NewLinkedList(removedValues[:i:i]...)
}

func (d *LinkedList[T]) ReplaceIf(condition func(T) bool, newElement T, counts ...int) LinkedList[T] {
	count := d.parseCount(counts...)
	node := d.head.Next
	replacedValues := make([]T, count)
	i := 0
	for count > 0 && node.Next != nil {
		if condition(node.Value) {
			replacedValues[i] = node.Value
			i++
			node.Value = newElement
			count--
		}
		node = node.Next
	}
	return *NewLinkedList(replacedValues[:i:i]...)
}

func (d *LinkedList[T]) ReplaceRightIf(condition func(T) bool, newElement T, counts ...int) LinkedList[T] {
	count := d.parseCount(counts...)
	node := d.tail.Prev
	replacedValues := make([]T, count)
	i := 0
	for count > 0 && node.Prev != nil {
		if condition(node.Value) {
			replacedValues[i] = node.Value
			i++
			node.Value = newElement
			count--
		}
		node = node.Prev
	}
	return *NewLinkedList(replacedValues[:i:i]...)
}
