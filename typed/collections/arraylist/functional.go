package arraylist

import "flex/common"

func (l ArrayList[T]) Map(handler func(T) T) ArrayList[T] {
	list := make(ArrayList[T], l.Len())
	for i, item := range l {
		list[i] = handler(item)
	}
	return list
}

func (l ArrayList[T]) Reduce(handler func(T, T) T, initial ...T) (result T, err error) {
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

func (l ArrayList[T]) ReduceRight(handler func(T, T) T, initial ...T) (result T, err error) {
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

func (l ArrayList[T]) Filter(condition func(T) bool) ArrayList[T] {
	list := make(ArrayList[T], 0)
	for _, item := range l {
		if condition(item) {
			list = append(list, item)
		}
	}
	return list
}

func (l ArrayList[T]) Some(condition func(T) bool) bool {
	for _, item := range l {
		if condition(item) {
			return true
		}
	}
	return false
}

func (l ArrayList[T]) Every(condition func(T) bool) bool {
	for _, item := range l {
		if !condition(item) {
			return false
		}
	}
	return true
}
