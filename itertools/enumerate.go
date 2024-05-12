// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
	"unicode/utf8"
)

// listEnumerator is an iterator that iterates over a slice.
type listEnumerator struct {
	entry   []any
	end     int
	pointer int
	value   any
	step    int
}

// NewListEnumerator creates a new listEnumerator with the given slice, start, end, and step.
func NewListEnumerator(entry []any, start, end, step int) Iterator {
	return &listEnumerator{
		entry:   entry,
		end:     end,
		pointer: start,
		value:   nil,
		step:    step,
	}
}

// clear release the resources used by the iterator.
func (iter *listEnumerator) clear() {
	iter.value = nil
	iter.entry = nil
}

// Next moves the pointer to the next position and returns true if there is a value to be returned.
func (iter *listEnumerator) Next() bool {
	if (iter.step > 0 && iter.pointer > iter.end) || (iter.step < 0 && iter.pointer < iter.end) {
		iter.clear()
		return false
	}
	iter.value = [2]any{iter.pointer, iter.entry[iter.pointer]}
	iter.pointer += iter.step
	return true
}

// Value returns [2]any{index, value} of the slice at the current position.
func (iter *listEnumerator) Value() any {
	return iter.value
}

// Pour returns a slice of [2]any{index, value} of the slice from the pointer to the end of the slice.
func (iter *listEnumerator) Pour() any {
	output := make([][2]any, 0)
	for iter.Next() {
		output = append(output, iter.Value().([2]any))
	}
	return output
}

// Enumerate creates an iterator that iterates over the given slice, start, end, and step.
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
	switch value.Kind() { //nolint
	case reflect.Array, reflect.Slice:
		iterator = NewListEnumerator(common.CopyList(value, length), start, end, step)
	case reflect.String:
		iterator = NewListEnumerator(common.ConvertStringToList(value.String()), start, end, step)
	}
	return
}
