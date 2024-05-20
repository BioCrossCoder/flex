package sortedlist

import (
	"github.com/biocrosscoder/flex/common"
	"slices"
)

// IndexOf returns the index of the first occurrence of the specified element in the sorted list,
// or -1 if the element is not found.
func (l SortedList[T]) IndexOf(element T) int {
	index, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	if !exist {
		return -1
	}
	for index >= 0 && l.cmp(l.elements[index], element) == 0 {
		index--
	}
	return index + 1
}

// LastIndexOf returns the index of the last occurrence of the specified element in the sorted list,
// or -1 if the element is not found.
func (l SortedList[T]) LastIndexOf(element T) int {
	index, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	if !exist {
		return -1
	}
	for index < l.Len() && l.cmp(l.elements[index], element) == 0 {
		index++
	}
	return index - 1
}

// At returns the element at the specified index in the sorted list, or an error if the index is out of range.
func (l SortedList[T]) At(index int) (T, error) {
	return l.elements.At(index)
}

// Find returns the first element satisfying the given predicate function, along with a boolean indicating its existence.
func (l SortedList[T]) Find(by func(T) bool) (T, bool) {
	return l.elements.Find(by)
}

// FindIndex returns the index of the first element satisfying the given predicate function, or -1 if not found.
func (l SortedList[T]) FindIndex(by func(T) bool) int {
	return l.elements.FindIndex(by)
}

// FindLast returns the last element satisfying the given predicate function, along with a boolean indicating its existence.
func (l SortedList[T]) FindLast(by func(T) bool) (T, bool) {
	return l.elements.FindLast(by)
}

// FindLastIndex returns the index of the last element satisfying the given predicate function, or -1 if not found.
func (l SortedList[T]) FindLastIndex(by func(T) bool) int {
	return l.elements.FindLastIndex(by)
}

// Head returns the first element of the sorted list, or an error if the list is empty.
func (l SortedList[T]) Head() (T, error) {
	return l.elements.Head()
}

// Tail returns the last element of the sorted list, or an error if the list is empty.
func (l SortedList[T]) Tail() (T, error) {
	return l.elements.Tail()
}

// Max returns the maximum element in the sorted list, along with an error if the list is empty.
func (l SortedList[T]) Max() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = slices.MaxFunc([]T{l.elements[0], l.elements[l.Len()-1]}, l.cmp)
	return
}

// Min returns the minimum element in the sorted list, along with an error if the list is empty.
func (l SortedList[T]) Min() (element T, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = slices.MinFunc([]T{l.elements[0], l.elements[l.Len()-1]}, l.cmp)
	return
}
