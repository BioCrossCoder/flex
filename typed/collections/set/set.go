// Package set provides a set data structure.
package set

import "github.com/biocrosscoder/flex/common"

// Set is a generic type representing a set data structure with elements of type T.
type Set[T comparable] map[T]bool

// Size returns the number of elements in the set.
func (s Set[T]) Size() int {
	return len(s)
}

// Copy creates and returns a new set containing the same elements as the original set.
func (s Set[T]) Copy() Set[T] {
	backup := make(Set[T], common.GetMapInitialCapacity(s.Size()))
	for k := range s {
		backup.Add(k)
	}
	return backup
}

// Of creates and returns a new set containing the given elements.
func Of[T comparable](elements ...T) Set[T] {
	s := make(Set[T], common.GetMapInitialCapacity(len(elements)))
	for _, element := range elements {
		s.Add(element)
	}
	return s
}
