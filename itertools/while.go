// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

// DropWhile returns an iterator that drops elements from the iterable as long as the condition is true.
func DropWhile(condition func(any) bool, entry any) (iterator Iterator, err error) {
	err = common.IsList(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	length := value.Len()
	iterator = NewListIterator(make([]any, 0))
	for i := 0; i < length; i++ {
		if !condition(value.Index(i).Interface()) {
			iterator = NewListIterator(common.CopyList(value.Slice(i, length), length-i))
			break
		}
	}
	return
}

// TakeWhile returns an iterator that takes elements from the iterable as long as the condition is true.
func TakeWhile(condition func(any) bool, entry any) (iterator Iterator, err error) {
	err = common.IsList(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	length := value.Len()
	for i := 0; i < length; i++ {
		if !condition(value.Index(i).Interface()) {
			iterator = NewListIterator(common.CopyList(value.Slice(0, i), i))
			break
		}
	}
	if iterator == nil {
		iterator = NewListIterator(common.CopyList(value, length))
	}
	return
}
