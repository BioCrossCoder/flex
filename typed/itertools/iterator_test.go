package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListIterator(t *testing.T) {
	convey.Convey("iterate all", t, func() {
		entry := []int{1, 2, 3, 4, 5}
		iterator := NewListIterator(entry)
		result := iterator.Pour()
		assert.Equal(t, entry, result)
		assert.False(t, iterator.Next())
		assert.Zero(t, iterator.Value())
	})
}
