package arraylist

import "flex/common"

func (l ArrayList) IndexOf(element any) (index int) {
	index = -1
	for i, item := range l {
		if item == element {
			index = i
			break
		}
	}
	return
}

func (l ArrayList) LastIndexOf(element any) (index int) {
	index = -1
	for i := l.Len() - 1; i >= 0; i-- {
		if l[i] == element {
			index = i
			break
		}
	}
	return
}

func (l ArrayList) At(index int) (element any, err error) {
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

func (l ArrayList) Find(by func(any) bool) (element any) {
	for _, item := range l {
		if by(item) {
			element = item
			break
		}
	}
	return
}

func (l ArrayList) FindIndex(by func(any) bool) (index int) {
	index = -1
	for i, item := range l {
		if by(item) {
			index = i
			break
		}
	}
	return
}

func (l ArrayList) FindLast(by func(any) bool) (element any) {
	for i := l.Len() - 1; i >= 0; i-- {
		if by(l[i]) {
			element = l[i]
			break
		}
	}
	return
}

func (l ArrayList) FindLastIndex(by func(any) bool) (index int) {
	index = -1
	for i := l.Len() - 1; i >= 0; i-- {
		if by(l[i]) {
			index = i
			break
		}
	}
	return
}

func (l ArrayList) Head() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l[0]
	return
}

func (l ArrayList) Tail() (element any, err error) {
	if l.Empty() {
		err = common.ErrEmptyList
		return
	}
	element = l[l.Len()-1]
	return
}