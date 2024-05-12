// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
	"unicode/utf8"
)

// sliceIterator is an iterator for a slice.
type sliceIterator struct {
	entry   []any
	end     int
	pointer int
	value   any
	step    int
}

// NewSliceIterator creates a new slice iterator with the given slice, start, end, and step.
func NewSliceIterator(entry []any, start, end, step int) Iterator {
	return &sliceIterator{
		entry:   entry,
		end:     end,
		pointer: start,
		value:   nil,
		step:    step,
	}
}

// clear releases the resources used by the iterator.
func (iter *sliceIterator) clear() {
	iter.value = nil
	iter.entry = nil
}

// Next moves the pointer to the next position and returns true if there is a valid value or false otherwise.
func (iter *sliceIterator) Next() bool {
	if (iter.step > 0 && iter.pointer > iter.end) || (iter.step < 0 && iter.pointer < iter.end) {
		iter.clear()
		return false
	}
	iter.value = iter.entry[iter.pointer]
	iter.pointer += iter.step
	return true
}

// Value returns the current value pointed by the pointer.
func (iter *sliceIterator) Value() any {
	return iter.value
}

// Pour returns a slice containing all the unvisited values in the iterator.
func (iter *sliceIterator) Pour() any {
	output := make([]any, 0)
	for iter.Next() {
		output = append(output, iter.Value())
	}
	return output
}

// Slice returns a slice of the given entry, with the given start, end, and step from array, slice, or string.
func Slice(entry any, start, end, step int) (slice any, err error) {
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	var length int
	if value.Kind() == reflect.String {
		length = utf8.RuneCountInString(value.String())
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
	case reflect.Slice, reflect.Array:
		iterator := NewSliceIterator(common.CopyList(value, length), start, end, step)
		slice = iterator.Pour()
	case reflect.String:
		iterator := NewSliceIterator(common.ConvertStringToList(value.String()), start, end, step)
		slice = ""
		for _, c := range iterator.Pour().([]any) {
			slice = slice.(string) + c.(string)
		}
	}
	return
}
