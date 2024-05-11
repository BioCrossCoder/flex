package arraylist

import "github.com/biocrosscoder/flex/common"

func (l ArrayList) Map(handler func(any) any) ArrayList {
	list := make(ArrayList, l.Len())
	for i, item := range l {
		list[i] = handler(item)
	}
	return list
}

func (l ArrayList) Reduce(handler func(any, any) any, initial ...any) (result any, err error) {
	if l.Len() == 0 {
		err = common.ErrEmptyList
		return
	}
	initialCount := len(initial)
	if initialCount > 1 {
		err = common.ErrTooManyArguments
		return
	}
	startIndex := 0
	if initialCount == 0 {
		result = l[startIndex]
		startIndex++
	} else {
		result = initial[0]
	}
	for i := startIndex; i < l.Len(); i++ {
		result = handler(result, l[i])
	}
	return
}

func (l ArrayList) ReduceRight(handler func(any, any) any, initial ...any) (result any, err error) {
	if l.Len() == 0 {
		err = common.ErrEmptyList
		return
	}
	initialCount := len(initial)
	if initialCount > 1 {
		err = common.ErrTooManyArguments
		return
	}
	startIndex := l.Len() - 1
	if initialCount == 0 {
		result = l[startIndex]
		startIndex--
	} else {
		result = initial[0]
	}
	for i := startIndex; i >= 0; i-- {
		result = handler(result, l[i])
	}
	return
}

func (l ArrayList) Filter(condition func(any) bool) ArrayList {
	list := make(ArrayList, 0)
	for _, item := range l {
		if condition(item) {
			list = append(list, item)
		}
	}
	return list
}

func (l ArrayList) Some(condition func(any) bool) bool {
	for _, item := range l {
		if condition(item) {
			return true
		}
	}
	return false
}

func (l ArrayList) Every(condition func(any) bool) bool {
	for _, item := range l {
		if !condition(item) {
			return false
		}
	}
	return true
}
