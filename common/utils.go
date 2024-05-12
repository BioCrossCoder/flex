// Package common contains common constants, functions and errors used throughout the flex codebase.
package common

import (
	"math"
	"reflect"
	"strings"
	"unicode/utf8"
)

// IsInputFuncValid checks if the input parameter is a valid function and if it has the expected number of input and output parameters.
func IsInputFuncValid(f any, inputCount, outputCount int) error {
	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return ErrNotFunc
	}
	if fType.NumIn() != inputCount {
		return ErrUnexpectedParamCount
	}
	if fType.NumOut() != outputCount {
		return ErrUnexpectedReturnCount
	}
	return nil
}

// IsJudgeFunc checks if the input parameter is a function expected to return a boolean value.
func IsJudgeFunc(f any) (err error) {
	err = IsInputFuncValid(f, 1, 1)
	if err != nil {
		return
	}
	if reflect.TypeOf(f).Out(0).Kind() != reflect.Bool {
		err = ErrNotJudgeFunc
	}
	return
}

// IsList checks if the input parameter is array or slice.
func IsList(entry any) error {
	entryType := reflect.TypeOf(entry).Kind()
	for _, v := range ArrayType {
		if entryType == v {
			return nil
		}
	}
	return ErrNotList
}

// IsSequence checks if the input parameter is array, slice or string.
func IsSequence(entry any) error {
	entryType := reflect.TypeOf(entry).Kind()
	for _, v := range SequenceType {
		if entryType == v {
			return nil
		}
	}
	return ErrNotSeq
}

// IsIterable checks if the input parameter is array, slice, map or string.
func IsIterable(entry any) error {
	entryType := reflect.TypeOf(entry).Kind()
	for _, v := range IterableContainers {
		if entryType == v {
			return nil
		}
	}
	return ErrNotIterable
}

// CopyMap copies the map to a new map with the given capcacity.
func CopyMap(entry reflect.Value, capacity int) (output map[any]any) {
	output = make(map[any]any, capacity)
	iter := entry.MapRange()
	for iter.Next() {
		key := iter.Key().Interface()
		value := iter.Value().Interface()
		output[key] = value
	}
	return
}

// CopyList copies the list to a new list with the given length.
func CopyList(entry reflect.Value, length int) (output []any) {
	output = make([]any, length)
	for i := 0; i < length; i++ {
		output[i] = entry.Index(i).Interface()
	}
	return
}

// ConvertStringToList converts a string to a list of characters.
func ConvertStringToList(entry string) (output []any) {
	chars := strings.Split(entry, "")
	output = make([]any, len(chars))
	for i, r := range chars {
		output[i] = r
	}
	return
}

// ConvertMapToLists converts a map to a key list and a value list, and returns the length of the map.
func ConvertMapToLists(entry map[any]any) (keys, values []any, length int) {
	length = len(entry)
	keys = make([]any, length)
	values = make([]any, length)
	i := 0
	for k, v := range entry {
		keys[i] = k
		values[i] = v
		i++
	}
	return
}

// GetMapInitialCapacity calculates the initial capacity of a map based on the expected number of elements.
func GetMapInitialCapacity(elementCount int) int {
	return int(math.Ceil(float64(elementCount) / hashTableLoadFactor))
}

// WillReHash checks if the map will be rehashed based on the expected number of elements.
func WillReHash(oldElementCount, newElementCount int) bool {
	return float64(newElementCount)/float64(oldElementCount) >= reHashThreshold-1
}

// Equal checks if two values are equal, it can safely handle any type of values.
func Equal(a, b any) (equal bool) {
	defer func() {
		if r := recover(); r != nil {
			equal = reflect.DeepEqual(a, b)
		}
	}()
	equal = a == b
	return
}

// ParseIndex parses the index to be within the range of the length.
func ParseIndex(index, length int) int {
	if index < 0 {
		index += length
		if index < 0 {
			return 0
		}
	} else if index >= length {
		return length - 1
	}
	return index
}

// CheckRange checks if the range is valid and within the range of the length.
func CheckRange(start, end, step, length int) (err error) {
	if step == 0 {
		err = ErrZeroStep
		return
	}
	if (start < end && step < 0) || (start > end && step > 0) {
		err = ErrInvalidRange
		return
	}
	if start >= length {
		err = ErrOutOfRange
	}
	return
}

// Len returns the length of the input parameter, it will return -1 if the input parameter is not a valid type.
func Len(entry any) (length int) {
	if s, ok := entry.(string); ok {
		return utf8.RuneCountInString(s)
	}
	defer func() {
		if r := recover(); r != nil {
			length = -1
		}
	}()
	length = reflect.ValueOf(entry).Len()
	return
}

// Contains checks if the input parameter contains the value, it can handle any type of values.
func Contains(entry, value any) bool {
	str, ok1 := entry.(string)
	substr, ok2 := value.(string)
	if ok1 && ok2 {
		return strings.Contains(str, substr)
	}
	if ok1 && (!ok2) {
		return false
	}
	list := reflect.ValueOf(entry)
	switch list.Kind() { //nolint
	case reflect.Array, reflect.Slice:
		for i := 0; i < list.Len(); i++ {
			if Equal(list.Index(i).Interface(), value) {
				return true
			}
		}
	case reflect.Map:
		iter := list.MapRange()
		for iter.Next() {
			if Equal(iter.Value().Interface(), value) {
				return true
			}
		}
	}
	return false
}

// Count counts the number of occurrences of the value in the input parameter, it can handle any type of values.
func Count(entry, value any) (count int) {
	str, ok1 := entry.(string)
	substr, ok2 := value.(string)
	if ok1 && ok2 {
		return strings.Count(str, substr)
	}
	if ok1 && (!ok2) {
		return -1
	}
	list := reflect.ValueOf(entry)
	switch list.Kind() { //nolint
	case reflect.Array, reflect.Slice:
		for i := 0; i < list.Len(); i++ {
			if Equal(list.Index(i).Interface(), value) {
				count++
			}
		}
	case reflect.Map:
		iter := list.MapRange()
		for iter.Next() {
			if Equal(iter.Value().Interface(), value) {
				count++
			}
		}
	default:
		count = -1
	}
	return
}
