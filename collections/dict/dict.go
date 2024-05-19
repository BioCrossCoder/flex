// Package dict provides a flexible hash map implementation.
package dict

// Dict is an enhanced map[any]any that provides additional methods for working with dictionaries.
type Dict map[any]any

// Size returns the number of key-value pairs in the dictionary.
func (d Dict) Size() int {
	return len(d)
}

// Get retrieves the value associated with the specified key. If the key is not found, it returns the default value.
func (d Dict) Get(key any, defaultValue ...any) (value any) {
	if d.Has(key) {
		value = d[key]
	} else if len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return
}

// Keys returns a slice containing all the keys present in the dictionary.
func (d Dict) Keys() []any {
	keys := make([]any, d.Size())
	i := 0
	for k := range d {
		keys[i] = k
		i++
	}
	return keys
}

// Values returns a slice containing all the values present in the dictionary.
func (d Dict) Values() []any {
	values := make([]any, d.Size())
	i := 0
	for _, v := range d {
		values[i] = v
		i++
	}
	return values
}

// Items returns a slice containing all the key-value pairs present in the dictionary.
func (d Dict) Items() [][2]any {
	items := make([][2]any, d.Size())
	i := 0
	for k, v := range d {
		items[i] = [2]any{k, v}
		i++
	}
	return items
}
