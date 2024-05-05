package collections

import (
	"flex/common"
	"flex/typed/collections/dict"
	"flex/typed/collections/set"
)

type DefaultDict[K comparable, V any] struct {
	dict.Dict[K, V]
	builder func() V
}

func NewDefaultDict[K comparable, V any](items dict.Dict[K, V], defaultVal V) *DefaultDict[K, V] {
	return &DefaultDict[K, V]{
		items.Copy(),
		func() V {
			return defaultVal
		},
	}
}

func (d *DefaultDict[K, V]) Get(key K) (value V) {
	if d.Has(key) {
		return d.Dict.Get(key)
	}
	value = d.builder()
	_ = d.Set(key, value)
	return
}

func (d *DefaultDict[K, V]) Pop(key K) (value V) {
	value, _ = d.Dict.Pop(key, d.builder())
	return
}

func (d DefaultDict[K, V]) Copy() DefaultDict[K, V] {
	return DefaultDict[K, V]{
		d.Dict.Copy(),
		d.builder,
	}
}

func (d *DefaultDict[K, V]) Clear() *DefaultDict[K, V] {
	_ = d.Dict.Clear()
	return d
}

func (d *DefaultDict[K, V]) Set(key K, value V) *DefaultDict[K, V] {
	_ = d.Dict.Set(key, value)
	return d
}

func (d *DefaultDict[K, V]) Update(another DefaultDict[K, V]) *DefaultDict[K, V] {
	_ = d.Dict.Update(another.Dict)
	return d
}

func (d *DefaultDict[K, V]) SetDefault(value V) *DefaultDict[K, V] {
	d.builder = func() V {
		return value
	}
	return d
}

func (d DefaultDict[K, V]) Equal(another DefaultDict[K, V]) bool {
	keys1 := set.Of(d.Keys()...)
	keys2 := set.Of(another.Keys()...)
	if !keys1.SymmetricDifference(keys2).Empty() {
		return false
	}
	for k, v := range d.Dict {
		if !common.Equal(v, another.Get(k)) {
			return false
		}
	}
	return true
}
