// Package sortedlist provides a sorted list data structure.
package sortedlist

import (
	"cmp"
	"github.com/biocrosscoder/flex/typed/collections/arraylist"
	"slices"
)

// SortedList is a type representing a sorted list data structure.
type SortedList[T any] struct {
	elements arraylist.ArrayList[T]
	cmp      func(a, b T) int
}

// AscendOrder is a comparison function that returns the result of comparing two elements in ascending order.
func AscendOrder[T cmp.Ordered](a, b T) int {
	return cmp.Compare(a, b)
}

// DescendOrder is a comparison function that returns the result of comparing two elements in descending order.
func DescendOrder[T cmp.Ordered](a, b T) int {
	return -cmp.Compare(a, b)
}

// NewSortedList creates and returns a new SortedList with the provided comparison function and initial elements.
func NewSortedList[T any](f func(a, b T) int, elements ...T) *SortedList[T] {
	arr := arraylist.Of(elements...)
	if !slices.IsSortedFunc(arr, f) {
		slices.SortFunc(arr, f)
	}
	return &SortedList[T]{arr, f}
}

// Len returns the number of elements in the sorted list.
func (l SortedList[T]) Len() int {
	return l.elements.Len()
}

// Count returns the number of occurrences of the specified element in the sorted list.
func (l SortedList[T]) Count(element T) (count int) {
	index, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	if !exist {
		return
	}
	for i := index; i < l.Len() && l.cmp(l.elements[i], element) == 0; i++ {
		count++
	}
	for i := index - 1; i >= 0 && l.cmp(l.elements[i], element) == 0; i-- {
		count++
	}
	return
}

// Includes checks if the specified element is present in the sorted list.
func (l SortedList[T]) Includes(element T) bool {
	_, exist := slices.BinarySearchFunc(l.elements, element, l.cmp)
	return exist
}

// Empty returns true if the sorted list is empty, otherwise false.
func (l SortedList[T]) Empty() bool {
	return l.elements.Empty()
}

// Copy creates a copy of the sorted list and returns it.
func (l SortedList[T]) Copy() SortedList[T] {
	return SortedList[T]{l.elements.Copy(), l.cmp}
}

// Slice returns a new sorted list that is a subset of the original sorted list based on the provided slice parameters.
func (l SortedList[T]) Slice(args ...int) SortedList[T] {
	f := l.cmp
	if len(args) >= 3 && args[2] < 0 {
		f = func(a, b T) int {
			return -l.cmp(a, b)
		}
	}
	return SortedList[T]{l.elements.Slice(args...), f}
}

// ToReversed returns a new sorted list with the elements in reversed order.
func (l SortedList[T]) ToReversed() SortedList[T] {
	list := l.Copy()
	_ = list.Reverse()
	return list
}

// ToArray returns a copy of the elements in the sorted list as a regular slice.
func (l SortedList[T]) ToArray() []T {
	return l.elements.Copy()
}

// Equal compares the sorted list with another sorted list and returns true if they are equal, otherwise false.
func (l SortedList[T]) Equal(another SortedList[T]) bool {
	return l.elements.Equal(another.elements)
}

// ToList returns a copy of the sorted list's underlying arraylist.
func (l SortedList[T]) ToList() arraylist.ArrayList[T] {
	return l.elements.Copy()
}
