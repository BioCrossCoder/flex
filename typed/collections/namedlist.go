package collections

import (
	"flex/common"
	"flex/typed/collections/arraylist"
	"flex/typed/collections/dict"
)

type NamedList[T any] struct {
	fields   []string
	mappings dict.Dict[string, int]
	elements arraylist.ArrayList[T]
}

func NewNamedList[T any](fields []string) *NamedList[T] {
	count := len(fields)
	mapping := make(dict.Dict[string, int], count)
	for i, field := range fields {
		_ = mapping.Set(field, i)
	}
	elements := make(arraylist.ArrayList[T], count)
	return &NamedList[T]{
		append(make([]string, 0), fields...),
		mapping,
		elements,
	}
}

func (nl NamedList[T]) Fields() []string {
	return nl.fields
}

func (nl NamedList[T]) Elements() arraylist.ArrayList[T] {
	return nl.elements
}

func (nl *NamedList[T]) SetByName(field string, value T) (err error) {
	exist := nl.Contains(field)
	if !exist {
		err = common.ErrKeyNotFound
		return
	}
	index := nl.mappings.Get(field)
	return nl.SetByIndex(index, value)
}

func (nl NamedList[T]) GetByName(field string) (value T, err error) {
	exist := nl.Contains(field)
	if !exist {
		err = common.ErrKeyNotFound
		return
	}
	index := nl.mappings.Get(field)
	return nl.GetByIndex(index)
}

func (nl NamedList[T]) GetByIndex(index int) (value T, err error) {
	return nl.elements.At(index)
}

func (nl *NamedList[T]) SetByIndex(index int, value T) (err error) {
	return nl.elements.Set(index, value)
}

func (nl *NamedList[T]) Add(field string, values ...T) (ok bool) {
	exist := nl.Contains(field)
	if exist {
		return false
	}
	index := nl.elements.Len()
	_ = nl.mappings.Set(field, index)
	nl.fields = append(nl.fields, field)
	if len(values) == 0 {
		values = make([]T, 1)
	}
	_ = nl.elements.Push(values[0])
	return true
}

func (nl *NamedList[T]) Remove(field string) (ok bool) {
	exist := nl.Contains(field)
	if !exist {
		return false
	}
	index := nl.mappings.Get(field)
	_ = nl.mappings.Delete(field)
	_, _ = nl.elements.Pop(index)
	prefix := nl.fields[:index]
	suffix := nl.fields[index+1:]
	for _, field := range suffix {
		oldIndex := nl.mappings.Get(field)
		newIndex := oldIndex - 1
		_ = nl.mappings.Set(field, newIndex)
	}
	nl.fields = append(prefix, suffix...)
	return true
}

func (nl NamedList[T]) Contains(field string) bool {
	return nl.mappings.Has(field)
}

func (nl NamedList[T]) Empty() bool {
	if !nl.elements.Empty() {
		return false
	}
	if !nl.mappings.Empty() {
		return false
	}
	if len(nl.fields) != 0 {
		return false
	}
	return true
}

func (nl NamedList[T]) Count(value T) int {
	return nl.elements.Count(value)
}

func (nl NamedList[T]) Len() int {
	return nl.elements.Len()
}

func (nl NamedList[T]) Copy() NamedList[T] {
	newList := NewNamedList[T](nl.fields)
	for i, v := range nl.elements {
		_ = newList.SetByIndex(i, v)
	}
	return *newList
}

func (nl NamedList[T]) With(field string, value T) NamedList[T] {
	newList := nl.Copy()
	if !newList.Contains(field) {
		_ = newList.Add(field)
	}
	_ = newList.SetByName(field, value)
	return newList
}

func (nl *NamedList[T]) Update(another NamedList[T]) *NamedList[T] {
	for _, field := range another.Fields() {
		value, _ := another.GetByName(field)
		if !nl.Contains(field) {
			_ = nl.Add(field)
		}
		_ = nl.SetByName(field, value)
	}
	return nl
}

func (nl NamedList[T]) Index(field string) int {
	if nl.Contains(field) {
		return nl.mappings.Get(field)
	}
	return -1
}

func (nl *NamedList[T]) Clear() *NamedList[T] {
	*nl = *NewNamedList[T](make([]string, 0))
	return nl
}

func (nl *NamedList[T]) Reset() *NamedList[T] {
	nl.elements = make(arraylist.ArrayList[T], nl.Len())
	return nl
}

type fieldItem[T any] struct {
	FieldName  string
	FieldValue T
}

func (nl NamedList[T]) Items() arraylist.ArrayList[*fieldItem[T]] {
	items := make(arraylist.ArrayList[*fieldItem[T]], nl.Len())
	for i, field := range nl.fields {
		items[i] = &fieldItem[T]{field, nl.elements[i]}
	}
	return items
}

func (nl NamedList[T]) Equal(another NamedList[T]) bool {
	length1 := nl.Len()
	length2 := another.Len()
	if length1 != length2 {
		return false
	}
	items1 := nl.Items()
	items2 := another.Items()
	for i := 0; i < length1; i++ {
		if items1[i].FieldName != items2[i].FieldName {
			return false
		}
		if !common.Equal(items1[i].FieldValue, items2[i].FieldValue) {
			return false
		}
	}
	return true
}
