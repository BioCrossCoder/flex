package dict

import (
	"flex/typed/collections/arraylist"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newTestDict() Dict[string, int] {
	return Dict[string, int]{"a": 1, "b": 2, "c": 3}
}

func TestSize(t *testing.T) {
	convey.Convey("get size of a dict", t, func() {
		d := newTestDict()
		assert.Equal(t, len(d), d.Size())
	})
}

func TestGet(t *testing.T) {
	m := newTestDict()
	convey.Convey("get value by key from a dict", t, func() {
		assert.Equal(t, m["a"], m.Get("a"))
	})
	convey.Convey("get value by default of a key not exist in dict", t, func() {
		assert.Equal(t, 0, m.Get("d", 0))
	})
}

func TestCopy(t *testing.T) {
	convey.Convey("copy a dict", t, func() {
		d := newTestDict()
		assert.True(t, d.Equal(d.Copy()))
	})
}

func TestHas(t *testing.T) {
	convey.Convey("check the value matches the key in the dict", t, func() {
		d := newTestDict()
		for _, item := range d.Items() {
			assert.Equal(t, item.Value, d.Get(item.Key))
		}
	})
	convey.Convey("check if the values are included in the dict", t, func() {
		d := newTestDict()
		values := d.Values()
		for _, k := range d.Keys() {
			assert.True(t, d.Has(k))
			assert.True(t, arraylist.Of(values...).Includes(d.Get(k)))
		}
	})
}

func TestEmpty(t *testing.T) {
	convey.Convey("check if the dict is empty", t, func() {
		d := newTestDict()
		assert.False(t, d.Empty())
		_ = d.Clear()
		assert.True(t, d.Empty())
	})
}
