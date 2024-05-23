package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// parseCount parses the counts and returns the integer value.
func (l LinkedList) parseCount(counts ...int) int {
	return list.ParseCount(l.Len(), counts...)
}

// removeNode removes the specified node from the linked list and returns the value, previous node, and next node.
func (l *LinkedList) removeNode(node *listNode) (value any, prev, next *listNode) {
	value = node.Value
	prev = node.Prev
	next = node.Next
	prev.Next = next
	next.Prev = prev
	l.size--
	return
}

// Remove removes the specified element from the linked list and returns the modified linked list.
func (l *LinkedList) Remove(element any, counts ...int) *LinkedList {
	count := l.parseCount(counts...)
	node := l.head.Next
	for count > 0 && node.Next != nil {
		if !common.Equal(node.Value, element) {
			node = node.Next
			continue
		}
		_, _, node = l.removeNode(node)
		count--
	}
	return l
}

// RemoveRight removes the specified element from the right end of the linked list and returns the modified linked list.
func (l *LinkedList) RemoveRight(element any, counts ...int) *LinkedList {
	count := l.parseCount(counts...)
	node := l.tail.Prev
	for count > 0 && node.Prev != nil {
		if !common.Equal(node.Value, element) {
			node = node.Prev
			continue
		}
		_, node, _ = l.removeNode(node)
		count--
	}
	return l
}

// Clear clears the linked list and returns the modified linked list.
func (l *LinkedList) Clear() *LinkedList {
	l.head.Next = l.tail
	l.tail.Prev = l.head
	l.size = 0
	return l
}

// insertNode inserts a new node with the specified element after the specified previous node and returns the newly inserted node.
func (l *LinkedList) insertNode(prev *listNode, element any) *listNode {
	following := prev.Next
	node := &listNode{
		Value: element,
		Prev:  prev,
		Next:  following,
	}
	prev.Next = node
	following.Prev = node
	l.size++
	return node
}

// Append appends the specified element to the end of the linked list and returns the modified linked list.
func (l *LinkedList) Append(element any) *LinkedList {
	_ = l.insertNode(l.tail.Prev, element)
	return l
}

// AppendLeft appends the specified element to the beginning of the linked list and returns the modified linked list.
func (l *LinkedList) AppendLeft(element any) *LinkedList {
	_ = l.insertNode(l.head, element)
	return l
}

// Pop removes and returns the last element from the linked list.
func (l *LinkedList) Pop() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element, _, _ = l.removeNode(l.tail.Prev)
	return
}

// PopLeft removes and returns the first element from the linked list.
func (l *LinkedList) PopLeft() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element, _, _ = l.removeNode(l.head.Next)
	return
}

// Extend appends all elements from another linked list to the end of the current linked list and returns the modified linked list.
func (l *LinkedList) Extend(another *LinkedList) *LinkedList {
	for node := another.head.Next; node != another.tail; node = node.Next {
		_ = l.Append(node.Value)
	}
	return l
}

// ExtendLeft appends all elements from another linked list to the beginning of the current linked list and returns the modified linked list.
func (l *LinkedList) ExtendLeft(another *LinkedList) *LinkedList {
	for node := another.head.Next; node != another.tail; node = node.Next {
		_ = l.AppendLeft(node.Value)
	}
	return l
}

// Insert inserts the specified element at the specified index in the linked list and returns the modified linked list.
func (l *LinkedList) Insert(index int, element any) *LinkedList {
	_ = l.insertNode(l.getNodeByIndex(l.parseIndex(index)-1), element)
	return l
}

// RemoveByIndex removes and returns the element at the specified index in the linked list.
func (l *LinkedList) RemoveByIndex(index int) (element any, err error) {
	if index < 0 {
		index += l.size
	}
	err = l.isIndexValid(index)
	if err != nil {
		return
	}
	element, _, _ = l.removeNode(l.getNodeByIndex(index))
	return
}

// reduceSteps reduces the steps to rotate the linked list by the specified number of steps and keep the result of rotation correct.
func (l LinkedList) reduceSteps(step int) int {
	step %= l.size
	if l.nearTail(step) {
		step = l.reverseIndex(step)
	}
	return step
}

// Rotate rotates the linked list by the specified number of steps and returns the modified linked list.
func (l *LinkedList) Rotate(steps ...int) *LinkedList {
	if l.size <= 1 {
		return l
	}
	step := 1
	if len(steps) >= 1 {
		step = steps[0]
	}
	step = l.reduceSteps(step)
	for step < 0 {
		value, _ := l.PopLeft()
		_ = l.Append(value)
		step++
	}
	for step > 0 {
		value, _ := l.Pop()
		_ = l.AppendLeft(value)
		step--
	}
	return l
}

