package dict

import "flex/common"

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

type dictItem[K comparable, V any] struct {
	Key   K
	Value V
}

func (d Dict[K, V]) Items() []*dictItem[K, V] {
	items := make([]*dictItem[K, V], d.Size())
	i := 0
	for k, v := range d {
		items[i] = &dictItem[K, V]{k, v}
		i++
	}
	return items
}

func (d Dict[K, V]) Copy() Dict[K, V] {
	backup := make(Dict[K, V], common.GetMapInitialCapacity(d.Size()))
	for k, v := range d {
		backup.Set(k, v)
	}
	return backup
}

func (d Dict[K, V]) Has(key K) bool {
	_, ok := d[key]
	return ok
}

func (d Dict[K, V]) Empty() bool {
	return d.Size() == 0
}
