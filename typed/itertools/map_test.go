package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListConvertor(t *testing.T) {
	convey.Convey("iterate all", t, func() {
		entry := []int{1, 2, 3}
		f := func(x int) int {
			return x * 2
		}
		iterator := Map(f, entry)
		assert.Equal(t, []int{2, 4, 6}, iterator.Pour())
		assert.False(t, iterator.Next())
		assert.Zero(t, iterator.Value())
	})
}
