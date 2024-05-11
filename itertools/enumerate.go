package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
	"unicode/utf8"
)

type listEnumerator struct {
	entry   []any
	end     int
	pointer int
	value   any
	step    int
}

func NewListEnumerator(entry []any, start, end, step int) Iterator {
	return &listEnumerator{
		entry:   entry,
		end:     end,
		pointer: start,
		value:   nil,
		step:    step,
	}
}

func (iter *listEnumerator) clear() {
	iter.value = nil
	iter.entry = nil
}

func (iter *listEnumerator) Next() bool {
	if (iter.step > 0 && iter.pointer > iter.end) || (iter.step < 0 && iter.pointer < iter.end) {
		iter.clear()
		return false
	}
	iter.value = [2]any{iter.pointer, iter.entry[iter.pointer]}
	iter.pointer += iter.step
	return true
}

func (iter *listEnumerator) Value() any {
	return iter.value
}

func (iter *listEnumerator) Pour() any {
	output := make([][2]any, 0)
	for iter.Next() {
		output = append(output, iter.Value().([2]any))
	}
	return output
}

func Enumerate(entry any, start, end, step int) (iterator Iterator, err error) {
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	var length int
	if value.Kind() == reflect.String {
		length = utf8.RuneCountInString(entry.(string))
	} else {
		length = value.Len()
	}
	err = common.CheckRange(start, end, step, length)
	if err != nil {
		return
	}
	start = common.ParseIndex(start, length)
	end = common.ParseIndex(end, length)
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		iterator = NewListEnumerator(common.CopyList(value, length), start, end, step)
	case reflect.String:
		iterator = NewListEnumerator(common.ConvertStringToList(entry.(string)), start, end, step)
	}
	return
}
