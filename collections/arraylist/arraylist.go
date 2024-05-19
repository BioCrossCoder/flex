package arraylist

import (
	"github.com/biocrosscoder/flex/collections/list"
	"github.com/biocrosscoder/flex/common"
)

type ArrayList []any

func (l ArrayList) sliceIndex(index int, accessible bool) int {
	return list.SliceIndex(index, l.Len(), accessible)
}

func (l ArrayList) parseIndex(index int) int {
	return list.ParseIndex(index, l.Len())
}

func (l ArrayList) isIndexValid(index int) error {
	return list.IsIndexValid(index, l.Len())
}

func (l ArrayList) Len() int {
	return len(l)
}

func (l ArrayList) Count(element any) (count int) {
	for _, item := range l {
		if common.Equal(item, element) {
			count++
		}
	}
	return
}

func (l ArrayList) Includes(element any) bool {
	return l.IndexOf(element) != -1
}

func (l ArrayList) Empty() bool {
	return l.Len() == 0
}

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
