package arraylist

import "github.com/biocrosscoder/flex/common"

// Map applies the given handler function to each item in the list and returns a new list containing the results.
func (l ArrayList) Map(handler func(any) any) ArrayList {
	list := make(ArrayList, l.Len())
	for i, item := range l {
		list[i] = handler(item)
	}
	return list
}

// Reduce reduces the list to a single value by applying the handler function cumulatively to the items. It returns the accumulated result value and an error.
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

// ReduceRight reduces the list to a single value by applying the handler function cumulatively from right to left. It returns the accumulated result value and an error.
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

// Filter creates a new list with all items that pass the condition function.
func (l ArrayList) Filter(condition func(any) bool) ArrayList {
	list := make(ArrayList, 0)
	for _, item := range l {
		if condition(item) {
			list = append(list, item)
		}
	}
	return list
}

// Some checks if at least one item in the list satisfies the given condition function.
func (l ArrayList) Some(condition func(any) bool) bool {
	for _, item := range l {
		if condition(item) {
			return true
		}
	}
	return false
}

// Every checks if all items in the list satisfy the given condition function.
func (l ArrayList) Every(condition func(any) bool) bool {
	for _, item := range l {
		if !condition(item) {
			return false
		}
	}
	return true
}
