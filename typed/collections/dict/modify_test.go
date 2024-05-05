package dict

import (
	"flex/common"
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClear(t *testing.T) {
	convey.Convey("clear a dict", t, func() {
		d := newTestDict()
		assert.False(t, d.Empty())
		_ = d.Clear()
		assert.True(t, d.Empty())
	})
}

func TestSet(t *testing.T) {
	convey.Convey("set a value by key in a dict", t, func() {
		d := newTestDict()
		_ = d.Set("key1", 666)
		assert.Equal(t, 666, d.Get("key1"))
	})
}

func TestDelete(t *testing.T) {
	convey.Convey("delete a key in a dict", t, func() {
		d := newTestDict()
		_ = d.Set("key1", 233)
		assert.True(t, d.Has("key1"))
		ok := d.Delete("key1")
		assert.True(t, ok)
		assert.False(t, d.Has("key1"))
	})
}

func TestPop(t *testing.T) {
	convey.Convey("pop a key from a dict", t, func() {
		d := newTestDict()
		_ = d.Set("key1", 454)
		assert.True(t, d.Has("key1"))
		v, err := d.Pop("key1")
		assert.Nil(t, err)
		assert.Equal(t, 454, v)
		assert.False(t, d.Has("key1"))
	})
	convey.Convey("pop a non-exist key from a dict", t, func() {
		d := newTestDict()
		_ = d.Delete("f")
		assert.False(t, d.Has("f"))
		v, err := d.Pop("f")
		assert.Equal(t, common.ErrKeyNotFound, err)
		assert.Zero(t, v)
	})
	convey.Convey("pop a non-exist key from a dict with default value", t, func() {
		d := newTestDict()
		_ = d.Delete("f")
		assert.False(t, d.Has("f"))
		v, err := d.Pop("f", -1)
		assert.Nil(t, err)
		assert.Equal(t, -1, v)
	})
}

func TestPopItem(t *testing.T) {
	convey.Convey("pop an item from a dict", t, func() {
		d := newTestDict()
		d2 := d.Copy()
		k, v, err := d.PopItem()
		assert.Nil(t, err)
		assert.Equal(t, d2.Get(k), v)
	})
}

func TestUpdate(t *testing.T) {
	convey.Convey("update a dict with another dict", t, func() {
		d := newTestDict()
		d2 := Dict[string, int]{"key1": 1, "key2": 2}
		_ = d.Update(d2)
		assert.Equal(t, d.Get("key1"), 1)
		assert.Equal(t, d.Get("key2"), 2)
	})
}
