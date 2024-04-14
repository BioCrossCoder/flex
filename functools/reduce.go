package functools

import (
	"flex/common"
	"flex/itertools"
	"reflect"
)

func Reduce(handler, entry any) (output any, err error) {
	iterator, err := itertools.Accumulate(handler, entry)
	if err != nil {
		return
	}
	output = iterator.Pour()
	return
}

func ReduceMap(handler, entry any) (output any, err error) {
	entryValue := reflect.ValueOf(entry)
	entryType := entryValue.Type()
	if entryType.Kind() != reflect.Map {
		err = common.ErrNotMap
		return
	}
	values := make([]any, 0)
	iter := entryValue.MapRange()
	for iter.Next() {
		value := iter.Value().Interface()
		values = append(values, value)
	}
	return Reduce(handler, values)
}
