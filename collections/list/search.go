package list

import "flex/common"

func (l List) IndexOf(element any) (index int) {
	index = -1
	for i, item := range l {
		if item == element {
			index = i
			break
		}
	}
	return
}

func (l List) LastIndexOf(element any) (index int) {
	index = -1
	for i := l.Len() - 1; i >= 0; i-- {
		if l[i] == element {
			index = i
			break
		}
	}
	return
}

func (l List) At(index int) (element any, err error) {
	if index < 0 {
		index += l.Len()
	}
	err = l.isIndexValid(index)
	if err != nil {
		return
	}
	element = l[index]
	return
}

func (l List) Find(by func(any) bool) (element any) {
	for _, item := range l {
		if by(item) {
			element = item
			break
		}
	}
	return
}

func (l List) FindIndex(by func(any) bool) (index int) {
	index = -1
	for i, item := range l {
		if by(item) {
			index = i
			break
		}
	}
	return
}

func (l List) FindLast(by func(any) bool) (element any) {
	for i := l.Len() - 1; i >= 0; i-- {
		if by(l[i]) {
			element = l[i]
			break
		}
	}
	return
}

func (l List) FindLastIndex(by func(any) bool) (index int) {
	index = -1
	for i := l.Len() - 1; i >= 0; i-- {
		if by(l[i]) {
			index = i
			break
		}
	}
	return
}

func (l List) Head() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l[0]
	return
}

func (l List) Tail() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l[l.Len()-1]
	return
}
