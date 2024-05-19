package collections

import (
	"github.com/biocrosscoder/flex/collections/dict"
	"github.com/biocrosscoder/flex/common"
	"strings"
)

// Translator is a struct that represents a dictionary for translating words or characters.
type Translator struct {
	dict.Dict
}

// NewTranslator creates a new Translator with the provided entries.
func NewTranslator(entries ...[2]string) *Translator {
	d := make(dict.Dict, common.GetMapInitialCapacity(len(entries)))
	for _, entry := range entries {
		_ = d.Set(entry[0], entry[1])
	}
	return &Translator{d}
}

// Clear removes all entries from the Translator.
func (t *Translator) Clear() *Translator {
	t.Dict.Clear()
	return t
}

// Copy creates and returns a copy of the Translator.
func (t Translator) Copy() Translator {
	return Translator{t.Dict.Copy()}
}

// Delete removes the entry with the specified key from the Translator.
func (t *Translator) Delete(key string) bool {
	return t.Dict.Delete(key)
}

// Get retrieves the value associated with the specified key from the Translator.
func (t Translator) Get(key string) (value string) {
	return t.Dict.Get(key, "").(string)
}

// Has checks if the Translator contains the specified key.
func (t Translator) Has(key string) bool {
	return t.Dict.Has(key)
}

// Pop removes the entry with the specified key from the Translator and returns its value.
func (t *Translator) Pop(key string) (value string, err error) {
	raw, err := t.Dict.Pop(key)
	if err == nil {
		value = raw.(string)
	}
	return
}

// PopItem removes and returns an arbitrary entry from the Translator as a key-value pair.
func (t *Translator) PopItem() (key, value string, err error) {
	rawKey, rawValue, err := t.Dict.PopItem()
	if err == nil {
		key = rawKey.(string)
		value = rawValue.(string)
	}
	return
}

// Set adds or updates the entry with the specified key and value in the Translator.
func (t *Translator) Set(key, value string) *Translator {
	_ = t.Dict.Set(key, value)
	return t
}

// Update merges the entries from another Translator into the current Translator.
func (t *Translator) Update(another Translator) *Translator {
	_ = t.Dict.Update(another.Dict)
	return t
}

// Keys returns all keys from the Translator as a slice of strings.
func (t Translator) Keys() []string {
	keys := make([]string, t.Size())
	i := 0
	for k := range t.Dict {
		keys[i] = k.(string)
		i++
	}
	return keys
}

// Values returns all values from the Translator as a slice of strings.
func (t Translator) Values() []string {
	values := make([]string, t.Size())
	i := 0
	for _, v := range t.Dict {
		values[i] = v.(string)
		i++
	}
	return values
}

// Items returns all key-value pairs from the Translator as a slice of [2]string arrays.
func (t Translator) Items() [][2]string {
	items := make([][2]string, t.Size())
	i := 0
	for k, v := range t.Dict {
		items[i] = [2]string{k.(string), v.(string)}
		i++
	}
	return items
}

// Translate translates the input string using the mappings in the Translator and returns the result.
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
