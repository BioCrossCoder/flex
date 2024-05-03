package dict

import (
	"flex/collections/arraylist"
)

type Dict map[any]any

func (d Dict) Size() int {
	return len(d)
}

func (d Dict) Get(key any, defaultValue ...any) (value any) {
	if d.Has(key) {
		value = d[key]
	} else if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return
}

func (d Dict) Keys() arraylist.ArrayList {
	keys := make(arraylist.ArrayList, d.Size())
	i := 0
	for k := range d {
		keys[i] = k
		i++
	}
	return keys
}

func (d Dict) Values() arraylist.ArrayList {
	values := make(arraylist.ArrayList, d.Size())
	i := 0
	for _, v := range d {
		values[i] = v
		i++
	}
	return values
}

func (d Dict) Items() arraylist.ArrayList {
	items := make(arraylist.ArrayList, d.Size())
	i := 0
	for k, v := range d {
		items[i] = [2]any{k, v}
		i++
	}
	return items
}
