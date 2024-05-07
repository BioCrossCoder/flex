package collections

import (
	"encoding/json"
	"flex/collections/dict"
	"flex/collections/set"
	"flex/common"
	"fmt"
)

type DefaultDict struct {
	dict.Dict
	builder func() any
}

func NewDefaultDict(items dict.Dict, defaultVal any) *DefaultDict {
	if items == nil {
		items = make(dict.Dict)
	}
	return &DefaultDict{
		items.Copy(),
		func() any {
			return defaultVal
		},
	}
}

func (d *DefaultDict) Get(key any) (value any) {
	if d.Has(key) {
		return d.Dict.Get(key)
	}
	value = d.builder()
	_ = d.Set(key, value)
	return
}

func (d *DefaultDict) Pop(key any) (value any) {
	value, _ = d.Dict.Pop(key, d.builder())
	return
}

func (d DefaultDict) Copy() DefaultDict {
	return DefaultDict{
		d.Dict.Copy(),
		d.builder,
	}
}

func (d *DefaultDict) Clear() *DefaultDict {
	_ = d.Dict.Clear()
	return d
}

func (d *DefaultDict) Set(key, value any) *DefaultDict {
	_ = d.Dict.Set(key, value)
	return d
}

func (d *DefaultDict) Update(another DefaultDict) *DefaultDict {
	_ = d.Dict.Update(another.Dict)
	return d
}

func (d *DefaultDict) SetDefault(value any) *DefaultDict {
	d.builder = func() any {
		return value
	}
	return d
}

func (d DefaultDict) Equal(another DefaultDict) bool {
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

func (d DefaultDict) String() string {
	return fmt.Sprint(d.Dict)
}

func (d DefaultDict) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Dict)
}

func (d *DefaultDict) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &d.Dict)
}
