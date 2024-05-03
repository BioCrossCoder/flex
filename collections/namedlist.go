package collections

import (
	"flex/collections/arraylist"
	"flex/collections/dict"
	"flex/common"
)

type NamedList struct {
	fields   []string
	mappings dict.Dict
	elements arraylist.ArrayList
}

func NewNamedList(fields []string) *NamedList {
	count := len(fields)
	mapping := make(dict.Dict, count)
	for i, field := range fields {
		_ = mapping.Set(field, i)
	}
	elements := make(arraylist.ArrayList, count)
	return &NamedList{
		append(make([]string, 0), fields...),
		mapping,
		elements,
	}
}

func (nl NamedList) Fields() []string {
	return nl.fields
}

func (nl NamedList) Elements() arraylist.ArrayList {
	return nl.elements
}

func (nl *NamedList) SetByName(field string, value any) (err error) {
	exist := nl.Contains(field)
	if !exist {
		err = common.ErrKeyNotFound
		return
	}
	index := nl.mappings.Get(field).(int)
	return nl.SetByIndex(index, value)
}

func (nl NamedList) GetByName(field string) (value any, err error) {
	exist := nl.Contains(field)
	if !exist {
		err = common.ErrKeyNotFound
		return
	}
	index := nl.mappings.Get(field).(int)
	return nl.GetByIndex(index)
}

func (nl NamedList) GetByIndex(index int) (value any, err error) {
	return nl.elements.At(index)
}

func (nl *NamedList) SetByIndex(index int, value any) (err error) {
	return nl.elements.Set(index, value)
}

func (nl *NamedList) Add(field string, values ...any) (ok bool) {
	exist := nl.Contains(field)
	if exist {
		return false
	}
	index := nl.elements.Len()
	_ = nl.mappings.Set(field, index)
	nl.fields = append(nl.fields, field)
	if len(values) == 0 {
		values = make([]any, 1)
	}
	_ = nl.elements.Push(values[0])
	return true
}

func (nl *NamedList) Remove(field string) (ok bool) {
	exist := nl.Contains(field)
	if !exist {
		return false
	}
	index := nl.mappings.Get(field).(int)
	_ = nl.mappings.Delete(field)
	_, _ = nl.elements.Pop(index)
	prefix := nl.fields[:index]
	suffix := nl.fields[index+1:]
	for _, field := range suffix {
		oldIndex := nl.mappings.Get(field).(int)
		newIndex := oldIndex - 1
		_ = nl.mappings.Set(field, newIndex)
	}
	nl.fields = append(prefix, suffix...)
	return true
}

func (nl NamedList) Contains(field string) bool {
	return nl.mappings.Has(field)
}

func (nl NamedList) Empty() bool {
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

func (nl NamedList) Count(value any) int {
	return nl.elements.Count(value)
}

func (nl NamedList) Len() int {
	return nl.elements.Len()
}

func (nl NamedList) Copy() NamedList {
	newList := NewNamedList(nl.fields)
	for i, v := range nl.elements {
		_ = newList.SetByIndex(i, v)
	}
	return *newList
}

func (nl NamedList) With(field string, value any) NamedList {
	newList := nl.Copy()
	if !newList.Contains(field) {
		_ = newList.Add(field)
	}
	_ = newList.SetByName(field, value)
	return newList
}

func (nl *NamedList) Update(another NamedList) *NamedList {
	for _, field := range another.Fields() {
		value, _ := another.GetByName(field)
		if !nl.Contains(field) {
			_ = nl.Add(field)
		}
		_ = nl.SetByName(field, value)
	}
	return nl
}

func (nl NamedList) Index(field string) int {
	if nl.Contains(field) {
		return nl.mappings.Get(field).(int)
	}
	return -1
}

func (nl *NamedList) Clear() *NamedList {
	*nl = *NewNamedList(make([]string, 0))
	return nl
}

func (nl *NamedList) Reset() *NamedList {
	_ = nl.elements.Fill(nil)
	return nl
}

func (nl NamedList) Items() arraylist.ArrayList {
	items := make(arraylist.ArrayList, nl.Len())
	for i, field := range nl.fields {
		items[i] = [2]any{field, nl.elements[i]}
	}
	return items
}

func (nl NamedList) Equal(another NamedList) bool {
	length1 := nl.Len()
	length2 := another.Len()
	if length1 != length2 {
		return false
	}
	items1 := nl.Items()
	items2 := another.Items()
	for i := 0; i < length1; i++ {
		if items1[i] != items2[i] {
			return false
		}
	}
	return true
}
