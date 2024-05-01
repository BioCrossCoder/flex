package collections

import "flex/collections/dict"

type DefaultDict struct {
	dict.Dict
	builder func() any
}

func NewDefaultDict(items dict.Dict, defaultVal any) *DefaultDict {
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
