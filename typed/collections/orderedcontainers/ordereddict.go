package orderedcontainers

import (
	"encoding/json"
	"flex/common"
	"flex/typed/collections/arraylist"
	"flex/typed/collections/dict"
	"fmt"
	"strings"
)

type OrderedDict[K comparable, V any] struct {
	dict.Dict[K, V]
	sequence arraylist.ArrayList[K]
}

func NewOrderedDict[K comparable, V any]() *OrderedDict[K, V] {
	return &OrderedDict[K, V]{
		make(dict.Dict[K, V]),
		make(arraylist.ArrayList[K], 0),
	}
}

func (d *OrderedDict[K, V]) Clear() *OrderedDict[K, V] {
	_ = d.Dict.Clear()
	_ = d.sequence.Clear()
	return d
}

func (d *OrderedDict[K, V]) Set(key K, value V) *OrderedDict[K, V] {
	if !d.Has(key) {
		_ = d.sequence.Push(key)
	}
	_ = d.Dict.Set(key, value)
	return d
}

func (d *OrderedDict[K, V]) Delete(key K) bool {
	if d.Dict.Delete(key) {
		_ = d.sequence.Remove(key)
		return true
	}
	return false
}

func (d *OrderedDict[K, V]) Pop(key K, args ...V) (value V, err error) {
	if d.Has(key) {
		_ = d.sequence.Remove(key)
	}
	return d.Dict.Pop(key, args...)
}

func (d *OrderedDict[K, V]) PopItem() (key K, value V, err error) {
	key, err = d.sequence.Tail()
	if err != nil {
		err = common.ErrEmptyDict
		return
	}
	value, _ = d.Pop(key)
	return
}

func (d *OrderedDict[K, V]) Update(another OrderedDict[K, V]) *OrderedDict[K, V] {
	for _, key := range another.Keys() {
		_ = d.Set(key, another.Get(key))
	}
	return d
}

func (d OrderedDict[K, V]) Keys() []K {
	return d.sequence.Copy()
}

func (d OrderedDict[K, V]) Values() []V {
	values := make([]V, d.Size())
	for i, key := range d.sequence {
		values[i] = d.Get(key)
	}
	return values
}

func (d OrderedDict[K, V]) Items() []*dict.DictItem[K, V] {
	items := make([]*dict.DictItem[K, V], d.Size())
	for i, key := range d.sequence {
		items[i] = &dict.DictItem[K, V]{key, d.Get(key)}
	}
	return items
}

func (d OrderedDict[K, V]) Copy() OrderedDict[K, V] {
	return OrderedDict[K, V]{
		d.Dict.Copy(),
		d.sequence.Copy(),
	}
}

func (d OrderedDict[K, V]) Equal(another OrderedDict[K, V]) bool {
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

func (d OrderedDict[K, V]) KeyAt(index int) (K, error) {
	return d.sequence.At(index)
}

func (d OrderedDict[K, V]) IndexOf(key K) int {
	return d.sequence.IndexOf(key)
}

func (d OrderedDict[K, V]) String() string {
	items := make([]string, d.Size())
	for i, key := range d.sequence {
		items[i] = fmt.Sprintf("%v:%v", key, d.Get(key))
	}
	return "map[" + strings.Join(items, " ") + "]"
}

func (d OrderedDict[K, V]) MarshalJSON() ([]byte, error) {
	items := make([][2]any, d.Size())
	for i, key := range d.sequence {
		items[i] = [2]any{key, d.Get(key)}
	}
	return json.Marshal(items)
}

func (d *OrderedDict[K, V]) UnmarshalJSON(data []byte) (err error) {
	var items [][2]any
	err = json.Unmarshal(data, &items)
	if err != nil {
		return
	}
	newDict := NewOrderedDict[K, V]()
	for _, item := range items {
		key, ok := item[0].(K)
		if !ok {
			err = common.ErrInvalidType
			return
		}
		value, ok := item[1].(V)
		if !ok {
			err = common.ErrInvalidType
			return
		}
		_ = newDict.Set(key, value)
	}
	*d = *newDict
	return
}
