package itertools

import (
	"flex/common"
	"reflect"
)

type cycler struct {
	entry   []any
	length  int
	pointer int
	value   any
}

func NewCycler(entry []any) *cycler {
	return &cycler{
		entry:   entry,
		length:  len(entry),
		pointer: -1,
		value:   nil,
	}
}

func (iter *cycler) Next() any {
	iter.pointer = (iter.pointer + 1) % iter.length
	iter.value = iter.entry[iter.pointer]
	return iter.value
}

func Cycle(entry any) (iterator *cycler, err error) {
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		iterator = NewCycler(common.CopyList(value, value.Len()))
	case reflect.String:
		iterator = NewCycler(common.ConvertStringToList(entry.(string)))
	}
	return
}
