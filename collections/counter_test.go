package collections

import (
	"flex/collections/list"
	"flex/collections/set"
	"math"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	l := list.List{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	convey.Convey("count elements in a list", t, func() {
		c := NewCounter(l, -1)
		nums := c.Elements()
		assert.Equal(t, set.Of(nums...), set.Of(l...))
		assert.Equal(t, c.Total(), len(l))
		maxCount := 0
		minCount := math.MaxInt
		for _, num := range nums {
			count := c.Get(num)
			if count > maxCount {
				maxCount = count
			}
			if count < minCount {
				minCount = count
			}
			assert.Equal(t, count, l.Count(num))
		}
		for _, num := range c.MostCommon() {
			assert.Equal(t, c.Get(num), maxCount)
		}
		for _, num := range c.LeastCommon() {
			assert.Equal(t, c.Get(num), minCount)
		}
	})
	convey.Convey("modify a counter", t, func() {
		c := NewCounter(l, -1)
		assert.Equal(t, c.Get(0), -1)
		_ = c.Set(0, 10)
		assert.Equal(t, c.Get(0), 10)
		_ = c.Increment(0, 2)
		assert.Equal(t, c.Get(0), 12)
		_ = c.Subtract(0, 1)
		assert.Equal(t, c.Get(0), 11)
		assert.True(t, c.Remove(0))
		assert.Equal(t, c.Get(0), -1)
		_ = c.SetDefault(6)
		assert.Equal(t, c.Get(0), 6)
	})
	convey.Convey("reset a counter", t, func() {
		c := NewCounter(l, -1)
		nums := c.Elements()
		rest := list.List{}
		for i, num := range nums {
			assert.Equal(t, c.Get(num), l.Count(num))
			if i%2 == 0 {
				assert.True(t, c.Remove(num))
				assert.Equal(t, c.Get(num), -1)
			} else {
				_ = rest.Push(num)
			}
		}
		_ = c.Reset()
		for _, num := range rest {
			assert.Equal(t, c.Get(num), -1)
			assert.True(t, c.Remove(num))
		}
		assert.Equal(t, len(c.Clear().Elements()), 0)
	})
}
