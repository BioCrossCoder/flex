package sortedcontainers

import (
	"github.com/biocrosscoder/flex/typed/collections/dict"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrderedDict(t *testing.T) {
	src := dict.Dict[string, int]{
		"d": 1,
		"a": 2,
		"b": 3,
		"c": 4,
	}
	d := NewSortedDict[string, int](nil, nil)
	for k, v := range src {
		_ = d.Set(k, v)
	}
	assert.Equal(t, d.Keys(), []string{"a", "b", "c", "d"})
	size := d.Size()
	convey.Convey("add a key-value pair to dict", t, func() {
		d2 := d.Copy()
		assert.Equal(t, d2.Size(), size)
		assert.False(t, d2.Has("e"))
		_ = d2.Set("e", 5)
		assert.True(t, d2.Has("e"))
		assert.Equal(t, d2.Size(), size+1)
		assert.Equal(t, d2.Get("e"), 5)
		assert.Equal(t, d2.IndexOf("e"), size)
		key, err := d2.KeyAt(size)
		assert.Nil(t, err)
		assert.Equal(t, key, "e")
	})
	convey.Convey("remove a key from dict", t, func() {
		d2 := d.Copy()
		assert.Equal(t, d2.Size(), size)
		assert.True(t, d2.Has("b"))
		assert.True(t, d2.Delete("b"))
		assert.False(t, d2.Has("b"))
		assert.Equal(t, d2.Size(), size-1)
	})
	convey.Convey("pop a key from dict", t, func() {
		d2 := d.Copy()
		assert.Equal(t, d2.Size(), size)
		assert.True(t, d2.Has("a"))
		expected := d2.Get("a")
		value, err := d2.Pop("a")
		assert.Nil(t, err)
		assert.Equal(t, expected, value)
		assert.Equal(t, d2.Size(), size-1)
		assert.False(t, d2.Has("a"))
	})
	convey.Convey("pop a key-value pair from dict", t, func() {
		d2 := d.Copy()
		assert.Equal(t, d2.Size(), size)
		expectedKey := d2.Keys()[d2.Size()-1]
		expectedValue := d2.Get(expectedKey)
		key, value, err := d2.PopItem()
		assert.Nil(t, err)
		assert.Equal(t, key, expectedKey)
		assert.Equal(t, value, expectedValue)
	})
	convey.Convey("iterate over dict", t, func() {
		d2 := d.Copy()
		keys := d2.Keys()
		values := d2.Values()
		items := d2.Items()
		for i := 0; i < d2.Size(); i++ {
			assert.Equal(t, keys[i], items[i].Key)
			assert.Equal(t, values[i], items[i].Value)
		}
	})
	convey.Convey("clear and update a dict", t, func() {
		d2 := d.Copy()
		assert.True(t, d.Equal(d2))
		assert.False(t, d2.Empty())
		_ = d2.Clear()
		assert.True(t, d2.Empty())
		_ = d2.Update(d.Dict)
		assert.True(t, d.Equal(d2))
	})
}