// Reverse reverses the order of elements in the linked list and returns the modified linked list.
func (l *LinkedList) Reverse() *LinkedList {
	previous := l.head
	node := previous.Next
	for node != nil {
		newNode := node.Next
		node.Next = previous
		previous = node
		node = newNode
	}
	follow := l.tail
	node = follow.Prev
	for node != nil {
		newNode := node.Prev
		node.Prev = follow
		follow = node
		node = newNode
	}
	l.head, l.tail = l.tail, l.head
	return l
}

// ForEach performs the specified action on each element in the linked list and returns the modified linked list.
func (l *LinkedList) ForEach(action func(any) any) *LinkedList {
	node := l.head.Next
	for node != l.tail {
		node.Value = action(node.Value)
		node = node.Next
	}
	return l
}

// Replace replaces all occurrences of the specified old element with the new element and returns the modified linked list.
func (l *LinkedList) Replace(oldElement, newElement any, counts ...int) *LinkedList {
	if common.Equal(oldElement, newElement) {
		return l
	}
	count := l.parseCount(counts...)
	node := l.head.Next
	for count > 0 && node.Next != nil {
		if common.Equal(node.Value, oldElement) {
			node.Value = newElement
			count--
		}
		node = node.Next
	}
	return l
}

// ReplaceRight replaces all occurrences of the specified old element with the new element from the right end and returns the modified linked list.
func (l *LinkedList) ReplaceRight(oldElement, newElement any, counts ...int) *LinkedList {
	if common.Equal(oldElement, newElement) {
		return l
	}
	count := l.parseCount(counts...)
	node := l.tail.Prev
	for count > 0 && node.Prev != nil {
		if common.Equal(node.Value, oldElement) {
			node.Value = newElement
			count--
		}
		node = node.Prev
	}
	return l
}

// Splice removes and/or adds elements in the linked list and returns a new linked list containing the removed elements.
func (l *LinkedList) Splice(start, deleteCount int, items ...any) LinkedList {
	result := NewLinkedList()
	if deleteCount <= 0 {
		return *result
	}
	node := l.getNodeByIndex(l.parseIndex(start))
	var value any
	for node != l.tail {
		if deleteCount == 0 {
			break
		}
		value, _, node = l.removeNode(node)
		_ = result.Append(value)
		deleteCount--
	}
	node = node.Prev
	for _, item := range items {
		node = l.insertNode(node, item)
	}
	return *result
}

// Fill fills the specified area of the linked list with the given element and returns the modified linked list.
func (l *LinkedList) Fill(element any, area ...int) *LinkedList {
	argCount := len(area)
	start := 0
	end := l.size
	if argCount >= 1 {
		start = l.parseIndex(area[0])
	}
	if argCount >= 2 {
		end = l.parseIndex(area[1])
	}
	count := end - start
	if l.size-end < start {
		node := l.getNodeByIndex(end - 1)
		for count > 0 {
			node.Value = element
			count--
			node = node.Prev
		}
	} else {
		node := l.getNodeByIndex(start)
		for count > 0 {
			node.Value = element
			count--
			node = node.Next
		}
	}
	return l
}

// Set sets the element at the specified index in the linked list and returns an error if the index is invalid.
func (l *LinkedList) Set(index int, element any) (err error) {
	if index < 0 {
		index += l.size
	}
	err = l.isIndexValid(index)
	if err != nil {
		return
	}
	l.getNodeByIndex(index).Value = element
	return
}

// RemoveIf removes elements that satisfy the specified condition and returns a new linked list containing the removed elements.
func (l *LinkedList) RemoveIf(condition func(any) bool, counts ...int) LinkedList {
	count := l.parseCount(counts...)
	node := l.head.Next
	result := NewLinkedList()
	var value any
	for count > 0 && node.Next != nil {
		if !condition(node.Value) {
			node = node.Next
			continue
		}
		value, _, node = l.removeNode(node)
		_ = result.Append(value)
		count--
	}
	return *result
}

// RemoveRightIf removes elements from the right end that satisfy the specified condition and returns a new linked list containing the removed elements.
func (l *LinkedList) RemoveRightIf(condition func(any) bool, counts ...int) LinkedList {
	count := l.parseCount(counts...)
	node := l.tail.Prev
	result := NewLinkedList()
	var value any
	for count > 0 && node.Prev != nil {
		if !condition(node.Value) {
			node = node.Prev
			continue
		}
		value, node, _ = l.removeNode(node)
		_ = result.Append(value)
		count--
	}
	return *result
}

// ReplaceIf replaces elements that satisfy the specified condition with the new element and returns a new linked list containing the original elements.
func (l *LinkedList) ReplaceIf(condition func(any) bool, newElement any, counts ...int) LinkedList {
	count := l.parseCount(counts...)
	node := l.head.Next
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

// ReplaceRightIf replaces elements from the right end that satisfy the specified condition with the new element and returns a new linked list containing the original elements.
func (l *LinkedList) ReplaceRightIf(condition func(any) bool, newElement any, counts ...int) LinkedList {
	count := l.parseCount(counts...)
	node := l.tail.Prev
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
