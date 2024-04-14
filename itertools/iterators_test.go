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

func pass(data any) bool {
	return true
}

func TestMapConvertor(t *testing.T) {
	convey.Convey("iterate all", t, func() {
		entry := map[any]any{
			"one":   1,
			"two":   2,
			"three": 3,
		}
		iterator := NewMapIterator(entry, self)
		result := iterator.Pour()
		assert.Equal(t, entry, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}

func TestMapFilter(t *testing.T) {
	convey.Convey("filter all", t, func() {
		entry := map[any]any{
			1: "one",
			2: "two",
		}
		iterator := NewMapFilter(entry, pass)
		result := iterator.Pour()
		assert.Equal(t, entry, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}

func TestListConvertor(t *testing.T) {
	convey.Convey("iterate all", t, func() {
		entry := []any{1, 2, 3}
		iterator := NewListIterator(entry, self)
		result := iterator.Pour()
		assert.Equal(t, entry, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
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

func TestAccumulator(t *testing.T) {
	convey.Convey("accumulate all", t, func() {
		entry := []any{1, 2, 3}
		f := func(x, y int) int {
			return x + y
		}
		iterator := NewAccumulator(entry, func(p1, p2 any) any {
			params := []reflect.Value{reflect.ValueOf(p1), reflect.ValueOf(p2)}
			return reflect.ValueOf(f).Call(params)[0].Interface()
		})
		result := iterator.Pour()
		expected := 0
		for _, v := range entry {
			expected += v.(int)
		}
		assert.Equal(t, expected, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}

func TestZipIterator(t *testing.T) {
	convey.Convey("zip all", t, func() {
		entry1 := []any{1, 2, 3}
		entry2 := []any{"a", "b", "c"}
		expected := [][2]any{
			{1, "a"},
			{2, "b"},
		}
		iterator := NewZipIterator(entry1, entry2, 2)
		result := iterator.Pour()
		assert.Equal(t, expected, result)
		assert.False(t, iterator.Next())
		assert.Nil(t, iterator.Value())
	})
}
