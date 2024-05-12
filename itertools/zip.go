// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/biocrosscoder/flex/common"
	"reflect"
	"unicode/utf8"
)

// zipIterator is an iterator that iterates over a list of lists.
type zipIterator struct {
	entries [][]any
	length  int
	pointer int
	value   any
	count   int
}

// NewZipIterator creates a new zipIterator.
func NewZipIterator(entries [][]any, length int) Iterator {
	return &zipIterator{
		entries: entries,
		length:  length,
		pointer: 0,
		value:   nil,
		count:   len(entries),
	}
}

// clear releases the resources used by the iterator.
func (iter *zipIterator) clear() {
	iter.entries = nil
	iter.value = nil
}

// Next moves the iterator to the next position.
func (iter *zipIterator) Next() bool {
	if iter.pointer == iter.length {
		iter.clear()
		return false
	}
	values := make([]any, iter.count)
	for i := 0; i < iter.count; i++ {
		values[i] = iter.entries[i][iter.pointer]
	}
	iter.value = values
	iter.pointer++
	return true
}

// Value returns the current value of the iterator.
func (iter *zipIterator) Value() any {
	return iter.value
}

// Pour returns all remaining values of the iterator.
func (iter *zipIterator) Pour() any {
	length := iter.length - iter.pointer
	output := make([]any, length)
	i := 0
	for iter.Next() {
		output[i] = iter.Value()
		i++
	}
	return output
}

// Zip creates an iterator that iterates over the elements of the given sequences in parallel, and the iteration will end when the shortest input sequence is exhausted.
func Zip(entries ...any) (iterator Iterator, err error) {
	entryCount := len(entries)
	if entryCount < 2 { //nolint
		err = common.ErrUnexpectedParamCount
		return
	}
	entryLength := -1
	iterEntries := make([][]any, entryCount)
	for i, entry := range entries {
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
		if entryLength == -1 {
			entryLength = length
		} else if length < entryLength {
			entryLength = length
		}
		if value.Kind() == reflect.String {
			iterEntries[i] = common.ConvertStringToList(value.String())
		} else {
			iterEntries[i] = common.CopyList(value, length)
		}
	}
	iterator = NewZipIterator(iterEntries, entryLength)
	return
}

// ZipLongest creates an iterator that iterates over the elements of the given sequences in parallel, and the iteration will end when the longest input sequence is exhausted, filling missing values with nil.
func ZipLongest(entries ...any) (iterator Iterator, err error) {
	entryCount := len(entries)
	if entryCount < 2 { //nolint
		err = common.ErrUnexpectedParamCount
		return
	}
	entryLength := 0
	iterEntries := make([][]any, entryCount)
	for _, entry := range entries {
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
		if length > entryLength {
			entryLength = length
		}
	}
	repeater := Repeat(nil)
	for i, entry := range entries {
		value := reflect.ValueOf(entry)
		var list []any
		var tailLength int
		if value.Kind() == reflect.String {
			list = common.ConvertStringToList(value.String())
			tailLength = entryLength - len(list)
		} else {
			list = common.CopyList(value, value.Len())
			tailLength = entryLength - value.Len()
		}
		iterEntries[i] = append(list, repeater.Repeat(tailLength)...)
	}
	iterator = NewZipIterator(iterEntries, entryLength)
	return
}
