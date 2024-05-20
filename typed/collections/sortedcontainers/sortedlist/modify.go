package sortedlist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"slices"
)

// parseCount returns the count of elements to be parsed based on the length of the sorted list and the given counts.
func (l SortedList[T]) parseCount(counts ...int) int {
	return list.ParseCount(l.Len(), counts...)
}

// Remove removes the specified element from the sorted list, according to the given count, and returns the modified list.
func (l *SortedList[T]) Remove(element T, counts ...int) *SortedList[T] {
	count := l.parseCount(counts...)
	index, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	if !exist {
		return l
	}
	for index >= 0 && l.cmp(l.elements[index], element) == 0 {
		index--
	}
	index++
	i := index
	for count > 0 && i < l.Len() && l.cmp(l.elements[i], element) == 0 {
		count--
		i++
	}
	_ = l.RemoveRange(index, i)
	return l
}

// RemoveRange removes a range of elements from the sorted list and returns a new sorted list.
func (l *SortedList[T]) RemoveRange(start, end int) SortedList[T] {
	return SortedList[T]{
		l.elements.Splice(start, end-start),
		l.cmp,
	}
}

// Clear removes all elements from the sorted list and returns the modified list.
func (l *SortedList[T]) Clear() *SortedList[T] {
	_ = l.elements.Clear()
	return l
}

// Pop removes and returns the element at the specified indexes from the sorted list.
func (l *SortedList[T]) Pop(indexes ...int) (element T, err error) {
	return l.elements.Pop(indexes...)
}

// Insert inserts an element into the sorted list at the appropriate position and returns the modified list.
func (l *SortedList[T]) Insert(element T) *SortedList[T] {
	index, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	if exist {
		for index < l.Len() && l.cmp(l.elements[index], element) == 0 {
			index++
		}
	}
	_ = l.elements.Insert(index, element)
	return l
}

// Reverse reverses the order of elements in the sorted list and returns the modified list.
func (l *SortedList[T]) Reverse() *SortedList[T] {
	_ = l.elements.Reverse()
	l.cmp = func(a, b T) int {
		return -l.cmp(a, b)
	}
	return l
}

// RemoveIf removes elements from the sorted list based on the specified condition and returns a new sorted list.
func (l *SortedList[T]) RemoveIf(condition func(T) bool, counts ...int) SortedList[T] {
	return SortedList[T]{
		l.elements.RemoveIf(condition, counts...),
		l.cmp,
	}
}

// RemoveRightIf removes elements from the end of the sorted list based on the specified condition and returns a new sorted list.
func (l *SortedList[T]) RemoveRightIf(condition func(T) bool, counts ...int) SortedList[T] {
	return SortedList[T]{
		l.elements.RemoveRightIf(condition, counts...),
		l.cmp,
	}
}
