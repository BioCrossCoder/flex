package linkedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// parseCount parses the count argument for the Remove and RemoveRight methods.
func (l LinkedList[T]) parseCount(counts ...int) int {
	return list.ParseCount(l.Len(), counts...)
}

// removeNode removes a node from the list and returns its value, previous and next nodes.
func (l *LinkedList[T]) removeNode(node *listNode[T]) (value T, prev, next *listNode[T]) {
	value = node.Value
	prev = node.Prev
	next = node.Next
	prev.Next = next
	next.Prev = prev
	l.size--
	return
}

// Remove removes the occurrences of the specified element from the list.
// If the count argument is not specified, only the first occurrence of the element will be removed.
// Returns the modified list.
func (l *LinkedList[T]) Remove(element T, counts ...int) *LinkedList[T] {
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

// RemoveRight removes the occurrences of the specified element from the right end of the list.
// If the count argument is not specified, only the first occurrence of the element will be removed.
// Returns the modified list.
func (l *LinkedList[T]) RemoveRight(element T, counts ...int) *LinkedList[T] {
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

// Clear removes all elements from the list.
func (l *LinkedList[T]) Clear() *LinkedList[T] {
	l.head.Next = l.tail
	l.tail.Prev = l.head
	l.size = 0
	return l
}

// insertNode inserts a new node with the specified value after the specified previous node.
// Returns the new node.
func (l *LinkedList[T]) insertNode(prev *listNode[T], element T) *listNode[T] {
	following := prev.Next
	node := &listNode[T]{
		Value: element,
		Prev:  prev,
		Next:  following,
	}
	prev.Next = node
	following.Prev = node
	l.size++
	return node
}

// Append appends a new element to the end of the list.
// Returns the modified list.
func (l *LinkedList[T]) Append(element T) *LinkedList[T] {
	_ = l.insertNode(l.tail.Prev, element)
	return l
}

// AppendLeft appends a new element to the beginning of the list.
// Returns the modified list.
func (l *LinkedList[T]) AppendLeft(element T) *LinkedList[T] {
	_ = l.insertNode(l.head, element)
	return l
}

// Pop removes and returns the last element of the list.
func (l *LinkedList[T]) Pop() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element, _, _ = l.removeNode(l.tail.Prev)
	return
}

// PopLeft removes and returns the first element of the list.
func (l *LinkedList[T]) PopLeft() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element, _, _ = l.removeNode(l.head.Next)
	return
}

// Extend extends the list by appending all the elements from the specified list.
// Returns the modified list.
func (l *LinkedList[T]) Extend(another *LinkedList[T]) *LinkedList[T] {
	for node := another.head.Next; node != another.tail; node = node.Next {
		_ = l.Append(node.Value)
	}
	return l
}

// ExtendLeft extends the list by prepending all the elements from the specified list.
// Returns the modified list.
func (l *LinkedList[T]) ExtendLeft(another *LinkedList[T]) *LinkedList[T] {
	for node := another.head.Next; node != another.tail; node = node.Next {
		_ = l.AppendLeft(node.Value)
	}
	return l
}

// Insert inserts a new element at the specified index.
// Returns the modified list.
func (l *LinkedList[T]) Insert(index int, element T) *LinkedList[T] {
	_ = l.insertNode(l.getNodeByIndex(l.parseIndex(index)-1), element)
	return l
}

// RemoveByIndex removes the element at the specified index.
// Returns the removed element and an error if the index is out of range.
func (l *LinkedList[T]) RemoveByIndex(index int) (element T, err error) {
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
func (l LinkedList[T]) reduceSteps(step int) int {
	step %= l.size
	if l.nearTail(step) {
		step = l.reverseIndex(step)
	}
	return step
}

// Rotate rotates the list by the specified number of steps.
// If the count argument is not specified, only one step will be performed.
// If the count is negative, the list rotates in a reverse direction.
// Returns the modified list.
func (l *LinkedList[T]) Rotate(steps ...int) *LinkedList[T] {
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

// Reverse reverses the order of the elements in the list.
// Returns the modified list.
func (l *LinkedList[T]) Reverse() *LinkedList[T] {
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

// ForEach applies the specified action function to each element of the list.
// Returns the modified list.
func (l *LinkedList[T]) ForEach(action func(T) T) *LinkedList[T] {
	node := l.head.Next
	for node != l.tail {
		node.Value = action(node.Value)
		node = node.Next
	}
	return l
}

// Replace replaces occurences of the specified old element with the new element.
// If the count argument is not specified, only the first occurrence of the old element will be replaced.
// Returns the modified list.
func (l *LinkedList[T]) Replace(oldElement, newElement T, counts ...int) *LinkedList[T] {
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

// ReplaceRight replaces occurences of the specified old element with the new element from the right end of the list.
// If the count argument is not specified, only the first occurrence of the old element will be replaced.
// Returns the modified list.
func (l *LinkedList[T]) ReplaceRight(oldElement, newElement T, counts ...int) *LinkedList[T] {
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

// Splice removes the specified number of elements from the list starting at the specified index,
// and then inserts the specified elements at the same index.
// Returns the removed elements as a new list.
func (l *LinkedList[T]) Splice(start, deleteCount int, items ...T) LinkedList[T] {
	result := NewLinkedList[T]()
	if deleteCount <= 0 {
		return *result
	}
	node := l.getNodeByIndex(l.parseIndex(start))
	var value T
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

// Fill fills the specified area of the list with the specified element.
// If the area argument is not specified, the entire list will be filled with the specified element.
// Returns the modified list.
func (l *LinkedList[T]) Fill(element T, area ...int) *LinkedList[T] {
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

// Set sets the value of the element at the specified index to the specified element.
// Returns an error if the index is out of range.
func (l *LinkedList[T]) Set(index int, element T) (err error) {
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

// RemoveIf removes the elements that satisfy the specified condition from the list.
// If the count argument is not specified, only the first element that satisfies the condition will be removed.
// Returns the removed elements as a new list.
func (l *LinkedList[T]) RemoveIf(condition func(T) bool, counts ...int) LinkedList[T] {
	count := l.parseCount(counts...)
	node := l.head.Next
	result := NewLinkedList[T]()
	var value T
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

// RemoveRightIf removes the elements that satisfy the specified condition from the right end of the list.
// If the count argument is not specified, only the first element that satisfies the condition will be removed.
// Returns the removed elements as a new list.
func (l *LinkedList[T]) RemoveRightIf(condition func(T) bool, counts ...int) LinkedList[T] {
	count := l.parseCount(counts...)
	node := l.tail.Prev
	result := NewLinkedList[T]()
	var value T
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

// ReplaceIf replaces the elements that satisfy the specified condition with the specified new element.
// If the count argument is not specified, only the first element that satisfies the condition will be replaced.
// Returns the replaced elements as a new list.
func (l *LinkedList[T]) ReplaceIf(condition func(T) bool, newElement T, counts ...int) LinkedList[T] {
	count := l.parseCount(counts...)
	node := l.head.Next
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

// ReplaceRightIf replaces the elements that satisfy the specified condition with the specified new element from the right end of the list.
// If the count argument is not specified, only the first element that satisfies the condition will be replaced.
// Returns the replaced elements as a new list.
func (l *LinkedList[T]) ReplaceRightIf(condition func(T) bool, newElement T, counts ...int) LinkedList[T] {
	count := l.parseCount(counts...)
	node := l.tail.Prev
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
