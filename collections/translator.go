package collections

import (
	"flex/collections/dict"
	"flex/common"
	"strings"
)

type Translator struct {
	dict.Dict
}

func NewTranslator(entries ...[2]string) *Translator {
	d := make(dict.Dict, common.GetMapInitialCapacity(len(entries)))
	for _, entry := range entries {
		_ = d.Set(entry[0], entry[1])
	}
	return &Translator{d}
}

func (t *Translator) Clear() *Translator {
	t.Dict.Clear()
	return t
}

func (t Translator) Copy() Translator {
	return Translator{t.Dict.Copy()}
}

func (t *Translator) Delete(key string) bool {
	return t.Dict.Delete(key)
}

func (t Translator) Get(key string) (value string) {
	return t.Dict.Get(key, "").(string)
}

func (t Translator) Has(key string) bool {
	return t.Dict.Has(key)
}

func (t *Translator) Pop(key string) (value string, err error) {
	raw, err := t.Dict.Pop(key)
	if err == nil {
		value = raw.(string)
	}
	return
}

func (t *Translator) PopItem() (key, value string, err error) {
	rawKey, rawValue, err := t.Dict.PopItem()
	if err == nil {
		key = rawKey.(string)
		value = rawValue.(string)
	}
	return
}

func (t *Translator) Set(key, value string) *Translator {
	_ = t.Dict.Set(key, value)
	return t
}

func (t *Translator) Update(another Translator) *Translator {
	_ = t.Dict.Update(another.Dict)
	return t
}

func (t Translator) Keys() []string {
	keys := make([]string, t.Size())
	i := 0
	for k := range t.Dict {
		keys[i] = k.(string)
		i++
	}
	return keys
}

func (t Translator) Values() []string {
	values := make([]string, t.Size())
	i := 0
	for _, v := range t.Dict {
		values[i] = v.(string)
		i++
	}
	return values
}

func (t Translator) Items() [][2]string {
	items := make([][2]string, t.Size())
	i := 0
	for k, v := range t.Dict {
		items[i] = [2]string{k.(string), v.(string)}
		i++
	}
	return items
}

func (t Translator) Translate(entry string) (result string) {
	chars := []rune(entry)
	charCount := len(chars)
	parts := make([]string, charCount)
	for i, c := range chars {
		key := string(c)
		parts[i] = t.Dict.Get(key, key).(string)
	}
	result = strings.Join(parts, "")
	return
}
