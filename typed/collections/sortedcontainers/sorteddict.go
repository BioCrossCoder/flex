package sortedcontainers

import (
	"cmp"
	"flex/common"
	"flex/typed/collections/dict"
	"flex/typed/collections/sortedcontainers/sortedlist"
)

type SortedDict[K cmp.Ordered, V any] struct {
	dict.Dict[K, V]
	sequence sortedlist.SortedList[K]
}

func NewSortedDict[K cmp.Ordered, V any](cmp func(a, b K) int, src dict.Dict[K, V]) *SortedDict[K, V] {
	if src == nil {
		src = make(dict.Dict[K, V])
	}
	return &SortedDict[K, V]{
		src,
		*sortedlist.NewSortedList(cmp, src.Keys()...),
	}
}

func (d *SortedDict[K, V]) Clear() *SortedDict[K, V] {
	_ = d.Dict.Clear()
	_ = d.sequence.Clear()
	return d
}

func (d *SortedDict[K, V]) Set(key K, value V) *SortedDict[K, V] {
	if !d.Has(key) {
		_ = d.sequence.Insert(key)
	}
	_ = d.Dict.Set(key, value)
	return d
}

func (d *SortedDict[K, V]) Delete(key K) bool {
	if d.Dict.Delete(key) {
		_ = d.sequence.Remove(key)
		return true
	}
	return false
}

func (d *SortedDict[K, V]) Pop(key K, args ...V) (value V, err error) {
	if d.Has(key) {
		_ = d.sequence.Remove(key)
	}
	return d.Dict.Pop(key, args...)
}

func (d *SortedDict[K, V]) PopItem() (key K, value V, err error) {
	key, err = d.sequence.Tail()
	if err != nil {
		err = common.ErrEmptyDict
		return
	}
	value, _ = d.Pop(key)
	return
}

func (d *SortedDict[K, V]) Update(another dict.Dict[K, V]) *SortedDict[K, V] {
	for k, v := range another {
		_ = d.Set(k, v)
	}
	return d
}

func (d SortedDict[K, V]) Keys() []K {
	return d.sequence.ToArray()
}

func (d SortedDict[K, V]) Values() []V {
	length := d.Size()
	values := make([]V, length)
	for i := 0; i < length; i++ {
		key, _ := d.KeyAt(i)
		values[i] = d.Get(key)
	}
	return values
}

func (d SortedDict[K, V]) Items() []*dict.DictItem[K, V] {
	length := d.Size()
	items := make([]*dict.DictItem[K, V], length)
	for i := 0; i < length; i++ {
		key, _ := d.KeyAt(i)
		value := d.Get(key)
		items[i] = &dict.DictItem[K, V]{key, value}
	}
	return items
}

func (d SortedDict[K, V]) Copy() SortedDict[K, V] {
	return SortedDict[K, V]{
		d.Dict.Copy(),
		d.sequence.Copy(),
	}
}

func (d SortedDict[K, V]) Equal(another SortedDict[K, V]) bool {
	if d.Size() != another.Size() {
		return false
	}
	keys1 := d.Keys()
	keys2 := another.Keys()
	for i := 0; i < d.Size(); i++ {
		key1 := keys1[i]
		key2 := keys2[i]
		if key1 != key2 {
			return false
		}
		if !common.Equal(d.Get(key1), another.Get(key2)) {
			return false
		}
	}
	return true
}

func (d SortedDict[K, V]) KeyAt(index int) (K, error) {
	return d.sequence.At(index)
}

func (d SortedDict[K, V]) IndexOf(key K) int {
	return d.sequence.IndexOf(key)
}
