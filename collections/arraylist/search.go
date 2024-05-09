package arraylist

import "flex/common"

func (l ArrayList) IndexOf(element any) (index int) {
	index = -1
	for i, item := range l {
		if common.Equal(item, element) {
			index = i
			break
		}
	}
	return
}

func (l ArrayList) LastIndexOf(element any) (index int) {
	index = -1
	for i := l.Len() - 1; i >= 0; i-- {
		if common.Equal(l[i], element) {
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

func (l ArrayList) Find(by func(any) bool) (element any, found bool) {
	index := l.FindIndex(by)
	if index != -1 {
		found = true
		element = l[index]
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

func (l ArrayList) FindLast(by func(any) bool) (element any, found bool) {
	index := l.FindLastIndex(by)
	if index != -1 {
		found = true
		element = l[index]
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

func (l ArrayList) FindIndexes(by func(any) bool, counts ...int) (indexes []int) {
	count := l.searchCount(counts...)
	indexes = make([]int, 0)
	for i, item := range l {
		if count == 0 {
			break
		}
		if by(item) {
			indexes = append(indexes, i)
			count--
		}
	}
	return
}

func (l ArrayList) FindLastIndexes(by func(any) bool, counts ...int) (indexes []int) {
	count := l.searchCount(counts...)
	indexes = make([]int, 0)
	for i := l.Len() - 1; i >= 0; i-- {
		if count == 0 {
			break
		}
		if by(l[i]) {
			indexes = append(indexes, i)
			count--
		}
	}
	return
}

func (l ArrayList) Finds(by func(any) bool, counts ...int) (elements []any) {
	count := l.searchCount(counts...)
	elements = make([]any, 0)
	for _, item := range l {
		if count == 0 {
			break
		}
		if by(item) {
			elements = append(elements, item)
			count--
		}
	}
	return
}

func (l ArrayList) FindLasts(by func(any) bool, counts ...int) (elements []any) {
	count := l.searchCount(counts...)
	elements = make([]any, 0)
	for i := l.Len() - 1; i >= 0; i-- {
		if count == 0 {
			break
		}
		if by(l[i]) {
			elements = append(elements, l[i])
			count--
		}
	}
	return
}

func (l ArrayList) searchCount(counts ...int) int {
	if len(counts) == 0 || counts[0] <= 0 {
		return l.Len()
	}
	return counts[0]
}
