// Package set provides a set data structure.
package set

import "github.com/biocrosscoder/flex/common"

// Set is a set data structure.
type Set map[any]bool

// Size returns the number of elements in the set.
func (s Set) Size() int {
	return len(s)
}

// Copy creates a new set with the same elements as the original set and returns it.
func (s Set) Copy() Set {
	backup := make(Set, common.GetMapInitialCapacity(s.Size()))
	for k := range s {
		backup.Add(k)
	}
	return backup
}

// Of creates a new set with the provided elements and returns it.
func Of(elements ...any) Set {
	s := make(Set, common.GetMapInitialCapacity(len(elements)))
	for _, element := range elements {
		s.Add(element)
	}
	return s
}
