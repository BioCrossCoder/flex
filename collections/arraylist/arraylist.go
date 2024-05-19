// Package arraylist provides a resizable array implementation.
package arraylist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

// ArrayList is an enhanced []any, which provides many convenient methods for array manipulation.
type ArrayList []any

// sliceIndex returns the adjusted index based on collection length and accessibility.
func (l ArrayList) sliceIndex(index int, accessible bool) int {
	return list.SliceIndex(index, l.Len(), accessible)
}

// parseIndex returns the adjusted index based on collection length.
func (l ArrayList) parseIndex(index int) int {
	return list.ParseIndex(index, l.Len())
}

// isIndexValid checks if the index is valid for the collection length.
func (l ArrayList) isIndexValid(index int) error {
	return list.IsIndexValid(index, l.Len())
}

// Len returns the length of the array list.
func (l ArrayList) Len() int {
	return len(l)
}

// Count returns the number of occurrences of an element in the array list.
func (l ArrayList) Count(element any) (count int) {
	for _, item := range l {
		if common.Equal(item, element) {
			count++
		}
	}
	return
}

// Includes checks whether the array list includes a specific element.
func (l ArrayList) Includes(element any) bool {
	return l.IndexOf(element) != -1
}

// Empty checks if the array list is empty.
func (l ArrayList) Empty() bool {
	return l.Len() == 0
}

// Equal checks if two array lists are equal in terms of length and elements.
func (l ArrayList) Equal(another ArrayList) bool {
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
