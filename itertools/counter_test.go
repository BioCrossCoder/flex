package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCounter(t *testing.T) {
	convey.Convey("count", t, func() {
		c := NewCounter(10, 2, true)
		assert.Equal(t, 10, c.Count())
		for i := 1; i < 10; i++ {
			assert.Equal(t, 10-i*2, c.Count())
		}
	})
	convey.Convey("reset", t, func() {
		c := NewCounter(6, 3, false)
		assert.Equal(t, 6, c.Count())
		for i := 0; i < 10; i++ {
			_ = c.Count()
		}
		assert.NotEqual(t, 6, c.Count())
		c.Reset()
		assert.Equal(t, 6, c.Count())
	})
	convey.Convey("jump", t, func() {
		c := NewCounter(1, 5, false)
		assert.Equal(t, 1, c.Count())
		assert.Equal(t, 11, c.Jump(2))
	})
}
