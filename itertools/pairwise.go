// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

// pairIterator is an iterator that returns pairs of elements from a slice.
type pairIterator struct {
	entry   []any
	length  int
	pointer int
	value   any
}

// NewPairIterator creates a new pairIterator from a slice.
func NewPairIterator(entry []any) Iterator {
	return &pairIterator{
		entry:   entry,
		length:  len(entry) - 1,
		pointer: 0,
		value:   nil,
	}
}

// clear releases the resources used by the iterator.
func (iter *pairIterator) clear() {
	iter.value = nil
	iter.entry = nil
}

// Next moves the iterator pointer to the next position and returns true if there is a value to return, or false if there are no more values.
func (iter *pairIterator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = [2]any{iter.entry[iter.pointer], iter.entry[iter.pointer+1]}
	iter.pointer++
	return true
}

// Value returns [2]any{currentValue,nextValue} of the slice at the current position.
func (iter *pairIterator) Value() any {
	return iter.value
}

// Pour returns a slice of [2]any{currentValue,nextValue} of the slice from the pointer to the end of the slice.
func (iter *pairIterator) Pour() any {
	length := iter.length - iter.pointer
	output := make([]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}

// PairWise returns an iterator that returns pairs of elements from array, slice or string.
func PairWise(entry any) (iterator Iterator, err error) {
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	switch value.Kind() { //nolint
	case reflect.Slice, reflect.Array:
		iterator = NewPairIterator(common.CopyList(value, value.Len()))
	case reflect.String:
		iterator = NewPairIterator(common.ConvertStringToList(value.String()))
	}
	return
}
