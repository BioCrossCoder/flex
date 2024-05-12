// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
)

// accumulator is an iterator that accumulates the results of applying a function to the elements of an iterable.
type accumulator struct {
	entry   []any
	length  int
	handler func(any, any) any
	pointer int
	value   any
}

// NewAccumulator creates a new accumulator iterator from a slice of elements and a function to apply to each pair of elements.
func NewAccumulator(entry []any, handler func(any, any) any) Iterator {
	return &accumulator{
		entry:   entry,
		length:  len(entry),
		handler: handler,
		pointer: 1,
		value:   entry[0],
	}
}

// clear release the resources used by the iterator.
func (iter *accumulator) clear() {
	iter.value = nil
	iter.entry = nil
	iter.handler = nil
}

// Next moves the pointer to the next element and returns true if there is a next element, or just return false if there is no next element.
func (iter *accumulator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	iter.value = iter.handler(iter.value, iter.entry[iter.pointer])
	iter.pointer++
	return true
}

// Value returns the current value of the iterator element pointed by the pointer.
func (iter *accumulator) Value() any {
	return iter.value
}

// Pour returns the accumulated result of applying the function to all the unvisited elements of the iterator.
func (iter *accumulator) Pour() any {
	result := iter.Value()
	for iter.Next() {
		result = iter.Value()
	}
	return result
}

// Accumulate applies a function to the elements of an iterable and returns an iterator that accumulates the results.
func Accumulate(handler, entry any) (iterator Iterator, err error) {
	err = common.IsInputFuncValid(handler, 2, 1) //nolint
	if err != nil {
		return
	}
	err = common.IsSequence(entry)
	if err != nil {
		return
	}
	value := reflect.ValueOf(entry)
	iterHandler := func(p1, p2 any) any {
		params := []reflect.Value{reflect.ValueOf(p1), reflect.ValueOf(p2)}
		return reflect.ValueOf(handler).Call(params)[0].Interface()
	}
	switch value.Kind() { //nolint
	case reflect.Array, reflect.Slice:
		iterator = NewAccumulator(common.CopyList(value, value.Len()), iterHandler)
	case reflect.String:
		iterator = NewAccumulator(common.ConvertStringToList(value.String()), iterHandler)
	}
	return
}
