package set

import "github.com/biocrosscoder/flex/common"

type Set[T comparable] map[T]bool

func (s Set[T]) Size() int {
	return len(s)
}

func (s Set[T]) Copy() Set[T] {
	backup := make(Set[T], common.GetMapInitialCapacity(s.Size()))
	for k := range s {
		backup.Add(k)
	}
	return backup
}

func Of[T comparable](elements ...T) Set[T] {
	s := make(Set[T], common.GetMapInitialCapacity(len(elements)))
	for _, element := range elements {
		s.Add(element)
	}
	return s
}
