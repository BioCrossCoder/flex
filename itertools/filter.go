package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

// listFilter is an iterator that filters a slice.
type listFilter struct {
	listConvertor
}

// NewListFilter creates a new listFilter iterator.
func NewListFilter(entry []any, filter func(any) bool) Iterator {
	iter := new(listFilter)
	iter.entry = entry
	iter.length = len(entry)
	iter.handler = func(p any) any {
		return filter(p)
	}
	iter.pointer = 0
	iter.value = false
	return iter
}

// Pour filters the list and returns a new list containing only the elements that satisfy the filter.
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

// Filter creates an iterator that filters a sequence based on a given function.
func Filter(filter, entry any) (iterator Iterator, err error) {
	err = common.IsJudgeFunc(filter)
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
		return reflect.ValueOf(filter).Call(params)[0].Bool()
	}
	switch value.Kind() { //nolint
	case reflect.Array, reflect.Slice:
		iterator = NewListFilter(common.CopyList(value, length), iterHandler)
	case reflect.String:
		iterator = NewListFilter(common.ConvertStringToList(value.String()), iterHandler)
	}
	return
}
