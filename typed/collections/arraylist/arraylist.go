// Package arraylist provides a resizable array implementation.
package arraylist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// ArrayList is a generic type representing a list of elements
type ArrayList[T any] []T

// sliceIndex returns the adjusted index based on the list length and accessibility requirement
func (l ArrayList[T]) sliceIndex(index int, accessible bool) int {
	return list.SliceIndex(index, l.Len(), accessible)
}

// parseIndex returns the adjusted index based on the list length
func (l ArrayList[T]) parseIndex(index int) int {
	return list.ParseIndex(index, l.Len())
}

// isIndexValid checks if the index is valid for the list
func (l ArrayList[T]) isIndexValid(index int) error {
	return list.IsIndexValid(index, l.Len())
}

// Len returns the length of the list
func (l ArrayList[T]) Len() int {
	return len(l)
}

// Count returns the number of occurrences of the specified element in the list
func (l ArrayList[T]) Count(element T) (count int) {
	for _, item := range l {
		if common.Equal(item, element) {
			count++
		}
	}
	return
}

// Includes checks if the specified element is present in the list
func (l ArrayList[T]) Includes(element T) bool {
	return l.IndexOf(element) != -1
}

// Empty checks if the list is empty
func (l ArrayList[T]) Empty() bool {
	return l.Len() == 0
}

// Equal checks if two lists are equal by comparing their elements
func (l ArrayList[T]) Equal(another ArrayList[T]) bool {
	if l.Len() != another.Len() {
		return false
	}
	for i := 0; i < l.Len(); i++ {
		if !common.Equal(l[i], another[i]) {
			return false
		}
	}
	return true
}
