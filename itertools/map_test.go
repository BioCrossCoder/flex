// Package itertools provides iterator functions to create iterators and perform common operations on iterables.
package itertools

import (
	"github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func self(data any) any {
	return reflect.ValueOf(data).Interface()
}

func TestListConvertor(t *testing.T) {
	convey.Convey("iterate all", t, func() {
		entry := []any{1, 2, 3}
		iterator := NewListConvertor(entry, self)
		result := iterator.Pour()
		assert.Equal(t, entry, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}
