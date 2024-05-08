package common

import (
	"math"
	"reflect"
	"strings"
	"unicode/utf8"
)

func IsInputFuncValid(f any, inputCount, outputCount int) error {
	fType := reflect.TypeOf(f)
	if fType.Kind() != reflect.Func {
		return ErrNotFunc
	}
	if fType.NumIn() != inputCount {
		return ErrIllegalParamCount
	}
	if fType.NumOut() != outputCount {
		return ErrIllegalReturnCount
	}
	return nil
}

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

func IsList(entry any) error {
	entryType := reflect.TypeOf(entry).Kind()
	for _, v := range ArrayType {
		if entryType == v {
			return nil
		}
	}
	return ErrNotList
}

func IsSequence(entry any) error {
	entryType := reflect.TypeOf(entry).Kind()
	for _, v := range SequenceType {
		if entryType == v {
			return nil
		}
	}
	return ErrNotSeq
}

func IsIterable(entry any) error {
	entryType := reflect.TypeOf(entry).Kind()
	for _, v := range IterableContainers {
		if entryType == v {
			return nil
		}
	}
	return ErrNotIterable
}

func CopyMap(entry reflect.Value, length int) (output map[any]any) {
	output = make(map[any]any, length)
	iter := entry.MapRange()
	for iter.Next() {
		key := iter.Key().Interface()
		value := iter.Value().Interface()
		output[key] = value
	}
	return
}

func CopyList(entry reflect.Value, length int) (output []any) {
	output = make([]any, length)
	for i := 0; i < length; i++ {
		output[i] = entry.Index(i).Interface()
	}
	return
}

func ConvertStringToList(entry string) (output []any) {
	chars := strings.Split(entry, "")
	output = make([]any, len(chars))
	for i, r := range chars {
		output[i] = r
	}
	return
}

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

func GetMapInitialCapacity(elementCount int) int {
	return int(math.Ceil(float64(elementCount) / hashTableFillFactor))
}

func WillReHash(oldElementCount, newElementCount int) bool {
	return float64(newElementCount)/float64(oldElementCount) >= reHashThreshold-1
}

func Equal(a, b any) (equal bool) {
	defer func() {
		if r := recover(); r != nil {
			equal = reflect.DeepEqual(a, b)
		}
	}()
	equal = a == b
	return
}

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

type hasLength interface {
	Len() int
}

type hasSize interface {
	Size() int
}

func Len(entry any) (length int) {
	if l, ok := entry.(hasLength); ok {
		return l.Len()
	}
	if h, ok := entry.(hasSize); ok {
		return h.Size()
	}
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
