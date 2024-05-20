package itertools

import "github.com/biocrosscoder/flex/common"

// Enumerator is an interface for iterating over a collection of items.
type Enumerator[T any] interface {
	// Next advances the iterator to the next item in the collection and returns true if there is another item to iterate over.
	Next() bool
	// Value returns the current item in the collection.
	Value() *enumItem[T]
	// Pour iterates over the entire collection and returns all items.
	Pour() []*enumItem[T]
}

// enumItem represents an item in the collection with its index and value.
type enumItem[T any] struct {
	Index int
	Value T
}

// enumerator is a struct that implements the Enumerator interface and provides the iteration functionality.
type enumerator[T any] struct {
	entry   []T
	end     int
	pointer int
	value   *enumItem[T]
	step    int
}

// Enumerate creates and returns an iterator for the given collection with specified start, end, and step parameters.
func Enumerate[T any](entry []T, start, end, step int) (iterator Enumerator[T], err error) {
	length := len(entry)
	err = common.CheckRange(start, end, step, length)
	if err != nil {
		return
	}
	iterator = &enumerator[T]{
		entry:   entry,
		end:     common.ParseIndex(end, length),
		pointer: common.ParseIndex(start, length),
		value:   nil,
		step:    step,
	}
	return
}

// clear resets the iterator's internal state.
func (iter *enumerator[T]) clear() {
	iter.value = nil
	iter.entry = nil
}

// Next advances the iterator to the next item in the collection and returns true if there is another item to iterate over.
func (iter *enumerator[T]) Next() bool {
	if (iter.step > 0 && iter.pointer > iter.end) || (iter.step < 0 && iter.pointer < iter.end) {
		iter.clear()
		return false
	}
	iter.value = &enumItem[T]{
		Index: iter.pointer,
		Value: iter.entry[iter.pointer],
	}
	iter.pointer += iter.step
	return true
}

// Value returns the current item in the collection.
func (iter *enumerator[T]) Value() *enumItem[T] {
	return iter.value
}

// Pour iterates over the entire collection and returns all items.
func (iter *enumerator[T]) Pour() []*enumItem[T] {
	output := make([]*enumItem[T], 0)
	for iter.Next() {
		output = append(output, iter.Value())
	}
	return output
}
