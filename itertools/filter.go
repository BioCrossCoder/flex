package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

type listFilter struct {
	listConvertor
}

func NewListFilter(entry []any, handler func(any) bool) Iterator {
	iter := new(listFilter)
	iter.entry = entry
	iter.length = len(entry)
	iter.handler = func(p any) any {
		return handler(p)
	}
	iter.pointer = 0
	iter.value = false
	return iter
}

func (iter *listFilter) Pour() any {
	output := make([]any, 0)
	for iter.Next() {
		if iter.Value().(bool) {
			value := iter.entry[iter.pointer-1]
			output = append(output, value)
		}
	}
	return output
}

func Filter(handler, entry any) (iterator Iterator, err error) {
	err = common.IsJudgeFunc(handler)
	if err != nil {
		return
	}
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	length := value.Len()
	iterHandler := func(a any) bool {
		params := []reflect.Value{reflect.ValueOf(a)}
		return reflect.ValueOf(handler).Call(params)[0].Bool()
	}
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		iterator = NewListFilter(common.CopyList(value, length), iterHandler)
	case reflect.String:
		iterator = NewListFilter(common.ConvertStringToList(entry.(string)), iterHandler)
	}
	return
}
