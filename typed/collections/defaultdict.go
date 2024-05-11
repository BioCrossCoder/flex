package collections

import (
	"encoding/json"
	"github.com/biocrosscoder/flex/common"
	"github.com/biocrosscoder/flex/typed/collections/dict"
	"fmt"
)

type DefaultDict[K comparable, V any] struct {
	dict.Dict[K, V]
	builder func() V
}

func NewDefaultDict[K comparable, V any](items dict.Dict[K, V], defaultVal V) *DefaultDict[K, V] {
	if items == nil {
		items = make(dict.Dict[K, V])
	}
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
	return d.Dict.Equal(another.Dict) && common.Equal(d.builder(), another.builder())
}

func (d DefaultDict[K, V]) String() string {
	return fmt.Sprint(d.Dict)
}

func (d DefaultDict[K, V]) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Dict)
}

func (d *DefaultDict[K, V]) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &d.Dict)
}
