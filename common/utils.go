package common

import "reflect"

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

func IsSequence(entry any) error {
	entryType := reflect.TypeOf(entry).Kind()
	for _, v := range Sequence {
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
	chars := []rune(entry)
	output = make([]any, len(chars))
	for i, r := range chars {
		output[i] = string(r)
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
