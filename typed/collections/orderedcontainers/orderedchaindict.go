package orderedcontainers

import (
	"encoding/json"
	"flex/common"
	"flex/typed/collections/dict"
	"flex/typed/collections/linkedlist"
	"fmt"
	"strings"
)

type OrderedChainDict[K comparable, V any] struct {
	dict.Dict[K, V]
	sequence *linkedlist.LinkedList[K]
}

func NewOrderedChainDict[K comparable, V any]() *OrderedChainDict[K, V] {
	return &OrderedChainDict[K, V]{
		dict.Dict[K, V]{},
		linkedlist.NewLinkedList[K](),
	}
}

func (d *OrderedChainDict[K, V]) Clear() *OrderedChainDict[K, V] {
	_ = d.Dict.Clear()
	_ = d.sequence.Clear()
	return d
}

func (d *OrderedChainDict[K, V]) Set(key K, value V) *OrderedChainDict[K, V] {
	if !d.Has(key) {
		_ = d.sequence.Append(key)
	}
	_ = d.Dict.Set(key, value)
	return d
}

func (d *OrderedChainDict[K, V]) Delete(key K) bool {
	if d.Has(key) {
		_ = d.Dict.Delete(key)
		_ = d.sequence.Remove(key)
		return true
	}
	return false
}

func (d *OrderedChainDict[K, V]) Pop(key K, args ...V) (value V, err error) {
	if d.Has(key) {
		_ = d.sequence.Remove(key)
	}
	return d.Dict.Pop(key, args...)
}

func (d *OrderedChainDict[K, V]) PopItem() (key K, value V, err error) {
	if d.Empty() {
		err = common.ErrEmptyDict
		return
	}
	key, _ = d.sequence.Tail()
	value, _ = d.Pop(key)
	return
}

func (d *OrderedChainDict[K, V]) Update(another OrderedChainDict[K, V]) *OrderedChainDict[K, V] {
	for _, key := range another.Keys() {
		_ = d.Set(key, another.Get(key))
	}
	return d
}

func (d OrderedChainDict[K, V]) Keys() []K {
	return d.sequence.ToArray()
}

func (d OrderedChainDict[K, V]) Values() []V {
	values := make([]V, d.Size())
	i := 0
	_ = d.sequence.ForEach(func(key K) K {
		values[i] = d.Get(key)
		i++
		return key
	})
	return values
}

func (d OrderedChainDict[K, V]) Items() []*dict.DictItem[K, V] {
	items := make([]*dict.DictItem[K, V], d.Size())
	i := 0
	_ = d.sequence.ForEach(func(key K) K {
		items[i] = &dict.DictItem[K, V]{key, d.Get(key)}
		i++
		return key
	})
	return items
}

func (d OrderedChainDict[K, V]) Copy() OrderedChainDict[K, V] {
	newSeq := d.sequence.Copy()
	return OrderedChainDict[K, V]{
		d.Dict.Copy(),
		&newSeq,
	}
}

func (d OrderedChainDict[K, V]) Equal(another OrderedChainDict[K, V]) bool {
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

func (d OrderedChainDict[K, V]) KeyAt(index int) (K, error) {
	return d.sequence.At(index)
}

func (d OrderedChainDict[K, V]) IndexOf(key K) int {
	return d.sequence.IndexOf(key)
}

func (d OrderedChainDict[K, V]) String() string {
	items := make([]string, d.Size())
	i := 0
	_ = d.sequence.ForEach(func(key K) K {
		items[i] = fmt.Sprintf("%v:%v", key, d.Get(key))
		i++
		return key
	})
	return "map[" + strings.Join(items, " ") + "]"
}

func (d OrderedChainDict[K, V]) MarshalJSON() ([]byte, error) {
	items := make([][2]any, d.Size())
	i := 0
	_ = d.sequence.ForEach(func(key K) K {
		items[i] = [2]any{key, d.Get(key)}
		i++
		return key
	})
	return json.Marshal(items)
}

func (d *OrderedChainDict[K, V]) UnmarshalJSON(data []byte) (err error) {
	var items [][2]any
	err = json.Unmarshal(data, &items)
	if err != nil {
		return
	}
	newDict := NewOrderedChainDict[K, V]()
	dictItem := dict.DictItem[K, V]{}
	var rawKey, rawValue []byte
	for _, item := range items {
		rawKey, err = json.Marshal(item[0])
		if err != nil {
			return
		}
		err = json.Unmarshal(rawKey, &dictItem.Key)
		if err != nil {
			return
		}
		rawValue, err = json.Marshal(item[1])
		if err != nil {
			return
		}
		err = json.Unmarshal(rawValue, &dictItem.Value)
		if err != nil {
			return
		}
		_ = newDict.Set(dictItem.Key, dictItem.Value)
	}
	*d = *newDict
	return
}
