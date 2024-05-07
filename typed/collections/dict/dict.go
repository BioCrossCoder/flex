package dict

import (
	"maps"
)

type Dict[K comparable, V any] map[K]V

func (d Dict[K, V]) Size() int {
	return len(d)
}

func (d Dict[K, V]) Get(key K, defaultValue ...V) (value V) {
	if d.Has(key) {
		value = d[key]
	} else if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return
}

func (d Dict[K, V]) Keys() []K {
	keys := make([]K, d.Size())
	i := 0
	for k := range d {
		keys[i] = k
		i++
	}
	return keys
}

func (d Dict[K, V]) Values() []V {
	values := make([]V, d.Size())
	i := 0
	for _, v := range d {
		values[i] = v
		i++
	}
	return values
}

type DictItem[K comparable, V any] struct {
	Key   K
	Value V
}

func (d Dict[K, V]) Items() []*DictItem[K, V] {
	items := make([]*DictItem[K, V], d.Size())
	i := 0
	for k, v := range d {
		items[i] = &DictItem[K, V]{k, v}
		i++
	}
	return items
}

func (d Dict[K, V]) Copy() Dict[K, V] {
	return maps.Clone(d)
}

func (d Dict[K, V]) Has(key K) bool {
	_, ok := d[key]
	return ok
}

func (d Dict[K, V]) Empty() bool {
	return d.Size() == 0
}
