package arraylist

import "github.com/biocrosscoder/flex/common"

// Map applies the given handler function to each item in the ArrayList and returns a new ArrayList of the results.
func (l ArrayList[T]) Map(handler func(T) T) ArrayList[T] {
	list := make(ArrayList[T], l.Len())
	for i, item := range l {
		list[i] = handler(item)
	}
	return list
}

// Reduce reduces the ArrayList to a single value by applying the handler function cumulatively to each item, starting with the initial value if provided.
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

// ReduceRight reduces the ArrayList to a single value by applying the handler function cumulatively from right to left, starting with the initial value if provided.
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

// Filter creates a new ArrayList with all items that pass the condition function.
func (l ArrayList[T]) Filter(condition func(T) bool) ArrayList[T] {
	list := make(ArrayList[T], 0)
	for _, item := range l {
		if condition(item) {
			list = append(list, item)
		}
	}
	return list
}

// Some checks if at least one item in the ArrayList satisfies the condition function.
func (l ArrayList[T]) Some(condition func(T) bool) bool {
	for _, item := range l {
		if condition(item) {
			return true
		}
	}
	return false
}

// Every checks if every item in the ArrayList satisfies the condition function.
func (l ArrayList[T]) Every(condition func(T) bool) bool {
	for _, item := range l {
		if !condition(item) {
			return false
		}
	}
	return true
}
