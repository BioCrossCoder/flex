package arraylist

import (
	"flex/common"
)

type ArrayList []any

func (l ArrayList) parseIndex(index int) int {
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

func (l ArrayList) isIndexValid(index int) (err error) {
	if index < 0 || index >= l.Len() {
		err = common.ErrOutOfRange
	}
	return
}

func (l ArrayList) Len() int {
	return len(l)
}

func (l ArrayList) Count(element any) (count int) {
	for _, item := range l {
		if item == element {
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
