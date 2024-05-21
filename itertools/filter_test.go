package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"testing"
)

func pass(any) bool {
	return true
}

func TestListFilter(t *testing.T) {
	convey.Convey("filter all", t, func() {
		entry := []any{1, 2, 3}
		iterator := NewListFilter(entry, pass)
		result := iterator.Pour()
		assert.Equal(t, entry, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}
