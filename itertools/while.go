package itertools

import (
	"flex/common"
	"reflect"
)

func DropWhile(condition func(any) bool, entry any) (iterator common.Iterator, err error) {
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

func TakeWhile(condition func(any) bool, entry any) (iterator common.Iterator, err error) {
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
