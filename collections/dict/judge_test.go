package dict

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHas(t *testing.T) {
	convey.Convey("check the value matches the key in the dict", t, func() {
		d := newTestDict()
		for _, item := range d.Items() {
			pair := item.([2]any)
			assert.Equal(t, pair[1], d.Get(pair[0]))
		}
	})
	convey.Convey("check if the values are included in the dict", t, func() {
		d := newTestDict()
		values := d.Values()
		for _, k := range d.Keys() {
			assert.True(t, d.Has(k))
			assert.True(t, values.Includes(d.Get(k)))
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
