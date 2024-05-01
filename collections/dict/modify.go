package dict

import "flex/common"

func (d *Dict) Clear() *Dict {
	*d = make(Dict)
	return d
}

func (d *Dict) Set(key, value any) *Dict {
	(*d)[key] = value
	return d
}

func (d *Dict) Delete(key any) bool {
	ok := d.Has(key)
	if ok {
		delete(*d, key)
	}
	return ok
}

func (d *Dict) Pop(key any, args ...any) (value any, err error) {
	var defaultVal any
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

func (d *Dict) PopItem() (key, value any, err error) {
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

func (d *Dict) Update(another Dict) *Dict {
	count1 := d.Size()
	count2 := another.Size()
	if common.WillReHash(count1, count2) {
		capacity := common.GetMapInitialCapacity(count1 + count2)
		newDict := make(Dict, capacity)
		for k, v := range *d {
			newDict.Set(k, v)
		}
		*d = newDict
	}
	for k, v := range another {
		d.Set(k, v)
	}
	return d
}
