package itertools

import (
	"flex/common"
	"reflect"
)

type listEnumerator struct {
	entry   []any
	end     int
	pointer int
	value   any
	step    int
}

func NewListEnumerator(entry []any, start, end, step int) common.Iterator {
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

func Enumerate(entry any, start, end, step int) (iterator common.Iterator, err error) {
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	if step == 0 {
		err = common.ErrZeroStep
		return
	}
	if (start < end && step < 0) || (start > end && step > 0) {
		err = common.ErrInvalidRange
		return
	}
	value := reflect.ValueOf(entry)
	length := value.Len()
	if value.Kind() == reflect.String {
		length = len([]rune(entry.(string)))
	}
	if start >= length {
		err = common.ErrOutOfRange
		return
	}
	if length-1 < end {
		end = length - 1
	}
	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		iterator = NewListEnumerator(common.CopyList(value, length), start, end, step)
	case reflect.String:
		list := common.ConvertStringToList(entry.(string))
		iterator = NewListEnumerator(list, start, end, step)
	}
	return
}
