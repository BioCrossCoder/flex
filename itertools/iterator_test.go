package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListIterator(t *testing.T) {
	convey.Convey("iterate all", t, func() {
		entry := []any{1, 2, 3, 4, 5}
		iterator := NewListIterator(entry)
		result := iterator.Pour()
		assert.Equal(t, entry, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}

func TestMapIterator(t *testing.T) {
	convey.Convey("iterate all", t, func() {
		entry := map[any]any{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
		iterator := NewMapIterator(entry)
		result := iterator.Pour()
		assert.Equal(t, entry, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}