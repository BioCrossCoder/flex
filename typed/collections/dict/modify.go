package dict

import (
	"flex/common"
	"maps"
)

func (d *Dict[K, V]) Clear() *Dict[K, V] {
	*d = make(Dict[K, V])
	return d
}

func (d *Dict[K, V]) Set(key K, value V) *Dict[K, V] {
	(*d)[key] = value
	return d
}

func (d *Dict[K, V]) Delete(key K) bool {
	ok := d.Has(key)
	if ok {
		delete(*d, key)
	}
	return ok
}

func (d *Dict[K, V]) Pop(key K, args ...V) (value V, err error) {
	var defaultVal V
	argCount := len(args)
	if !d.Has(key) && argCount == 0 {
		err = common.ErrKeyNotFound
		return
	}
	if argCount >= 1 {
		defaultVal = args[0]
	}
	value = d.Get(key, defaultVal)
	_ = d.Delete(key)
	return
}

func (d *Dict[K, V]) PopItem() (key K, value V, err error) {
	if d.Empty() {
		err = common.ErrEmptyDict
		return
	}
	for k, v := range *d {
		key = k
		value = v
		_ = d.Delete(k)
		break
	}
	return
}

func (d *Dict[K, V]) Update(another Dict[K, V]) *Dict[K, V] {
	maps.Copy(*d, another)
	return d
}
