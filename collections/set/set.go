package set

import "github.com/biocrosscoder/flex/common"

type Set map[any]bool

func (s Set) Size() int {
	return len(s)
}

func (s Set) Copy() Set {
	backup := make(Set, common.GetMapInitialCapacity(s.Size()))
	for k := range s {
		backup.Add(k)
	}
	return backup
}

func Of(elements ...any) Set {
	s := make(Set, common.GetMapInitialCapacity(len(elements)))
	for _, element := range elements {
		s.Add(element)
	}
	return s
}
