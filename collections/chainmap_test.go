package collections

import (
	"flex/collections/dict"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestChainMap(t *testing.T) {
	d1 := dict.Dict{"a": 1, "b": 2}
	d2 := dict.Dict{"a": 4, "c": 5}
	convey.Convey("search value by key in ChainMap", t, func() {
		cm := NewChainMap(&d1, &d2)
		a, ok := cm.Get("a")
		assert.True(t, ok)
		assert.Equal(t, d1.Get("a"), a)
		b, ok := cm.Get("b")
		assert.True(t, ok)
		assert.Equal(t, d1.Get("b"), b)
		c, ok := cm.Get("c")
		assert.True(t, ok)
		assert.Equal(t, d2.Get("c"), c)
		_ = d2.Set("c", 3)
		c, ok = cm.Get("c")
		assert.True(t, ok)
		assert.Equal(t, 3, c)
		d, ok := cm.Get("d")
		assert.False(t, ok)
		assert.Nil(t, d)
	})
	convey.Convey("set value of key in ChainMap", t, func() {
		cm := NewChainMap(&d1, &d2)
		a, ok := cm.Get("a")
		assert.True(t, ok)
		assert.Equal(t, d1.Get("a"), a)
		assert.False(t, cm.Items().Has("a"))
		cm.Set("a", 3)
		assert.True(t, cm.Items().Has("a"))
		a, ok = cm.Get("a")
		assert.True(t, ok)
		assert.Equal(t, 3, a)
	})
	convey.Convey("operate chain nodes on ChainMap", t, func() {
		cm := NewChainMap(&d1, &d2)
		child := cm.NewChild()
		assert.Equal(t, cm, child.Parent())
		assert.Zero(t, child.Items().Size())
		a, ok := child.Get("a")
		assert.True(t, ok)
		assert.Equal(t, 1, a)
		parents := child.Parents()
		assert.Equal(t, parents[0], cm)
		assert.Equal(t, parents[1].Items(), &d1)
		assert.Equal(t, parents[2].Items(), &d2)
		maps := child.Maps()
		assert.Equal(t, maps[0], child.Items())
		assert.Equal(t, maps[1], cm.Items())
		assert.Equal(t, maps[2], &d1)
		assert.Equal(t, maps[3], &d2)
	})
}
