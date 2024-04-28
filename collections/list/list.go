package list

import (
	"flex/common"
)

type List []any

func (l List) parseIndex(index int) int {
	length := l.Len()
	if index < 0 {
		index += length
		if index < 0 {
			return 0
		}
	} else if index > length {
		return length
	}
	return index
}

func (l List) isIndexValid(index int) (err error) {
	if index < 0 || index >= l.Len() {
		err = common.ErrOutOfRange
	}
	return
}

func (l List) Len() int {
	return len(l)
}

func (l List) Count(element any) (count int) {
	for _, item := range l {
		if item == element {
			count++
		}
	}
	return
}

func (l List) Includes(element any) bool {
	return l.IndexOf(element) != -1
}

func (l List) Empty() bool {
	return l.Len() == 0
}
