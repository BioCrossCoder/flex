package itertools

import (
	"flex/common"
	"reflect"
	"unicode/utf8"
)

type sliceIterator struct {
	entry   []any
	end     int
	pointer int
	value   any
	step    int
}

func NewSliceIterator(entry []any, start, end, step int) Iterator {
	return &sliceIterator{
		entry:   entry,
		end:     end,
		pointer: start,
		value:   nil,
		step:    step,
	}
}

func (iter *sliceIterator) clear() {
	iter.value = nil
	iter.entry = nil
}

func (iter *sliceIterator) Next() bool {
	if (iter.step > 0 && iter.pointer > iter.end) || (iter.step < 0 && iter.pointer < iter.end) {
		iter.clear()
		return false
	}
	iter.value = iter.entry[iter.pointer]
	iter.pointer += iter.step
	return true
}

func (iter *sliceIterator) Value() any {
	return iter.value
}

func (iter *sliceIterator) Pour() any {
	output := make([]any, 0)
	for iter.Next() {
		output = append(output, iter.Value())
	}
	return output
}

func Slice(entry any, start, end, step int) (slice any, err error) {
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
	case reflect.Slice, reflect.Array:
		iterator := NewSliceIterator(common.CopyList(value, length), start, end, step)
		slice = iterator.Pour()
	case reflect.String:
		iterator := NewSliceIterator(common.ConvertStringToList(entry.(string)), start, end, step)
		slice = ""
		for _, c := range iterator.Pour().([]any) {
			slice = slice.(string) + c.(string)
		}
	}
	return
}
